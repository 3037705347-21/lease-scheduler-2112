package lease

import (
	"testing"
	"time"
)

func TestSnapshotOrdersActiveLeasesByAscendingKey(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	for _, key := range []string{"worker-z", "worker-a", "worker-m"} {
		if _, err := store.Acquire(now, AcquireRequest{Key: key, Holder: "node-1", Duration: time.Minute}); err != nil { t.Fatal(err) }
	}
	got := store.Snapshot(now)
	if len(got) != 3 || got[0].Key != "worker-a" || got[1].Key != "worker-m" || got[2].Key != "worker-z" {
		t.Fatalf("Snapshot() = %#v, want ascending keys", got)
	}
}
