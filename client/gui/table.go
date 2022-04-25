package gui

import (
	"fyne.io/fyne/v2"
)

type tableLayout struct {
}

func (tr tableLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	hPadding := float32(25)
	vPadding := float32(75)
	w := size.Width - hPadding*2
	h := size.Height/2 - vPadding
	objects[0].Resize(fyne.NewSize(w, h))
	objects[1].Resize(fyne.NewSize(w, h))
	objects[0].Move(fyne.NewPos(hPadding, vPadding))
	objects[1].Move(fyne.NewPos(hPadding, vPadding+h))
}

func (tr tableLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[0].MinSize()
}
