package domain

// Replica holds the state for a single key on a single node.
// Value plus CRDT metadata for merge.
type Replica struct {
	Key     Key
	Value   []byte
	Version []byte // serialized version vector from meld
	Deleted bool   // tombstone
}

// Merge combines two replicas using CRDT semantics.
// Commutative, associative, idempotent.
func Merge(a, b Replica) Replica {
	return Replica{} // TODO: delegate to meld CRDT merge, compare version vectors
}
