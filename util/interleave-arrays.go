package util

// InterleaveSlices helper function will interleave and combine two sequences.
func InterleaveSlices[T any](a []T, b []T) []T {
	var arr []T
	observed := a
	other := b
	var temp []T

	for len(observed) > 0 {
		arr = append(arr, observed[0])
		observed = append(observed[:0], observed[1:]...)
		temp = observed
		observed = other
		other = temp
	}

	if len(other) > 0 {
		arr = append(arr, other...)
	}
	return arr
}
