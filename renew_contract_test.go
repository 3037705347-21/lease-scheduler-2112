package lease

import (
	"testing"
	"time"
)

func TestRenewUsesTheRenewalTimeAsItsDeadlineBase(t *testing.T) {
	store := NewStore()
	started := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	lease, err := store.Acquire(started, AcquireRequest{Key: "worker-c", Holder: "node-1", Duration: time.Minute})
	if err != nil { t.Fatal(err) }
	renewedAt := started.Add(30 * time.Second)
	renewed, err := store.Renew(renewedAt, RenewRequest{Key: "worker-c", Holder: "node-1", Token: lease.Token, Duration: 2 * time.Minute})
	if err != nil { t.Fatal(err) }
	if want := renewedAt.Add(2 * time.Minute); !renewed.ExpiresAt.Equal(want) {
		t.Fatalf("Renew() expiry = %v, want %v", renewed.ExpiresAt, want)
	}
}
