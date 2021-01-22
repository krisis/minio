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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const StorageClassAPI = "transition-storage-class"

func (adm *AdminClient) AddStorageClass(ctx context.Context, cfg TransitionStorageClassConfig) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	encData, err := EncryptData(adm.getSecretKey(), data)
	if err != nil {
		return err
	}

	queryValues := url.Values{}
	queryValues.Set("add", "")

	reqData := requestData{
		relPath:     strings.Join([]string{adminAPIPrefix, StorageClassAPI}, "/"),
		queryValues: queryValues,
		content:     encData,
	}

	// Execute PUT on /minio/admin/v3/transition-storage-class to add a
	// transition storage-class.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return httpRespToErrorResponse(resp)
	}
	return nil
}

func (adm *AdminClient) ListStorageClasses(ctx context.Context) ([]TransitionStorageClassConfig, error) {
	reqData := requestData{
		relPath: strings.Join([]string{adminAPIPrefix, StorageClassAPI}, "/"),
	}

	// Execute GET on /minio/admin/v3/transition-storage-class to list
	// transition storage-classes configured.
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	var storageClasses []TransitionStorageClassConfig
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return storageClasses, err
	}

	err = json.Unmarshal(b, &storageClasses)
	if err != nil {
		return storageClasses, err
	}

	return storageClasses, nil
}

type StorageClassCreds struct {
	AccessKey string `json:"access,omitempty"`
	SecretKey string `json:"secret,omitempty"`
	CredsJSON []byte `json:"creds,omitempty"`
}

func (adm *AdminClient) EditStorageClass(ctx context.Context, scName string, creds StorageClassCreds) error {
	data, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	encData, err := EncryptData(adm.getSecretKey(), data)
	reqData := requestData{
		relPath: strings.Join([]string{adminAPIPrefix, StorageClassAPI, scName}, "/"),
		content: encData,
	}

	// Execute POST on /minio/admin/v3/transition-storage-class/storageClassName" to edit
	// transition storage-classes configured.
	resp, err := adm.executeMethod(ctx, http.MethodPost, reqData)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return httpRespToErrorResponse(resp)
	}

	return nil
}
