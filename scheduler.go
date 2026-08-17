package lease

import "time"

// Scheduler combines a Store with a time source to expose lease operations.
type Scheduler struct {
	store *Store
	clock Clock
}

// NewScheduler constructs a scheduler. A nil clock uses the process wall clock.
func NewScheduler(store *Store, clock Clock) *Scheduler {
	if clock == nil {
		clock = WallClock{}
	}
	return &Scheduler{store: store, clock: clock}
}

// Acquire claims a resource for a holder.
func (s *Scheduler) Acquire(key, holder string, duration time.Duration) (Lease, error) {
	return s.store.Acquire(s.clock.Now(), AcquireRequest{Key: key, Holder: holder, Duration: duration})
}

// Renew extends a lease owned by holder and token.
func (s *Scheduler) Renew(key, holder string, token uint64, duration time.Duration) (Lease, error) {
	return s.store.Renew(s.clock.Now(), RenewRequest{Key: key, Holder: holder, Token: token, Duration: duration})
}

// Release relinquishes a lease owned by holder and token.
func (s *Scheduler) Release(key, holder string, token uint64) error {
	return s.store.Release(s.clock.Now(), key, holder, token)
}

// Active returns currently valid leases in stable key order.
func (s *Scheduler) Active() []Lease {
	return s.store.Snapshot(s.clock.Now())
}
