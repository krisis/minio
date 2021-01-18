/*
 * MinIO Cloud Storage, (C) 2018-2019 MinIO, Inc.
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
	"context"
	"encoding/json"
	"path"
	"sync"

	"github.com/minio/minio-go/v7/pkg/set"
	"github.com/minio/minio/pkg/hash"
	"github.com/minio/minio/pkg/madmin"
)

var transitionStorageClassConfigPath string = path.Join(minioConfigPrefix, "transition-storage-class-config.json")

type TransitionStorageClassConfigMgr struct {
	sync.RWMutex
	storageClassNames set.StringSet
	drivercache       map[string]warmBackend
	S3                map[string]madmin.TransitionStorageClassS3    `json:"s3"`
	Azure             map[string]madmin.TransitionStorageClassAzure `json:"azure"`
	GCS               map[string]madmin.TransitionStorageClassGCS   `json:"gcs"`
}

func (config *TransitionStorageClassConfigMgr) Add(sc madmin.TransitionStorageClassConfig) error {
	config.Lock()
	defer config.Unlock()

	scName := sc.Name()
	// storage-class name already in use
	if config.storageClassNames.Contains(scName) {
		return errInvalidArgument // FIXME(kp): errTransitionStorageClassAlreadyExists?
	}

	switch sc.Type {
	case madmin.S3:
		config.S3[scName] = *sc.S3

	case madmin.Azure:
		config.Azure[scName] = *sc.Azure

	case madmin.GCS:
		config.GCS[scName] = *sc.GCS
	}

	return errInvalidArgument
}

func (config *TransitionStorageClassConfigMgr) Edit(sc madmin.TransitionStorageClassConfig) error {
	config.Lock()
	defer config.Unlock()

	scName := sc.Name()
	// no storage-class by this name exists
	if !config.storageClassNames.Contains(scName) {
		return errInvalidArgument // FIXME(kp): errTransitionStorageClassNotFound?
	}

	switch sc.Type {
	case madmin.S3:
		config.S3[scName] = *sc.S3

	case madmin.Azure:
		config.Azure[scName] = *sc.Azure

	case madmin.GCS:
		config.GCS[scName] = *sc.GCS
	}
	return errInvalidArgument
}

func (config *TransitionStorageClassConfigMgr) RemoveStorageClass(name string) {
	config.Lock()
	defer config.Unlock()

	// FIXME: check if the SC is used by any of the ILM policies.

	delete(config.S3, name)
	delete(config.Azure, name)
	delete(config.GCS, name)
}

func (config *TransitionStorageClassConfigMgr) Bytes() ([]byte, error) {
	config.Lock()
	defer config.Unlock()
	return json.Marshal(config)
}

func (config *TransitionStorageClassConfigMgr) GetDriver(sc string) (warmBackend, error) {
	config.Lock()
	defer config.Unlock()

	d := config.drivercache[sc]
	if d != nil {
		return d, nil
	}
	for k, v := range config.S3 {
		if k == sc {
			return newWarmBackendS3(v)
		}
	}
	for k, v := range config.Azure {
		if k == sc {
			return newWarmBackendAzure(v)
		}
	}
	for k, v := range config.GCS {
		if k == sc {
			return newWarmBackendGCS(v)
		}
	}
	return nil, errInvalidArgument
}

func saveGlobalTransitionStorageClassConfig() error {
	b, err := json.Marshal(globalTransitionStorageClassConfigMgr)
	if err != nil {
		return err
	}
	r, err := hash.NewReader(bytes.NewReader(b), int64(len(b)), "", "", int64(len(b)), false)
	if err != nil {
		return err
	}
	_, err = globalObjectAPI.PutObject(context.Background(), minioMetaBucket, transitionStorageClassConfigPath, NewPutObjReader(r, nil, nil), ObjectOptions{})
	return err
}

func loadGlobalTransitionStorageClassConfig() error {
	var buf bytes.Buffer
	err := globalObjectAPI.GetObject(context.Background(), minioMetaBucket, transitionStorageClassConfigPath, 0, -1, &buf, "", ObjectOptions{})
	if err != nil {
		if isErrObjectNotFound(err) {
			globalTransitionStorageClassConfigMgr = &TransitionStorageClassConfigMgr{}
		}
		return err
	}
	var config TransitionStorageClassConfigMgr
	err = json.Unmarshal(buf.Bytes(), &config)
	if err != nil {
		return err
	}

	// Build the set of (unique) user-defined transition storage-class names
	// from transition-storage-class-config.json
	storageClassNames := set.NewStringSet()
	for scName, _ := range config.S3 {
		storageClassNames.Add(scName)
	}
	for scName, _ := range config.Azure {
		storageClassNames.Add(scName)
	}
	for scName, _ := range config.GCS {
		storageClassNames.Add(scName)
	}
	config.storageClassNames = storageClassNames

	globalTransitionStorageClassConfigMgr = &config
	return nil
}
