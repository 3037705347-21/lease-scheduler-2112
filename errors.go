package lease

import "errors"

var (
	// ErrInvalidRequest means a request is missing a required field or duration.
	ErrInvalidRequest = errors.New("invalid lease request")
	// ErrHeld means another holder owns an active lease.
	ErrHeld = errors.New("lease is held by another holder")
	// ErrNotFound means a requested lease does not exist.
	ErrNotFound = errors.New("lease not found")
	// ErrTokenMismatch means a caller tried to mutate a newer ownership generation.
	ErrTokenMismatch = errors.New("lease token mismatch")
)
