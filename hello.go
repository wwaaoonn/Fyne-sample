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
		container.NewVBox(
			container.NewAppTabs(
				container.NewTabItem("First", widget.NewLabel("This is the First tab item.")),
				container.NewTabItem("Second", widget.NewLabel("This is the Second tab item.")),
				container.NewTabItem("Third", widget.NewLabel("This is the Third tab item.")),
			),
		),
	)

	w.ShowAndRun()
}
