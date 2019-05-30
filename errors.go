package example

type NotFound struct{}

func (e NotFound) Error() string {
	return "Resource not found"
}
