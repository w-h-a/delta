package domain

// Key is a store key. Rendezvous hashing scores (node, Name) pairs
// directly, so there is no precomputed ring position.
type Key struct {
	Name string
}
