/*
 * MinIO Cloud Storage, (C) 2021 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/uuid"
	xhttp "github.com/minio/minio/cmd/http"
	"github.com/minio/minio/pkg/bucket/lifecycle"
)

func TestXLV2FormatData(t *testing.T) {
	failOnErr := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	data := []byte("some object data")
	data2 := []byte("some other object data")

	xl := xlMetaV2{}
	fi := FileInfo{
		Volume:           "volume",
		Name:             "object-name",
		VersionID:        "756100c6-b393-4981-928a-d49bbc164741",
		IsLatest:         true,
		Deleted:          false,
		TransitionStatus: "",
		DataDir:          "bffea160-ca7f-465f-98bc-9b4f1c3ba1ef",
		XLV1:             false,
		ModTime:          time.Now(),
		Size:             0,
		Mode:             0,
		Metadata:         nil,
		Parts:            nil,
		Erasure: ErasureInfo{
			Algorithm:    ReedSolomon.String(),
			DataBlocks:   4,
			ParityBlocks: 2,
			BlockSize:    10000,
			Index:        1,
			Distribution: []int{1, 2, 3, 4, 5, 6, 7, 8},
			Checksums: []ChecksumInfo{{
				PartNumber: 1,
				Algorithm:  HighwayHash256S,
				Hash:       nil,
			}},
		},
		MarkDeleted:                   false,
		DeleteMarkerReplicationStatus: "",
		VersionPurgeStatus:            "",
		Data:                          data,
		NumVersions:                   1,
		SuccessorModTime:              time.Time{},
	}

	failOnErr(xl.AddVersion(fi))

	fi.VersionID = mustGetUUID()
	fi.DataDir = mustGetUUID()
	fi.Data = data2
	failOnErr(xl.AddVersion(fi))

	serialized, err := xl.AppendTo(nil)
	failOnErr(err)
	// Roundtrip data
	var xl2 xlMetaV2
	failOnErr(xl2.Load(serialized))

	// We should have one data entry
	list, err := xl2.data.list()
	failOnErr(err)
	if len(list) != 2 {
		t.Fatalf("want 1 entry, got %d", len(list))
	}

	if !bytes.Equal(xl2.data.find("756100c6-b393-4981-928a-d49bbc164741"), data) {
		t.Fatal("Find data returned", xl2.data.find("756100c6-b393-4981-928a-d49bbc164741"))
	}
	if !bytes.Equal(xl2.data.find(fi.VersionID), data2) {
		t.Fatal("Find data returned", xl2.data.find(fi.VersionID))
	}

	// Remove entry
	xl2.data.remove(fi.VersionID)
	failOnErr(xl2.data.validate())
	if xl2.data.find(fi.VersionID) != nil {
		t.Fatal("Data was not removed:", xl2.data.find(fi.VersionID))
	}
	if xl2.data.entries() != 1 {
		t.Fatal("want 1 entry, got", xl2.data.entries())
	}
	// Re-add
	xl2.data.replace(fi.VersionID, fi.Data)
	failOnErr(xl2.data.validate())
	if xl2.data.entries() != 2 {
		t.Fatal("want 2 entries, got", xl2.data.entries())
	}

	// Replace entry
	xl2.data.replace("756100c6-b393-4981-928a-d49bbc164741", data2)
	failOnErr(xl2.data.validate())
	if xl2.data.entries() != 2 {
		t.Fatal("want 2 entries, got", xl2.data.entries())
	}
	if !bytes.Equal(xl2.data.find("756100c6-b393-4981-928a-d49bbc164741"), data2) {
		t.Fatal("Find data returned", xl2.data.find("756100c6-b393-4981-928a-d49bbc164741"))
	}

	if !xl2.data.rename("756100c6-b393-4981-928a-d49bbc164741", "new-key") {
		t.Fatal("old key was not found")
	}
	failOnErr(xl2.data.validate())
	if !bytes.Equal(xl2.data.find("new-key"), data2) {
		t.Fatal("Find data returned", xl2.data.find("756100c6-b393-4981-928a-d49bbc164741"))
	}
	if xl2.data.entries() != 2 {
		t.Fatal("want 2 entries, got", xl2.data.entries())
	}
	if !bytes.Equal(xl2.data.find(fi.VersionID), data2) {
		t.Fatal("Find data returned", xl2.data.find(fi.DataDir))
	}

	// Test trimmed
	xl2 = xlMetaV2{}
	trimmed := xlMetaV2TrimData(serialized)
	failOnErr(xl2.Load(trimmed))
	if len(xl2.data) != 0 {
		t.Fatal("data, was not trimmed, bytes left:", len(xl2.data))

	}
	// Corrupt metadata, last 5 bytes is the checksum, so go a bit further back.
	trimmed[len(trimmed)-10] += 10
	if err := xl2.Load(trimmed); err == nil {
		t.Fatal("metadata corruption not detected")
	}
}

// TestUsesDataDir tests xlMetaV2.UsesDataDir
func TestUsesDataDir(t *testing.T) {
	vID := uuid.New()
	dataDir := uuid.New()
	transitioned := make(map[string][]byte)
	transitioned[ReservedMetadataPrefixLower+TransitionStatus] = []byte(lifecycle.TransitionComplete)

	toBeRestored := make(map[string]string)
	toBeRestored[xhttp.AmzRestore] = ongoingRestoreObj().String()

	restored := make(map[string]string)
	restored[xhttp.AmzRestore] = completedRestoreObj(time.Now().UTC().Add(time.Hour)).String()

	restoredExpired := make(map[string]string)
	restoredExpired[xhttp.AmzRestore] = completedRestoreObj(time.Now().UTC().Add(-time.Hour)).String()

	testCases := []struct {
		xlmeta xlMetaV2Object
		uses   bool
	}{
		{ // transitioned object version
			xlmeta: xlMetaV2Object{
				VersionID: vID,
				DataDir:   dataDir,
				MetaSys:   transitioned,
			},
			uses: false,
		},
		{ // to be restored (requires object version to be transitioned)
			xlmeta: xlMetaV2Object{
				VersionID: vID,
				DataDir:   dataDir,
				MetaSys:   transitioned,
				MetaUser:  toBeRestored,
			},
			uses: false,
		},
		{ // restored object version (requires object version to be transitioned)
			xlmeta: xlMetaV2Object{
				VersionID: vID,
				DataDir:   dataDir,
				MetaSys:   transitioned,
				MetaUser:  restored,
			},
			uses: true,
		},
		{ // restored object version expired an hour back (requires object version to be transitioned)
			xlmeta: xlMetaV2Object{
				VersionID: vID,
				DataDir:   dataDir,
				MetaSys:   transitioned,
				MetaUser:  restoredExpired,
			},
			uses: false,
		},
		{ // object version with no ILM applied
			xlmeta: xlMetaV2Object{
				VersionID: vID,
				DataDir:   dataDir,
			},
			uses: true,
		},
	}
	for i, tc := range testCases {
		if got := tc.xlmeta.UsesDataDir(); got != tc.uses {
			t.Fatalf("Test %d: Expected %v but got %v for %v", i+1, tc.uses, got, tc.xlmeta)
		}
	}
}
