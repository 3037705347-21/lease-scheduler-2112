# Lease Scheduler

Lease Scheduler is a Go library for coordinating short-lived ownership of named
resources inside one process. It is intended for worker pools and service
components that need a small, deterministic lease registry without a networked
dependency.

## Packages

- `lease.go` defines lease state and expiration rules.
- `request.go` validates acquire and renew inputs.
- `store.go` serializes state transitions and exposes stable snapshots.
- `clock.go` provides wall-clock and deterministic time sources.
- `scheduler.go` offers the public coordination API.

## Use

```go
scheduler := lease.NewScheduler(nil, nil)
claim, err := scheduler.Acquire("worker-a", "node-1", time.Minute)
```

Build and test the module with:

```text
go build ./...
go test ./...
```

The library has no environment variables or external service dependencies.
