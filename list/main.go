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

	win := app.NewWindow("List View") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// list view - three arguments:
	// item count, (three items in list)
	// widget, label widget
	// update widget data
	list := widget.NewList(
		func() int { return 3 }, // list contains 3 items
		func() fyne.CanvasObject { return widget.NewLabel("label here") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			// update widget data
			co.(*widget.Label).SetText("here is my text")
		},
	)

	// setup data on screen
	win.SetContent(list)
	win.ShowAndRun() // run app
}
