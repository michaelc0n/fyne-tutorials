package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window

	win := app.NewWindow("My New Title") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// checkbox widget
	checkBox := widget.NewCheck("True/False, select checkbox", func(b bool) {
		fmt.Printf("my checkbox value %t\n", b)
		// when checkbox selected (true), exit fyne app
		if b {
			app.Quit()
		}
	})
	// set content checkBox
	win.SetContent(checkBox)
	win.ShowAndRun() // run app
}
