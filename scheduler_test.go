package lease

import (
	"errors"
	"testing"
	"time"
)

func TestSchedulerLifecycle(t *testing.T) {
	now := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	scheduler := NewScheduler(NewStore(), FixedClock{Time: now})

	lease, err := scheduler.Acquire("worker-a", "node-1", time.Minute)
	if err != nil {
		t.Fatalf("Acquire() error = %v", err)
	}
	if lease.Token == 0 {
		t.Fatal("Acquire() returned an empty token")
	}

	renewed, err := scheduler.Renew("worker-a", "node-1", lease.Token, 2*time.Minute)
	if err != nil {
		t.Fatalf("Renew() error = %v", err)
	}
	if !renewed.ExpiresAt.Equal(now.Add(2 * time.Minute)) {
		t.Fatalf("Renew() expiry = %v", renewed.ExpiresAt)
	}

	if err := scheduler.Release("worker-a", "node-1", renewed.Token); err != nil {
		t.Fatalf("Release() error = %v", err)
	}
	if active := scheduler.Active(); len(active) != 0 {
		t.Fatalf("Active() = %v, want empty", active)
	}
}

func TestSchedulerRejectsConflictingHolder(t *testing.T) {
	now := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	scheduler := NewScheduler(NewStore(), FixedClock{Time: now})
	if _, err := scheduler.Acquire("worker-a", "node-1", time.Minute); err != nil {
		t.Fatal(err)
	}
	if _, err := scheduler.Acquire("worker-a", "node-2", time.Minute); !errors.Is(err, ErrHeld) {
		t.Fatalf("Acquire() error = %v, want ErrHeld", err)
	}
}
