package gogame

import (
	"fmt"
)

type App struct {
	x int
}

func New() *App {
	return new(App)
}

func (app *App) Run() {
	fmt.Println("run")
}
