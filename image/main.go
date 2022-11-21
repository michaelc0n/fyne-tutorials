package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window

	win := app.NewWindow("Image") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// image widget
	img := canvas.NewImageFromFile("./game_on.jpg")

	// set content
	win.SetContent(img)
	win.ShowAndRun() // run app
}
