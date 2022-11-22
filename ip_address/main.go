package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// create a new app
	app := app.New()

	// create a new window
	win := app.NewWindow("What is my IP address?") // use any title for app

	// resize fyne app window
	win.Resize(fyne.NewSize(400, 400)) // first width, then height

	// UI - label/button setup
	labelTitle := widget.NewLabel("What is my IP address?")
	labelIP := widget.NewLabel("Your IP is...")
	labelValue := widget.NewLabel("...")
	btn := widget.NewButton("Run", func() {
		// logic
		labelValue.Text = myIP()
		labelValue.Refresh() // update labelValue(...) with received data (ip)
	})

	// set content, order matters...
	win.SetContent(container.NewVBox(
		labelTitle,
		labelIP,
		labelValue,
		btn,
	),
	)
	win.ShowAndRun() // run app
}

func myIP() string {
	req, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return err.Error()
	}

	return ip.Query
}

type IP struct {
	Query string
}
