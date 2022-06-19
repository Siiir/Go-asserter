package asserter

import (
	"fmt"
	"reflect"
	"testing"
)

// `Asserter` has its test in other files.
func TestAsserter(t *testing.T) {
	fail := func(string) {}

	/*
		This subtest covers Asserter members:
			.counter
			.Counter()
			.SetCounter(counter)
			.Assert(logical_val) aka .A(...)
			.Inc(ind)
			.IncLast()
	*/
	t.Run("{Assertions of `counter` state}", func(t *testing.T) {
		t.Parallel()

		// Successful construction
		a, e := New([]uint{0, 0}, fail)
		if e != nil {
			t.Fatal("e!=nil")
		}

		// auxialiary function
		assert_counter_eq := func(val []uint) {
			if !reflect.DeepEqual(a.counter, val) {
				t.Fatalf("!reflect.DeepEqual(a.counter, val);"+
					"\n\ta.counter = %v\n\tval = %v",
					a.counter, val)
			}
		}

		// Checks themselves
		{
			a.Assert(true)
			a.Assert(true)
			assert_counter_eq([]uint{0, 2})

			a.A(false)
			assert_counter_eq([]uint{0, 3})

			a.Inc(0)
			assert_counter_eq([]uint{1, 0})

			a.A(true)
			assert_counter_eq([]uint{1, 1})

			a.IncLast()
			assert_counter_eq([]uint{1, 2})
		}

	})
}

func Example_asserter() {
	a, e := New([]uint{0, 0}, func(s string) { fmt.Print(s) })
	if e != nil {
		panic("e!=nil")
	}

	// Assetions before a.Inc(0).
	// Enumeration tags be like: 0.x
	a.Assert(false)

	// The first pack of assertions.
	{
		a.Inc(0) // Enumeration tags be like: 1.x

		a.Assert(2 == 3)
		a.Assert(7+8 < 15)
		a.A(true)
		a.A(1 < 2)
	}

	// The second pack of assertions.
	{
		a.Inc(0) // Enumeration tags be like: 2.x
		a.A(6.0 == 6e0)
		a.Assert(false)
	}

	// Output:
	// 0.0. assertion failed!
	// 1.0. assertion failed!
	// 1.1. assertion failed!
	// 2.1. assertion failed!

}

// func TestEmptyCounter(t *testing.T) ommited, because the whole implementation is obvious.

// func TestNilFail(t *testing.T) ommited, because implementation is obvious.
