package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	win := app.NewWindow("Testing Tree Window")

	win.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Testing Tree Title")
	intro := widget.NewLabel("testing `tree` Label...")
	intro.Wrapping = fyne.TextWrapWord

	myContainer := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(), myContainer)

	split.Offset = 0
	win.SetContent(split)

	win.Resize(fyne.NewSize(440, 260))
	win.ShowAndRun()
}

func makeNav() fyne.CanvasObject {

	tree := widget.NewTreeWithStrings(menuItems)

	return container.NewBorder(nil, nil, nil, nil, tree)
}

var menuItems = map[string][]string{
	"":             {"welcome", "collections", "collections2", "advanced"},
	"collections":  {"list", "table"},
	"collections2": {"list2", "table2"},
	"welcome":      {"welcome again"},
}
