package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start!")

	a := app.New()
	w := a.NewWindow("Hello")

	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Bottom", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)

	w.SetContent(
		container.New(
			layout.NewBorderLayout(
				bt, bb, bl, br,
			),
			bt, bb, bl, br,
			widget.NewLabel("Center"),
		),
	)

	w.ShowAndRun()
}
