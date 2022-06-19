package asserter

import (
	"fmt"
	"strings"
)

// Generates enumeration tag for upcoming assertion case in form
// "a.b.c....z", where a = a.counter[0], b = a.counter[1], ..., z = a.counter[len(a.counter)-1]
// Example tags:  "0.0.0", "1.3", "5", "1.1.2"
func (a *Asserter) GenerateEnumerationTag() string {
	var sb strings.Builder

	sb.WriteByte('\n')
	for _, c := range a.counter {
		sb.WriteString(fmt.Sprint(c))
		sb.WriteByte('.')
	}

	return sb.String()
}

// Returns a.GenerateEnumerationTag() + " assertion failed!"
func (a *Asserter) GenerateFailerMsg() (failerMsg string) {
	return a.GenerateEnumerationTag() + " assertion failed!"
}

/*
	If logical_val==false,
		passes `a.GenerateFailerMsg()` to `a.fail`.
	Increments last subcounter in `a`.
	Returns `logical_val`.
*/
func (a *Asserter) Assert(logical_val bool) bool {
	if !logical_val {
		a.fail(a.GenerateFailerMsg())
	}

	a.IncLast()
	return logical_val
}

// Alias for `.assert(logical_val)`.
func (a *Asserter) A(logical_val bool) bool {
	return a.Assert(logical_val)
}

/*
func (a *Asserter) AssertEq
func (a *Asserter) AE
*/
