package domain

// Quorum evaluates whether R/W/N requirements are satisfied.
type Quorum struct {
	N int // total replicas
	R int // required reads
	W int // required writes
}

// ReadSatisfied returns true if enough read responses have been received.
func (q Quorum) ReadSatisfied(responses int) bool {
	return responses >= q.R
}

// WriteSatisfied returns true if enough write acknowledgments have been received.
func (q Quorum) WriteSatisfied(responses int) bool {
	return responses >= q.W
}
