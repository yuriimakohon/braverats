package gui

import (
	"braverats/brp"
	"fmt"
	"log"

	"fyne.io/fyne/v2/dialog"
)

// ApplicationErrDialog logs and show application side error dialog if it's not nil.
func (gui *GUI) ApplicationErrDialog(err error) {
	if err != nil {
		log.Println("Application error: ", err)
		gui.ApplicationInfoDialog("Application error", err.Error())
	}
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

// ServerErrDialog logs and show message in server side error dialog .
func (gui *GUI) ServerErrDialog(msg string) {
	log.Println("Server error: ", msg)
	dialog.NewInformation("Server error", msg, gui.W).Show()
}

// SendErrDialog wraps error with send err message and call ApplicationErrDialog.
func (gui *GUI) SendErrDialog(tag brp.TAG, err error) {
	if err != nil {
		gui.ApplicationErrDialog(fmt.Errorf("error sending %s TAG to server: %v\n", tag, err))
	}
}
