package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start!")

	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(
		container.NewScroll(
			container.NewVBox(
				widget.NewButton("One", func() { fmt.Println("One") }),
				widget.NewButton("Two", func() { fmt.Println("Two") }),
				widget.NewButton("Three", func() { fmt.Println("Three") }),
				widget.NewButton("Four", func() { fmt.Println("Four") }),
				widget.NewButton("Five", func() { fmt.Println("Five") }),
				widget.NewButton("Six", func() { fmt.Println("Six") }),
				widget.NewButton("Seven", func() { fmt.Println("Seven") }),
				widget.NewButton("Eight", func() { fmt.Println("Eight") }),
				widget.NewButton("Nine", func() { fmt.Println("Nine") }),
				widget.NewButton("Ten", func() { fmt.Println("Ten") }),
			),
		),
	)

	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
