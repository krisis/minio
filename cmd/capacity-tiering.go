package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aalpar/deheap"
	"github.com/minio/minio-go/v7/pkg/set"
	"github.com/minio/minio/internal/color"
	"github.com/minio/minio/internal/logger"
	"github.com/minio/pkg/console"
)

// tierCandidateCache contains objects older than 6 months in tiering enabled
// buckets
type tierCandidateCache struct {
	sync.RWMutex
	entries    *tierEntries
	set        set.StringSet
	tierHighWM int
	tierLowWM  int
	cap        int
	debug      bool
}

type tierEntry struct {
	Bucket    string
	Name      string
	VersionID string
	ModTime   time.Time
	Size      int64
}

func (t tierEntry) hash() string {
	return fmt.Sprintf("%s/%s/%s", t.Bucket, t.Name, t.VersionID)
}
func (oi ObjectInfo) TierEntry() tierEntry {
	return tierEntry{
		Name:      oi.Name,
		VersionID: oi.VersionID,
		Bucket:    oi.Bucket,
		ModTime:   oi.ModTime,
		Size:      oi.Size,
	}
}

func newTierCandidateCache() *tierCandidateCache {
	entries := new(tierEntries)
	maxEntry := 10000
	*entries = make(tierEntries, 0, maxEntry)
	fmt.Println("len", len(*entries), "cap", cap(*entries))
	return &tierCandidateCache{
		entries:    entries,
		set:        set.NewStringSet(),
		tierHighWM: 0,
		tierLowWM:  0,
		// tierHighWM: 85,
		// tierLowWM:  75,
		cap:   maxEntry,
		debug: serverDebugLog,
	}
}

func (tc *tierCandidateCache) tier(ctx context.Context) {
	capacityTieringLogPrefix := color.Green("capacityTiering:")
	ticker := time.NewTicker(2 * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("tierCandidateCache.tier", tc.debug)
			obj := newObjectLayerFn()
			// FIXME: use timedValue to periodically refresh storageInfo, usableSpace and dui
			storageInfo, _ := obj.StorageInfo(ctx)
			usableSpace := GetTotalUsableCapacity(storageInfo.Disks, storageInfo)
			dui, err := loadDataUsageFromBackend(ctx, obj)
			if err != nil {
				logger.LogIf(ctx, fmt.Errorf("tierCandidateCache.tier: failed to load data usage %w", err))
				continue
			}

			usedSpace := float64(dui.ObjectsTotalSize)
			if usedSpace*100 < float64(tc.tierHighWM)*usableSpace {
				continue
			}

			until := usedSpace - float64(tc.tierLowWM)*0.01*usableSpace
			for n := tc.Len(); n > 0 && until > 0; n-- {
				e := tc.Remove()
				until -= float64(e.Size)
				if tc.debug {
					console.Debugf(capacityTieringLogPrefix+" tiering %s %s %d %v", e.Bucket, e.Name, e.Size, e.ModTime)
				}
				// TODO: transition e to bucket's capacity tier
			}

		}
	}
}

// Add adds a tierEntry to the bucket's candidate list. If list is full, replace
// a smaller (newer) entry.
func (tc *tierCandidateCache) Add(e tierEntry) {
	tc.Lock()
	defer tc.Unlock()
	if tc.set.Contains(e.hash()) {
		return
	}
	fmt.Println("adding", e)

	if len(*tc.entries) >= tc.cap {
		// Replace smallest (newer) entry with e
		deheap.Pop(tc.entries)
	}
	deheap.Push(tc.entries, e)
	tc.set.Add(e.hash())
}

func (tc *tierCandidateCache) Remove() tierEntry {
	tc.Lock()
	defer tc.Unlock()
	e := deheap.PopMax(tc.entries).(tierEntry)
	tc.set.Remove(e.hash())
	fmt.Println("removing", e)
	return e
}

func (tc *tierCandidateCache) Len() int {
	tc.RLock()
	defer tc.RUnlock()
	return tc.entries.Len()
}

type tierEntries []tierEntry

func (te tierEntries) Len() int {
	return len(te)
}

func (te tierEntries) Swap(i, j int) {
	te[i], te[j] = te[j], te[i]
}

func (te tierEntries) Less(i, j int) bool {
	if te[i].Size == te[j].Size {
		return te[i].ModTime.After(te[j].ModTime)
	}
	return te[i].Size < te[j].Size
}

func (te *tierEntries) Push(x interface{}) {
	*te = append(*te, x.(tierEntry))
}

func (te *tierEntries) Pop() interface{} {
	old := *te
	n := len(old)
	item := old[n-1]
	*te = old[:n-1]
	return item
}
