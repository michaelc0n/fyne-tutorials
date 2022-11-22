package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window
	win := app.NewWindow("Layout NewVBox and NewHBox") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// new button
	btn := widget.NewButton("click me", func() {})

	label := widget.NewLabel("here is my text")

	// // NewVBox and NewHBox
	// box := container.NewVBox(btn, label)
	box := container.NewHBox(btn, label)

	// set content
	win.SetContent(box)
	win.ShowAndRun() // run app
}
