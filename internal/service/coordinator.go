// Package service contains the controller layer.
package service

// Coordinator routes requests to the preference list and gathers quorum.
// Thin controller: looks up ring, fans out to replicas, evaluates quorum, returns.
type Coordinator struct {
	// ports injected at construction
}
