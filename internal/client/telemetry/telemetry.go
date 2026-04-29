// Package telemetry is the port interface for observability.
// Wide events for coordinator requests, anti-entropy cycles,
// sync sessions, and CRDT merges.
package telemetry

import "context"

// Telemetry emits wide events for store operations.
type Telemetry interface {
	EmitCoordinatorRequest(ctx context.Context, event CoordinatorEvent) error
	EmitAntiEntropyCycle(ctx context.Context, event AntiEntropyEvent) error
	EmitSyncSession(ctx context.Context, event SyncEvent) error
}

type CoordinatorEvent struct {
	Key            string
	Operation      string // "get", "put", "delete"
	PreferenceList []string
	ResponseCount  int
	QuorumMet      bool
}

type AntiEntropyEvent struct {
	LocalNode    string
	RemoteNode   string
	KeysCompared int
	KeysRepaired int
}

type SyncEvent struct {
	ClientID      string
	EntriesSent   int
	EntriesRecvd  int
	BytesSent     int64
	BytesRecvd    int64
	Converged     bool
}
