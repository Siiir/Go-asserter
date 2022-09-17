package asserter

// Getters

// Returns a.counter
func (a *Asserter) Counter() []uint {
	return a.counter
}

// Returns a.fail
func (a *Asserter) Fail() func(failerMsg string) {
	return a.fail
}

// Setters

/*
	Description:

	if len(counter) == 0 {
		return EmptyCounterError{}
	}
	a.counter = counter
	return nil
*/
func (a *Asserter) SetCounter(counter []uint) error {
	if len(counter) == 0 {
		return EmptyCounterError{}
	}
	a.counter = counter
	return nil
}

/*
	Description:

	if fail == nil {
		return (NilFailError{})
	}
	a.fail = fail
	return nil
*/
func (a *Asserter) SetFail(fail func(failerMsg string)) error {
	if fail == nil {
		return NilFailError{}
	}
	a.fail = fail
	return nil
}
