package objects

type Callable interface {
	// Call should take an arbitrary number of arguments
	// and returns a return value and/or an error,
	// which the VM will consider as a run-time error.
	Call(args ...Object) (ret Object, err error)
}
