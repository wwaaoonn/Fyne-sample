package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start!")

	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne")
	r := widget.NewRadioGroup(
		[]string{"One", "Two", "Three"},
		func(s string) {
			if s == "" {
				l.SetText("not selected..")
			} else {
				l.SetText("selected: " + s)
			}
		},
	)
	r.SetSelected("One")
	w.SetContent(
		container.NewVBox(
			l,
			r,
		),
	)

	w.ShowAndRun()
}
