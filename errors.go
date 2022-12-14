package zc

type UnsupportedError struct {
	Name string
}

func (e UnsupportedError) Error() string {
	return "unsupported operation: " + e.Name
}

type EmptyStackError struct {
	Name string
}

func (e EmptyStackError) Error() string {
	return e.Name + ": empty stack"
}
