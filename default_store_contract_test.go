package lease

import "testing"

func TestNewSchedulerCreatesAStoreWhenNoneIsProvided(t *testing.T) {
	scheduler := NewScheduler(nil, nil)
	if active := scheduler.Active(); len(active) != 0 {
		t.Fatalf("Active() = %#v, want empty", active)
	}
}
