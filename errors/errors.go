package errors

type ModuleNotFound struct {
	Name string
}

func (e ModuleNotFound) Error() string {
	return "module not found: " + e.Name
}
