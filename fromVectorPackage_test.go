package asserter

/*
Returns true if both slices satisfy any of these conditions:
	1. are nil
	2. slice an empty array
	3. are slicing the same part of the same array
*/
func heapSame[T any](slA, slB []T) (equal bool) {
	// Checking 2 integer components.
	if cap(slA) != cap(slB) || len(slA) != len(slB) {
		return false
	}
	// Checking pointer
	if cap(slA) == 0 {
		if b1, b2 := slA == nil, slB == nil; b1 || b2 {
			return b1 && b2
		}
		return true
	}
	return &slA[:1][0] == &slB[:1][0]
}
