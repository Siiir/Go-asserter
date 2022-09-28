package asserter

import "fmt"

func (lhs *Asserter) Eq(rhs *Asserter) bool {
	var lhs_c, rhs_c = lhs.counter, rhs.counter
	if len(lhs_c) != len(rhs_c) {
		return false
	}
	for i := range lhs_c {
		if lhs_c[i] != rhs_c[i] {
			return false
		}
	}
	return fmt.Sprintf("%p", lhs.fail) == fmt.Sprintf("%p", rhs.fail)
}
