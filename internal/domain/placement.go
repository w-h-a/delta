package domain

// PreferenceList returns the n nodes responsible for the given key, in
// preference order. delta uses meld rendezvous hashing rather than a
// consistent-hashing ring, so there is no ring state to maintain: the
// list is computed from the current node set and the key. Pure and
// deterministic, the same nodes and key yield the same list everywhere.
func PreferenceList(nodes []string, key Key, n int) []string {
	return nil // TODO: rendezvous.AssignN(nodes, key.Name, n) from meld util/rendezvous (meld-gg25)
}
