package client

import (
	"braverats/brp"
	"fmt"
	"log"

	"fyne.io/fyne/v2/dialog"
)

func (gui *GUI) applicationErrDialog(msg string) {
	log.Println("Application error: ", msg)
	gui.applicationInfoDialog("Application error", msg)
}

func (gui *GUI) applicationInfoDialog(title, msg string) {
	dialog.ShowInformation(title, msg, gui.w)
}

// applicationErrDialog displays an information dialog and show another dialog by GID.
func (gui *GUI) applicationInfoPopup(title, msg string, gid GID) {
	dial := dialog.NewInformation(title, msg, gui.w)
	dial.SetOnClosed(func() {
		gui.showDialog(gid)
	})
	dial.Show()
}

func (gui *GUI) serverErrDialog(msg string) {
	log.Println("Server error: ", msg)
	dialog.NewInformation("Server error", msg, gui.w).Show()
}

func (gui *GUI) processSendErr(tag brp.TAG, err error) {
	if err != nil {
		msg := fmt.Sprintf("error sending %s TAG to server: %v\n", tag, err)
		gui.applicationErrDialog(msg)
	}
}
