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
 *
 */

package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type StorageClassType int

const (
	unsupported StorageClassType = iota
	S3
	Azure
	GCS
)

func (st StorageClassType) String() string {
	switch st {
	case S3:
		return "s3"
	case Azure:
		return "azure"
	case GCS:
		return "gcs"
	}
	return "unsupported"
}

func NewStorageClassType(scType string) (StorageClassType, error) {
	switch scType {
	case S3.String():
		return S3, nil
	case Azure.String():
		return Azure, nil
	case GCS.String():
		return GCS, nil
	}

	return unsupported, errors.New("Unsupported storage class type")
}

type TransitionStorageClassConfig struct {
	Type  StorageClassType
	S3    *TransitionStorageClassS3
	Azure *TransitionStorageClassAzure
	GCS   *TransitionStorageClassGCS
}

func (cfg *TransitionStorageClassConfig) Name() string {
	switch cfg.Type {
	case S3:
		return cfg.S3.Name
	case Azure:
		return cfg.Azure.Name
	case GCS:
		return cfg.Azure.Name
	}
	panic("unexpected transition storage-class type")
}

const StorageClassAPI = "transition-storage-class"

func (adm *AdminClient) AddStorageClass(ctx context.Context, cfg TransitionStorageClassConfig) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	// FIXME: encrypt storage-class config payload

	queryValues := url.Values{}
	queryValues.Set("add", "")

	reqData := requestData{
		relPath:     strings.Join([]string{adminAPIPrefix, StorageClassAPI}, "/"),
		queryValues: queryValues,
		content:     data,
	}

	// Execute PUT on /minio/admin/v3/transition-storage-class?add to add a transition storage-class.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)

	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}
	return nil
}
