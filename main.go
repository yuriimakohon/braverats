package main

import (
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app2.New()
	w := app.NewWindow("Test")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(200, 200))

	first := container.NewCenter(widget.NewLabel("First"))
	second := container.NewCenter(widget.NewLabel("Second"))

	w.SetContent(first)
	w.SetContent(second)
	w.ShowAndRun()
}
