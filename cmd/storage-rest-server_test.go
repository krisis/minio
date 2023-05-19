// Copyright (c) 2023 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bytes"
	"errors"
	"hash/crc64"
	"io"
	"net/http/httptest"
	"testing"
)

// TestKeepAliveHash tests if keepHTTPReqResponseAliveHash function works as
// expected with waitForHTTPResponseHash
func TestKeepAliveHash(t *testing.T) {
	tests := []struct {
		data        []byte
		expectedErr error
	}{
		{
			data:        []byte("Hello World!"),
			expectedErr: nil,
		},
		{
			data:        []byte("Hello!"),
			expectedErr: errors.New("sample error"),
		},
		{
			data:        bytes.Repeat([]byte("a"), 1000),
			expectedErr: errors.New("another error"),
		},
	}
	for i, test := range tests {
		r := httptest.NewRequest("GET", "http://example.com/foo", bytes.NewReader(test.data))
		w := httptest.NewRecorder()

		done, body := keepHTTPReqResponseAliveHash(w, r)
		h := crc64.New(crc64.MakeTable(crc64.ISO))
		tee := io.TeeReader(body, h)
		var expectedHash uint64
		_, err := io.ReadAll(tee)
		if err != nil {
			t.Fatalf("%d: failed with %v", i, err)
		}
		done(expectedHash, test.expectedErr)

		respBody, _ := io.ReadAll(w.Result().Body)
		_, gotHash, gotErr := waitForHTTPResponseHash(bytes.NewReader(respBody))
		if expectedHash != gotHash {
			t.Fatalf("Expected crc64 hash %d but got %d", expectedHash, gotHash)
		}
		if test.expectedErr != nil && test.expectedErr.Error() != gotErr.Error() {
			t.Fatalf("Expected error %v but got %v", test.expectedErr, gotErr)
		}
	}
}
