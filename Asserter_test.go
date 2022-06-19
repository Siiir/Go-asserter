package asserter

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_asserter(t *testing.T) {
	fail := func(string) {}

	t.Run("New", func(t *testing.T) {
		t.Parallel()
		counter := make([]uint, 40)

		t.Run("->unsuccessful", func(t *testing.T) {
			// Misconstructions.
			{
				if a, e := New(nil, nil); !(reflect.DeepEqual(a, asserter{}) && e == EmptyCounter{}) {
					t.Fatalf("\n!(reflect.DeepEqual(a, asserter{}) && e == EmptyCounter{})"+
						"\n\ta==%v\n\te==%v", a, e)
				}

				if a, e := New(nil, fail); !(reflect.DeepEqual(a, asserter{}) && e == EmptyCounter{}) {
					t.Fatalf("\n!(reflect.DeepEqual(a, asserter{}) && e == EmptyCounter{})"+
						"\n\ta==%v\n\te==%v", a, e)
				}

				if a, e := New(counter, nil); !(reflect.DeepEqual(a, asserter{counter: counter}) && e == NilFail{}) {
					t.Fatalf("\n!(reflect.DeepEqual(a, asserter{counter: %v}) && e == NilFail{})"+
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

				if *(*uintptr)(unsafe.Pointer(&fail)) != *(*uintptr)(unsafe.Pointer(&a.fail)) {
					t.Fatal("fail != a.fail")
				}
			}
		})
	})

	/*
		This subtest covers asserter members:
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
		var a asserter
		{
			counter := []uint{2, 3}
			var e error
			a, e = New(counter, fail)
			if e != nil {
				t.Fatal("e!=nil")
			}
		}

		// auxialiary function
		assert_counter_eq := func(val []uint) {
			if !reflect.DeepEqual(a.counter, val) {
				t.Fatalf("!reflect.DeepEqual(a.counter, %v);\n\tcounter=%v",
					val, a.counter)
			}
		}

		// Checks themselves
		{
			a.SetCounter([]uint{0, 1})
			assert_counter_eq([]uint{0, 1})

			a.counter[0] = 73
			assert_counter_eq([]uint{73, 1})

			a.Counter()[0] = 0
			assert_counter_eq([]uint{0, 1})

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

	// t.Run("[Set]Fail",...) ommited, because both functions is obvious.
}

func Example_asserter() {
	a, e := New([]uint{0, 0}, func(s string) { fmt.Print(s) })
	if e != nil {
		panic("e!=nil")
	}

	a.Inc(0)
	a.Assert(2 == 3)
	a.Assert(7+8 < 15)
	a.A(true)
	a.A(1 < 2)

	a.Inc(0)
	a.A(6.0 == 6e0)
	a.Assert(false)

	// Output:
	// 1.1. assertion failed!
	// 1.2. assertion failed!
	// 2.2. assertion failed!

}
