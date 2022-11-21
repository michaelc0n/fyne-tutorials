package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window

	win := app.NewWindow("My New Title") // use any title for app

	win.ShowAndRun() // run app
}
