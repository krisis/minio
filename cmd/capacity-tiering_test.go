package cmd

import (
	"container/heap"
	"fmt"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
)

func TestTierCandidateCache(t *testing.T) {
	entry := func(i int, now time.Time) tierEntry {
		return tierEntry{
			Bucket:  "bucket",
			Name:    fmt.Sprintf("object-%d", i),
			ModTime: now,
			Size:    int64(i * humanize.MiByte),
		}
	}
	tc := newTierCandidateCache()
	now := time.Now()

	i := 0
	for ; i < maxTierEntry; i++ {
		tc.Add(entry(i, now))
	}
	tc.Add(entry(i, now))
	entries := tc.entries
	got := heap.Pop(entries)
	expected := entry(1, now)
	if got != expected {
		t.Fatalf("expected %#v but got %#v", expected, got)
	}
}
