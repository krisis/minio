// Copyright (c) 2015-2021 MinIO, Inc.
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
	"context"

	"github.com/minio/minio/internal/bucket/lifecycle"
)

type transitioner interface {
	TransitionObject(ctx context.Context, objectAPI ObjectLayer) error
	ObjectInfo() ObjectInfo
}

type transitionInfo ObjectInfo

func (toi transitionInfo) ObjectInfo() ObjectInfo {
	return ObjectInfo(toi)
}

// TransitionObject to target specified by the transition ARN. When an object is transitioned to another
// storage specified by the transition ARN, the metadata is left behind on source cluster and original content
// is moved to the transition tier. Note that in the case of encrypted objects, entire encrypted stream is moved
// to the transition tier without decrypting or re-encrypting.
func (toi transitionInfo) TransitionObject(ctx context.Context, objectAPI ObjectLayer) error {
	lc, err := globalLifecycleSys.Get(toi.Bucket)
	if err != nil {
		return err
	}
	opts := ObjectOptions{
		Transition: TransitionOptions{
			Status: lifecycle.TransitionPending,
			Tier:   lc.TransitionTier(ObjectInfo(toi).ToLifecycleOpts()),
			ETag:   toi.ETag,
		},
		VersionID: toi.VersionID,
		Versioned: globalBucketVersioningSys.Enabled(toi.Bucket),
		MTime:     toi.ModTime,
	}
	return objectAPI.TransitionObject(ctx, toi.Bucket, toi.Name, opts)

}

type tierInfo ObjectInfo

func (ti tierInfo) ObjectInfo() ObjectInfo {
	return ObjectInfo(ti)
}
func (ti tierInfo) TransitionObject(ctx context.Context, objectAPI ObjectLayer) error {
	// TODO: Fetch tier name from storage_tiering config subsystem
	return nil
}
