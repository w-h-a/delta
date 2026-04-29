package domain

// Ring is a consistent hashing ring with virtual nodes.
// PreferenceList returns the N nodes responsible for a key.
// Deterministic: same ring + same key = same preference list on every node.
type Ring struct {
	Nodes        []string
	VirtualNodes int
}

// PreferenceList returns the n nodes responsible for the given key.
func (r Ring) PreferenceList(key Key, n int) []string {
	return nil // TODO: walk ring from key.Position, collect n distinct physical nodes
}
