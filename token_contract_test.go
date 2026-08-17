package lease

import (
	"testing"
	"time"
)

func TestReleaseRejectsAStaleTokenAfterReacquire(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	first, err := store.Acquire(now, AcquireRequest{Key: "worker-b", Holder: "node-1", Duration: time.Minute})
	if err != nil { t.Fatal(err) }
	second, err := store.Acquire(first.ExpiresAt.Add(time.Second), AcquireRequest{Key: "worker-b", Holder: "node-1", Duration: time.Minute})
	if err != nil { t.Fatal(err) }
	if err := store.Release(second.ExpiresAt.Add(-time.Second), "worker-b", "node-1", first.Token); err != ErrTokenMismatch {
		t.Fatalf("Release() error = %v, want ErrTokenMismatch", err)
	}
}
