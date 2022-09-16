package asserter
import (
	"fmt"
	"reflect"
)

/*
	If logical_val==false,
		passes `a.GenerateFailerMsg()+appendix` to `a.fail`.
	Increments last subcounter in `a`.
	Returns `logical_val`.
*/
func (a *Asserter) AssertWithFailMsgAppendix(logical_val bool, appendix string) bool {
	if !logical_val {
		a.fail(a.GenerateFailerMsg()+appendix)
	}

	a.IncLast()
	return logical_val
}

// Alias for `Asserter.AssertWithFailMsgAppendix`.
func (a *Asserter) AWFMA(logical_val bool, appendix string) bool {
	return a.AssertWithFailMsgAppendix(logical_val, appendix)
}

/*
	If logical_val==false,
		passes `a.GenerateFailerMsg()` to `a.fail`.
	Increments last subcounter in `a`.
	Returns `logical_val`.
*/
func (a *Asserter) Assert(logical_val bool) bool {
	if !logical_val {
		a.fail(a.GenerateFailerMsg())
	}

	a.IncLast()
	return logical_val
}

// Alias for `.Assert(logical_val)`.
func (a *Asserter) A(logical_val bool) bool {
	return a.Assert(logical_val)
}

/*
	Like `.Assert`,
	but appends string-printed values of `lhs` and `rhs` in "%v" format
	to basic fail message.
*/
func (a *Asserter) AssertEq(lhs any, rhs any) bool{
	var logical_val= reflect.DeepEqual(lhs, rhs)
	if !logical_val {
		var s1,s2= fmt.Sprint(lhs), fmt.Sprint(rhs)
		var basic_msg string= a.GenerateFailerMsg()

		if s1==s2{
			basic_msg+= fmt.Sprintf(
				" %s of type `%T` is not `reflect.DeepEqual` to %s of type `%T`",
				s1, lhs,
				s2, rhs,
			)
		}else{
			basic_msg+= fmt.Sprintf(" %s != %s", s1, s2)
		}
		a.fail(basic_msg)
	}

	a.IncLast()
	return logical_val
}

// Alias for `.AssertEq(lhs,rhs)`.
func (a *Asserter) AE(lhs any, rhs any) bool{
	return a.AssertEq(lhs,rhs)
}

