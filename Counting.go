package asserter

/*
	If possible, does a.counter[ind]++ and returns false.

	Otherwise returns true.
*/
func (a *Asserter) Inc(ind uint) (indexNotPresentInCounter bool) {
	if uint(len(a.counter)) <= ind {
		return true
	}
	a.counter[ind]++

	for ind++; ind < uint(len(a.counter)); ind++ {
		a.counter[ind] = 0
	}
	return
}

/*
	Description:
		a.counter[len(a.counter)-1]++
*/
func (a *Asserter) IncLast() {
	a.counter[len(a.counter)-1]++
}

// Zeroes the internal counter.
func (a *Asserter) ResetCounter() {
	for i := range a.counter {
		a.counter[i] = 0
	}
}
