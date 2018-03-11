package go_logstock

type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}
