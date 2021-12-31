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
	w.SetContent(
		container.NewHBox(
			container.NewVBox(
				widget.NewLabel("Hello Fyne!"),
				widget.NewLabel("This is the sample app!"),
			),
			container.NewVBox(
				widget.NewLabel("Hello Fyne!"),
				widget.NewLabel("This is the sample app!"),
			),
		),
	)

	w.ShowAndRun()
}
