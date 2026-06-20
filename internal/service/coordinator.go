// Package service contains the controller layer.
package service

// Coordinator routes requests to the preference list and gathers quorum.
// Thin controller: computes the preference list via rendezvous, fans out
// to replicas, evaluates quorum, returns.
type Coordinator struct {
	// ports injected at construction
}
