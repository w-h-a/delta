// Package persister is the port interface for per-node KV storage.
package persister

import (
	"context"

	"github.com/w-h-a/delta/internal/domain"
)

// Persister stores and retrieves replicas on a single node.
type Persister interface {
	Put(ctx context.Context, replica domain.Replica) error
	Get(ctx context.Context, key domain.Key) (domain.Replica, error)
	Delete(ctx context.Context, key domain.Key) error
	Keys(ctx context.Context) ([]domain.Key, error)
}
