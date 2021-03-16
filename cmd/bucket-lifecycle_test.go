/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
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
	"net/http"
	"testing"
	"time"

	xhttp "github.com/minio/minio/cmd/http"
)

// TestParseRestoreObjStatus tests parseRestoreObjStatus
func TestParseRestoreObjStatus(t *testing.T) {
	testCases := []struct {
		restoreHdr     map[string]string
		expectedStatus restoreObjStatus
		expectedErr    error
	}{
		{
			// valid: represents a restored object, 'pending' expiry.
			restoreHdr: map[string]string{
				xhttp.AmzRestore: "ongoing-request=false, expiry-date=Fri, 21 Dec 2012 00:00:00 GMT",
			},
			expectedStatus: restoreObjStatus{
				ongoing: false,
				expiry:  time.Date(2012, 12, 21, 0, 0, 0, 0, time.UTC),
			},
			expectedErr: nil,
		},
		{
			// valid: represents an ongoing restore object request.
			restoreHdr: map[string]string{
				xhttp.AmzRestore: "ongoing-request=true",
			},
			expectedStatus: restoreObjStatus{
				ongoing: true,
			},
			expectedErr: nil,
		},
		{
			// invalid; ongoing restore object request can't have expiry set on it.
			restoreHdr: map[string]string{
				xhttp.AmzRestore: "ongoing-request=true, expiry-date=Fri, 21 Dec 2012 00:00:00 GMT",
			},
			expectedStatus: restoreObjStatus{},
			expectedErr:    errRestoreHDRMalformed,
		},
		{
			// invalid; restore header missing.
			restoreHdr:     map[string]string{},
			expectedStatus: restoreObjStatus{},
			expectedErr:    errRestoreHDRMissing,
		},
	}
	for i, tc := range testCases {
		actual, err := parseRestoreObjStatus(tc.restoreHdr)
		if err != tc.expectedErr {
			t.Fatalf("Test %d: got %v expected %v", i+1, err, tc.expectedErr)
		}
		if actual != tc.expectedStatus {
			t.Fatalf("Test %d: got %v expected %v", i+1, actual, tc.expectedStatus)
		}
	}
}

// TestRestoreObjStatusRoundTrip restoreObjStatus roundtrip
func TestRestoreObjStatusRoundTrip(t *testing.T) {
	testCases := []restoreObjStatus{
		ongoingRestoreObj(),
		completedRestoreObj(time.Now().UTC()),
	}
	for i, tc := range testCases {
		actual, err := parseRestoreObjStatus(map[string]string{xhttp.AmzRestore: tc.String()})
		if err != nil {
			t.Fatalf("Test %d: parse restore object failed: %v", i+1, err)
		}
		if actual.ongoing != tc.ongoing || actual.expiry.Format(http.TimeFormat) != tc.expiry.Format(http.TimeFormat) {
			t.Fatalf("Test %d: got %v expected %v", i+1, actual, tc)
		}
	}
}
