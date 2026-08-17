package lease

import "time"

// AcquireRequest contains the resource, owner, and requested lifetime for a claim.
type AcquireRequest struct {
	Key      string
	Holder   string
	Duration time.Duration
}

// Valid reports whether the request can create a meaningful lease.
func (r AcquireRequest) Valid() bool {
	return r.Key != "" && r.Holder != "" && r.Duration > 0
}

// RenewRequest extends an existing claim only when its ownership token matches.
type RenewRequest struct {
	Key      string
	Holder   string
	Token    uint64
	Duration time.Duration
}

// Valid reports whether the request can safely renew an existing claim.
func (r RenewRequest) Valid() bool {
	return r.Key != "" && r.Holder != "" && r.Token > 0 && r.Duration > 0
}
