package asserter

import (
	"fmt"
	"strings"
)

type asserter struct {
	counter []uint
	fail    func(basic_msg string)
}

func New(
	counter []uint,
	fail func(basic_msg string),
) (a asserter, e error) {
	e = a.SetCounter(counter)
	if e != nil {
		return
	}
	e = a.SetFail(fail)
	return
}

func (a *asserter) SetCounter(counter []uint) error {
	if len(counter) == 0 {
		return EmptyCounter{}
	}
	a.counter = counter
	return nil
}

func (a *asserter) SetFail(fail func(basic_msg string)) error {
	if fail == nil {
		return (NilFail{})
	}
	a.fail = fail
	return nil
}

func (a *asserter) Assert(logical_val bool) bool {
	a.IncLast()

	if !logical_val {
		var sb strings.Builder

		sb.WriteByte('\n')
		for _, c := range a.counter {
			sb.WriteString(fmt.Sprint(c))
			sb.WriteByte('.')
		}
		sb.WriteString(" assertion failed!")

		a.fail(sb.String())
	}

	return logical_val
}

// Alias for `.assert(logical_val)`.
func (a *asserter) A(logical_val bool) bool {
	return a.Assert(logical_val)
}

func (a *asserter) Inc(ind uint) (indexNotPresentInCounter bool) {
	if uint(len(a.counter)) <= ind {
		return true
	}
	a.counter[ind]++

	for ind++; ind < uint(len(a.counter)); ind++ {
		a.counter[ind] = 0
	}
	return
}
func (a *asserter) IncLast() {
	a.Inc(uint(len(a.counter)) - 1)
}

func (a *asserter) Counter() []uint {
	return a.counter
}

func (a *asserter) Fail() func(basic_msg string) {
	return a.fail
}
