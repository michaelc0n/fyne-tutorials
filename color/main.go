package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window

	win := app.NewWindow("Canvas Text") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// set color, R=Red, G=Green, B=Blue, A=alpha for opacity, range=0-to-255
	colorx := color.NRGBA{R: 0, G: 255, B: 0, A: 100}

	// text widget in canvas
	newText := canvas.NewText(">>>here is my green text<<<", colorx) // first label/caption, second color value
	// center text
	content := container.New(layout.NewCenterLayout(), newText)

	// set content
	win.SetContent(content)
	win.ShowAndRun() // run app
}
