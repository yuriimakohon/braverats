package gui

import (
	"braverats/brp"
	"fmt"
	"log"

	"fyne.io/fyne/v2/dialog"
)

func (gui *GUI) ApplicationErrDialog(msg string) {
	log.Println("Application error: ", msg)
	gui.ApplicationInfoDialog("Application error", msg)
}

func (gui *GUI) ApplicationInfoDialog(title string, msg string) {
	dialog.ShowInformation(title, msg, gui.W)
}

// ApplicationInfoPopup displays an information dialog and show another dialog by GID.
func (gui *GUI) ApplicationInfoPopup(title, msg string, gid GID) {
	dial := dialog.NewInformation(title, msg, gui.W)
	dial.SetOnClosed(func() {
		gui.ShowDialog(gid)
	})
	dial.Show()
}

func (gui *GUI) ServerErrDialog(msg string) {
	log.Println("Server error: ", msg)
	dialog.NewInformation("Server error", msg, gui.W).Show()
}

func (gui *GUI) ProcessSendErr(tag brp.TAG, err error) {
	if err != nil {
		msg := fmt.Sprintf("error sending %s TAG to server: %v\n", tag, err)
		gui.ApplicationErrDialog(msg)
	}
}
