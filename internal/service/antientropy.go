package service

// AntiEntropy periodically compares Merkle trees between node pairs and
// triggers repair for divergent key ranges. It builds and diffs the
// trees with meld util/merkle (meld-37n7). The compare-and-repair
// orchestration here is delta's.
type AntiEntropy struct {
	// ports injected at construction
}
