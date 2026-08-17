package lease

import (
	"sort"
	"sync"
	"time"
)

// Store owns lease state and serializes all transitions through one mutex.
type Store struct {
	mu     sync.Mutex
	leases map[string]Lease
	next   uint64
}

// NewStore creates an empty lease registry.
func NewStore() *Store {
	return &Store{leases: make(map[string]Lease)}
}

// Acquire creates a lease or returns the caller's still-active lease unchanged.
func (s *Store) Acquire(now time.Time, request AcquireRequest) (Lease, error) {
	if !request.Valid() {
		return Lease{}, ErrInvalidRequest
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	current, found := s.leases[request.Key]
	if found && current.Active(now) {
		if current.Holder == request.Holder {
			return current, nil
		}
		return Lease{}, ErrHeld
	}

	s.next++
	created := Lease{
		Key:       request.Key,
		Holder:    request.Holder,
		Token:     s.next,
		ExpiresAt: now.Add(request.Duration),
	}
	s.leases[created.Key] = created
	return created, nil
}

// Renew extends an active lease while preserving its ownership generation.
func (s *Store) Renew(now time.Time, request RenewRequest) (Lease, error) {
	if !request.Valid() {
		return Lease{}, ErrInvalidRequest
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	current, found := s.leases[request.Key]
	if !found || current.Expired(now) {
		return Lease{}, ErrNotFound
	}
	if current.Holder != request.Holder || current.Token != request.Token {
		return Lease{}, ErrTokenMismatch
	}

	// Extend from the renewal moment, not the prior expiry, so a renew that
	// happens mid-lease resets the deadline to now + duration.
	current.ExpiresAt = now.Add(request.Duration)
	s.leases[current.Key] = current
	return current, nil
}

// Release removes an active lease after checking its ownership generation.
func (s *Store) Release(now time.Time, key, holder string, token uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	current, found := s.leases[key]
	if !found || current.Expired(now) {
		return ErrNotFound
	}
	if current.Holder != holder || current.Token != token {
		return ErrTokenMismatch
	}
	delete(s.leases, key)
	return nil
}

// Snapshot returns active leases sorted by key so callers get deterministic output.
func (s *Store) Snapshot(now time.Time) []Lease {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]Lease, 0, len(s.leases))
	for key, current := range s.leases {
		if current.Expired(now) {
			delete(s.leases, key)
			continue
		}
		result = append(result, current)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Key < result[j].Key })
	return result
}
