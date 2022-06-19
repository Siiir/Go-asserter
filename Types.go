package asserter

/*
	Is a type for making & enumerating dynamic assertions, also automatically logging failers.

	`counter` slice contains subcounters of assertions.
	Must have at least one subcounter automatically used (incremented) by Asserter.Assert method).
	If assertions come in one series, then
		one subcounter is enough, and so len(`counter`)==1.
	If assertions come in (nested) packs, then:
		1. counter[0] enumerates top-level packs and top-level packs can be switched with Asserter.Inc(0).
		2. counter[i] where i>0 enumerates lower-level packs and they can be switched with Asserter.Inc(i).
	`fail` member.

	`fail` logs all failers with `failerMsg` provided by asserter object.
	Most popular `fail` clousures:
		func(failerMsg string){t.Error(failerMsg)}, where t is *testing.T
		func(failerMsg string){print(failerMsg)}
		func(failerMsg string){logFile.WriteString(failerMsg)}, where logFile is a text stream
*/
type Asserter struct {
	counter []uint
	fail    func(failerMsg string)
}

// Errors with their implementation.

// Denotes that length of `counter` variable equals 0,
// which means it cannot be(come) asseter.counter & be used for assertion counting.
type EmptyCounterError struct{}

// Returns "len(counter) == 0"
func (EmptyCounterError) Error() string { return "len(counter) == 0" }

// Denotes that `fail` variable equals nil.
// which means it cannot be(come) Asseter.fail & be used to log failing assertions.
type NilFailError struct{}

// Returns "fail == nil"
func (NilFailError) Error() string { return "fail == nil" }
