package cmd

import (
	"testing"
	"time"

	"github.com/google/uuid"
	xhttp "github.com/minio/minio/cmd/http"
	"github.com/minio/minio/pkg/bucket/lifecycle"
)

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
