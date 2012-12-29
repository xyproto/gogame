package gogame

type Window struct {
	x int
}

func New() *Window {
	return new(Window)
}
