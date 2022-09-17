package asserter

import "testing"

func TestAsserter_GenerateEnumerationTag(t *testing.T) {
	const tcQuantity = 5
	argTab := [tcQuantity][]uint{{1, 4, 7, 0}, {5, 6}, {34}, {0,0,0,0}, {0}}
	expectTab := [tcQuantity]string{"1.4.7.0.", "5.6.", "34.", "0.0.0.0.", "0."}

	a := Asserter{ make([]uint,1), func(string) {} }
	for tcInd, counter := range argTab {
		a.counter= counter
		ex := expectTab[tcInd]
		got := a.GenerateEnumerationTag()

		if ex != got {
			t.Errorf("TC %d failed! expect!=got\n\t"+
				`"%s" != "%s"`,
				tcInd, ex, got)
		}
	}

}

// func TestAsserter_GenerateFailerMsg(t *testing.T) ommited, because the function is obvious and has been comfirmed to work.
