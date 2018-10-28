package main

import (
	"fmt"
	"github.com/indeedhat/ButtonPad/DesktopApp/gui"
)

func main() {
	fmt.Println("starting")
	app := gui.NewUiController(func(conf bool) {
		fmt.Println("confirm clicked")
	}, func(layer, macro int, typ byte, label string, payload []byte) {
		fmt.Println("layer clicked")
	})

	fmt.Println("App added")

	app.Start()
	fmt.Println("closing app")
}