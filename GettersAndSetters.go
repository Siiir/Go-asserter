package asserter

// Getters

// Returns private member `a.counter`.
func (a *Asserter) Counter() []uint {
	return a.counter
}

// Returns private member `a.fail`.
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
	Mad version of `SetCounter(counter)`.

	Panic on error − not returning it.
*/
func (a *Asserter) PSetCounter(counter []uint) {
	e := a.SetCounter(counter)
	onErrPanic(e)
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

/*
	Mad version of `SetFail(fail)`.
	Panic on error − not returning it.
*/
func (a *Asserter) PSetFail(fail func(failerMsg string)) {
	e := a.SetFail(fail)
	onErrPanic(e)
}
