package asserter

import (
	"reflect"
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
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
			if a, e := New(make([]uint,0), fail); !(reflect.DeepEqual(a, Asserter{}) && e == EmptyCounterError{}) {
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

// TestNewTiny ommited, because `NewTiny` is simple and has been checked visually.

// TestNewReseted ommited, because `NewReseted` is only a variation of `New`.

const ExampleAssertersQuantity= 4;
var ExampleAsserters= [ExampleAssertersQuantity]Asserter{
	Asserter{[]uint{4},func(string){}},
	Asserter{[]uint{0,6,2,},func(_ string){}},
	Asserter{[]uint{4,6,2,567,34,23},func(s string){println(s)}},
	Asserter{[]uint{4,6,2,567,0,0,0,0,0,0},func(s string){fmt.Println(s)}},
}

func TestAsserter_NewReseted(t *testing.T){
	// Most functionality is already tested.
	// We only need to make sure that condition
	// ```len(a.Counter()) == len(a.NewReseted.Counter())```
	// is always met.

	for idx, old := range ExampleAsserters{
		new:= old.NewReseted()
		if len(old.Counter()) != len(new.Counter()){
			t.Errorf(
				"TC %d failed!"+
				" len(old.Counter()) != len(new.Counter())"+
				"\n\told == %v\n\tnew == %v",
				idx, old, new,
			)
		}
	}
}

func TestAsserter_Clone(t *testing.T){
	for idx, old := range ExampleAsserters{
		new:= old.Clone()
		if !old.Eq(&new){
			t.Errorf(
				"\nTC %d failed!"+
				" old != new"+
				"\n\told == %v\n\tnew == %v",
				idx, old, new,
			)
		}
	}
}