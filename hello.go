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
	l := widget.NewLabel("Hello Fyne!")
	ne := widget.NewEntry()
	pe := widget.NewPasswordEntry()
	fr := widget.NewForm(
		widget.NewFormItem("Name", ne),
		widget.NewFormItem("Pass", pe),
	)
	b := widget.NewButton("Click", func() {
		dialog.ShowCustomConfirm("Enter message.", "OK", "Cancel", fr, func(b bool) {
			if b {
				l.SetText("Name: '" + ne.Text + "' Pass: '" + pe.Text + "'")
			} else {
				l.SetText("no message...")
			}
		}, w)
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
