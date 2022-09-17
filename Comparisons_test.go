package asserter

import (
	"fmt"
	"testing"
	"github.com/huandu/go-clone"
)

func TestAsserter_Eq(t *testing.T){
	for idx,old:= range ExampleAsserters{
		new:= clone.Clone(old).(Asserter)
		if !old.Eq(&new){
			t.Fatalf(
				"TC %d.0 failed! !old.Eq(new)"+
				"\n\told == %v\n\tnew == %v",
				idx, old, new,
			)
		}

		changed:= new
		changed.SetFail( func(string){fmt.Println("kupaa, ha ha")} )
		if old.Eq(&changed){
			t.Errorf(
				"TC %d.1 failed! old.Eq(changed)"+
				"\n\told == %v\n\tchanged == %v",
				idx, old, changed,
			)
		}
		changed.fail= old.fail

		changed.Counter()[0]+= 1 
		if old.Eq(&changed){
			t.Errorf(
				"TC %d.2 failed! old.Eq(changed)"+
				"\n\told == %v\n\tchanged == %v",
				idx, old, changed,
			)
		}
	}
}