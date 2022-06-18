package asserter

type EmptyCounter struct{}

func (EmptyCounter) Error() string { return "len(counter) == 0" }

type NilFail struct{}

func (NilFail) Error() string { return "fail == nil" }
