package domain

// Key is a store key hashed to a position on the consistent hashing ring.
type Key struct {
	Name     string
	Position uint64 // hash of Name, determines ring placement
}
