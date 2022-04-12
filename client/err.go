package client

import (
	"braverats/brp"
	"fmt"
	"log"

	"fyne.io/fyne/v2/dialog"
)

func (app *App) applicationErrDialog(msg string) {
	dialog.NewInformation("Application error", msg, app.w).Show()
}

func (app *App) serverErrDialog(msg string) {
	log.Println(msg)
	dialog.NewInformation("Server error", msg, app.w).Show()
}

func (app *App) processSendErr(tag brp.TAG, err error) {
	if err != nil {
		msg := fmt.Sprintf("Error sending %s TAG to server: %v\n", tag, err)
		log.Printf(msg)
		app.applicationErrDialog(msg)
	}
}
