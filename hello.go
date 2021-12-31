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
	c := widget.NewCheck("Check me!", func(b bool) {
		if b {
			l.SetText("CHECKED!!")
		} else {
			l.SetText("not checked..")
		}
	})
	c.SetChecked(true)
	w.SetContent(
		container.NewVBox(
			l,
			c,
		),
	)

	w.ShowAndRun()
}
