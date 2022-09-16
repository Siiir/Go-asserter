package asserter

import (
	"reflect"
	"testing"
)

func TestAsserter_New(t *testing.T) {
	fail := func(string) {}
	counter := make([]uint, 40)

	t.Run("->unsuccessful", func(t *testing.T) {
		// Misconstructions.
		{
			if a, e := New(nil, nil); !(reflect.DeepEqual(a, Asserter{}) && e == EmptyCounterError{}) {
				t.Fatalf("\n!(reflect.DeepEqual(a, Asserter{}) && e == EmptyCounterError{})"+
					"\n\ta==%v\n\te==%v", a, e)
			}

			if a, e := New(nil, fail); !(reflect.DeepEqual(a, Asserter{}) && e == EmptyCounterError{}) {
				t.Fatalf("\n!(reflect.DeepEqual(a, Asserter{}) && e == EmptyCounterError{})"+
					"\n\ta==%v\n\te==%v", a, e)
			}

			if a, e := New(counter, nil); !(reflect.DeepEqual(a, Asserter{counter: counter}) && e == NilFailError{}) {
				t.Fatalf("\n!(reflect.DeepEqual(a, Asserter{counter: %v}) && e == NilFailError{})"+
					"\n\ta==%v\n\te==%v", counter, a, e)
			}
		}
	})

	t.Run("->successful", func(t *testing.T) {
		// Successful construction.
		a, e := New(counter, fail)
		if e != nil {
			t.Fatalf("e != nil"+
				"e == %v", e)
		}
		// First use
		a.Assert(false)

		// Construction checks.
		{
			if !heapSame(counter, a.Counter()) {
				t.Fatalf("\n!heapSame(counter, a.Counter())"+
					"\n\tcounter == %v\n\ta.Counter() == %v",
					counter, a.Counter())
			}

			// I don't know how to check wheather a.fail==fail .
		}
	})

}
