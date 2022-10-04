package asserter

import (
	"fmt"
	"reflect"
	"testing"
)

// Many `Asserter` functions have their tests in other files.
func TestAsserter(t *testing.T) {
	fail := func(string) {}

	/*
		This subtest covers Asserter members:
			.Assert(logical_val) aka .A(...)
			.AssertWithFailMsgAppendix(logical_val, appendix) aka .A(...)
			.AssertEq(lhs,rhs) aka .AE(...)
			.Inc(ind)
			.IncLast()
	*/
	t.Run("{Assertions of `counter` state}", func(t *testing.T) {
		t.Parallel()

		// Successful construction
		a, e := New([]uint{9, 7}, fail)
		if e != nil {
			t.Fatal("e!=nil")
		}

		{
			a.ResetCounter()
			for i, v := range a.Counter() {
				if v != 0 {
					t.Fatalf(
						"After ```a.ResetCounter()``` found value %v at counter index %d .",
						v, i,
					)
				}
			}
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

			a.Inc(1)
			assert_counter_eq([]uint{1, 3})

			a.AssertWithFailMsgAppendix(true, " my not-displayed msg appendix.")
			assert_counter_eq([]uint{1, 4})

			a.AWFMA(false, " my msg appendix.")
			assert_counter_eq([]uint{1, 5})

			a.AssertEq(8.9, 7.1)
			assert_counter_eq([]uint{1, 6})
			a.AssertEq(false, false)
			assert_counter_eq([]uint{1, 7})
			a.AE(8, 8.0)
			assert_counter_eq([]uint{1, 8})

			a.IncLast()
			assert_counter_eq([]uint{1, 9})
		}
	})
}

func ExampleAsserter() {
	a, e := New([]uint{0, 0}, func(s string) { fmt.Println(s) })
	if e != nil {
		panic(fmt.Sprint("e == ", e))
	}

	// Assetions before a.Inc(0).
	{
		// Enumeration tags be like: 0.x
		a.Assert(false)
	}

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

	// The fird pack of assertions.
	{
		a.Inc(0) // Enumeration tags be like: 3.x
		a.AssertWithFailMsgAppendix(true, "my not-displayed msg appendix.")
		a.AWFMA(false, "my msg appendix.")
		a.AssertEq(8.9, 7.1)
		a.AssertEq(false, false)
		a.AE(8, 8.0)
	}

	a.ResetCounter()
	a.AE("abc", nil)

	// Output:
	// 0.0. assertion failed!
	// 1.0. assertion failed!
	// 1.1. assertion failed!
	// 2.1. assertion failed!
	// 3.1. assertion failed!my msg appendix.
	// 3.2. assertion failed! 8.9 != 7.1
	// 3.4. assertion failed! 8 of type `int` is not `reflect.DeepEqual` to 8 of type `float64`
	// 0.0. assertion failed! abc != <nil>
}

// func TestEmptyCounter(t *testing.T) ommited, because the whole implementation is obvious.

// func TestNilFail(t *testing.T) ommited, because implementation is obvious.
