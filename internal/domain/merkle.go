package domain

// MerkleTree is a hash tree built from sorted key-hash pairs.
// Used by anti-entropy to detect divergence between replicas.
type MerkleTree struct {
	Root []byte
}

// Build constructs a MerkleTree from sorted key-hash pairs. Pure function.
func Build(keys []Key, hashes [][]byte) MerkleTree {
	return MerkleTree{} // TODO
}

// Diff returns key ranges where two trees disagree. Pure comparison.
func Diff(local, remote MerkleTree) []KeyRange {
	return nil // TODO
}

// KeyRange identifies a contiguous range of keys that need repair.
type KeyRange struct {
	Start Key
	End   Key
}
