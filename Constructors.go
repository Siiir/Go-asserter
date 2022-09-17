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
	Equivalent to:
		1. New( make([]uint,1), fail )
		2. NewReseted( 1, fail )
*/
func NewTiny(
	fail func(failerMsg string),
) (a Asserter, e error) {
	a.counter= make([]uint,1)
	e = a.SetFail(fail)
	return
}


// Equivalent to: New( make([]uint,counterLen), fail )
func NewReseted(counterLen uint, fail func(string)) (Asserter, error){
	return New( make([]uint,counterLen), fail )
}

/*
	Similar to: NewReseter(uint(len(a.counter)), a.fail) ,
	but because it is bound to succeed it doesn't return `error` .
*/
func (a *Asserter) NewReseted() Asserter {
	return Asserter{
		make(
			[]uint,
			uint( len(a.counter) ),
		),
		a.fail,
	}
}

/*
	Returns semantic copy of `*a`.
*/
func (a *Asserter) Clone() Asserter {
	new := a.NewReseted()
	copy(new.counter, a.counter)
	return new
}

