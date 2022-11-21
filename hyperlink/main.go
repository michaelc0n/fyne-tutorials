package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window

	win := app.NewWindow("Hyperlink") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(100, 100)) // first width, then height

	// create url
	url, _ := url.Parse("https://apod.nasa.gov/apod/astropix.html")

	// hyperlink widget
	hyperlink := widget.NewHyperlink("Nasa Picture Of The Day!", url) // first value is label, second is url

	// set content
	win.SetContent(hyperlink)
	win.ShowAndRun() // run app
}
