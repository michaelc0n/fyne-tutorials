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

	// button widget, first value is button name
	// second value is any function
	btn := widget.NewButton("Button Name", func() {
		// test button, print to `stdout`
		fmt.Println("print button to stdout")
	})

	win.SetContent(btn) // using button
	win.ShowAndRun()    // run app
}
