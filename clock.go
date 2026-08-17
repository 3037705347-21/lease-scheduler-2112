package lease

import "time"

// Clock supplies the time source used by Scheduler.
type Clock interface {
	Now() time.Time
}

// WallClock reads the local process clock.
type WallClock struct{}

// Now returns the current wall-clock time.
func (WallClock) Now() time.Time {
	return time.Now()
}

// FixedClock is useful for deterministic integrations and tests.
type FixedClock struct {
	Time time.Time
}

// Now returns the configured time.
func (c FixedClock) Now() time.Time {
	return c.Time
}
