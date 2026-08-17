// Package lease provides an in-memory scheduler for short-lived ownership leases.
package lease

import "time"

// Lease represents a claimed resource and the instant at which its claim expires.
type Lease struct {
	Key       string
	Holder    string
	Token     uint64
	ExpiresAt time.Time
}

// Active reports whether the lease still belongs to its holder at now.
func (l Lease) Active(now time.Time) bool {
	return l.Key != "" && l.Holder != "" && now.Before(l.ExpiresAt)
}

// Expired reports whether a lease is available for a new holder.
func (l Lease) Expired(now time.Time) bool {
	return !l.Active(now)
}
