package lease

import (
	"testing"
	"time"
)

func TestAcquireReclaimsLeaseAtItsExactExpiry(t *testing.T) {
	store := NewStore()
	now := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	first, err := store.Acquire(now, AcquireRequest{Key: "worker-a", Holder: "node-1", Duration: time.Minute})
	if err != nil {
		t.Fatal(err)
	}
	reclaimed, err := store.Acquire(first.ExpiresAt, AcquireRequest{Key: "worker-a", Holder: "node-2", Duration: time.Minute})
	if err != nil {
		t.Fatalf("Acquire() at expiry error = %v", err)
	}
	if reclaimed.Holder != "node-2" || reclaimed.Token == first.Token {
		t.Fatalf("Acquire() at expiry = %#v, want a new node-2 lease", reclaimed)
	}
}
