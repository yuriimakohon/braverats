package client

import "fyne.io/fyne/v2"

const (
	gameClientTitle  = "Brave Rats"
	gameWindowWidth  = 1024
	gameWindowHeight = 768
)

func (app *App) sendNotification(title, message string) {
	app.a.SendNotification(fyne.NewNotification(title, message))
}
