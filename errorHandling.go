package asserter

/*
Description:
	if e != nil {
		panic(e)
	}
*/
func onErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}
