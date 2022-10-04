package asserter

import (
	"fmt"
	"strings"
)

/*
	Generates enumeration tag for upcoming assertion case in form "α.β.γ....ζ",

	where α = `a.counter[0]`, β = `a.counter[1]`, ..., ζ = `a.counter[len(a.counter)-1]` .

	Example tags:  "0.0.0", "1.3", "5", "1.1.2"
*/
func (a *Asserter) GenerateEnumerationTag() string {
	var sb strings.Builder

	for _, c := range a.counter {
		sb.WriteString(fmt.Sprint(c))
		sb.WriteByte('.')
	}

	return sb.String()
}

/*
	Description:
		return a.GenerateEnumerationTag() + " assertion failed!"
*/
func (a *Asserter) GenerateFailerMsg() (failerMsg string) {
	return a.GenerateEnumerationTag() + " assertion failed!"
}
