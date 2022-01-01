package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start!")

	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon!")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			l.SetText("Select Infomation Icon!")
		}),
		widget.NewToolbarAction(theme.ComputerIcon(), func() {
			l.SetText("Select Computer Icon!")
		}),
	)

	w.SetContent(
		container.New(
			layout.NewBorderLayout(
				nil, tb, nil, nil,
			),
			l,
			tb,
		),
	)

	w.Resize(fyne.NewSize(800, 500))
	w.ShowAndRun()
}
