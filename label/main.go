package main

import (
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

	// first widget
	newLabel := widget.NewLabel("New Label!")
	win.SetContent(newLabel)

	win.ShowAndRun() // run app
}
