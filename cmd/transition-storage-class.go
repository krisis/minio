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

	"github.com/minio/minio/pkg/hash"
	"github.com/minio/minio/pkg/madmin"
)

var transitionStorageClassConfigPath string = path.Join(minioConfigPrefix, "transition-storage-class-config.json")

type TransitionStorageClassConfig struct {
	sync.RWMutex
	drivercache map[string]warmBackend
	S3          map[string]madmin.TransitionStorageClassS3    `json:"s3"`
	Azure       map[string]madmin.TransitionStorageClassAzure `json:"azure"`
	GCS         map[string]madmin.TransitionStorageClassGCS   `json:"gcs"`
}

func (config *TransitionStorageClassConfig) Add(i interface{}) error {
	config.Lock()
	defer config.Unlock()

	switch c := i.(type) {
	case madmin.TransitionStorageClassS3:
		_, ok := config.S3[c.Name]
		if ok {
			return errInvalidArgument
		}
		config.S3[c.Name] = c
	case madmin.TransitionStorageClassAzure:
		_, ok := config.Azure[c.Name]
		if ok {
			return errInvalidArgument
		}
		config.Azure[c.Name] = c
	case madmin.TransitionStorageClassGCS:
		_, ok := config.GCS[c.Name]
		if ok {
			return errInvalidArgument
		}
		config.GCS[c.Name] = c
	default:
		return errInvalidArgument
	}
	return nil
}

func (config *TransitionStorageClassConfig) Edit(i interface{}) error {
	config.Lock()
	defer config.Unlock()

	switch c := i.(type) {
	case madmin.TransitionStorageClassS3:
		config.S3[c.Name] = c
	case madmin.TransitionStorageClassAzure:
		config.Azure[c.Name] = c
	case madmin.TransitionStorageClassGCS:
		config.GCS[c.Name] = c
	default:
		return errInvalidArgument
	}
	return nil
}

func (config *TransitionStorageClassConfig) RemoveStorageClass(name string) {
	config.Lock()
	defer config.Unlock()

	// FIXME: check if the SC is used by any of the ILM policies.

	delete(config.S3, name)
	delete(config.Azure, name)
	delete(config.GCS, name)
}

func (config *TransitionStorageClassConfig) Byte() ([]byte, error) {
	config.Lock()
	defer config.Unlock()
	return json.Marshal(config)
}

func (config *TransitionStorageClassConfig) GetDriver(sc string) (warmBackend, error) {
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
	b, err := json.Marshal(globalTransitionStorageClassConfig)
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
			globalTransitionStorageClassConfig = &TransitionStorageClassConfig{}
		}
		return err
	}
	var config TransitionStorageClassConfig
	err = json.Unmarshal(buf.Bytes(), &config)
	if err != nil {
		return err
	}
	globalTransitionStorageClassConfig = &config
	return nil
}
