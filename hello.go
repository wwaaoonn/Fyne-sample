package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("start!")

	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")
	b := widget.NewButton("Confirm", func() {
		dialog.ShowConfirm("Confirm",
			"Please check 'YES'!",
			func(b bool) {
				if b {
					l.SetText("OK, thank you!!")
				} else {
					l.SetText("oh...")
				}
			},
			w,
		)
	})

	w.SetContent(
		container.New(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l,
			b,
		),
	)

	w.Resize(fyne.NewSize(800, 500))
	w.ShowAndRun()
}
