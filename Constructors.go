package asserter

/*
	Creates a new Asserter.
	Passes its argument to corresponding field setters
	stopping at first error caught and returning it.
	[No error caught <==> `e`==nil] ==>
		Construction was successful. Returned Asserter object is valid.
*/
func New(
	counter []uint,
	fail func(failerMsg string),
) (a Asserter, e error) {
	e = a.SetCounter(counter)
	if e != nil {
		return
	}
	e = a.SetFail(fail)
	return
}

/*
func NewReseted(counterLen uint, fail func(string)) Asserter{
	return Asserter{make([]uint, counterLen), fail}
}

func (a *Asserter) NewReseted() Asserter {
	return NewReseter(uint(len(a.counter)), a.fail)
}

func (a *Asserter) Clone() Asserter {
	new := a.NewReseted()
	copy(new.counter, a.counter)
	return new
}
*/
