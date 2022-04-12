package client

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

const (
	gameClientTitle  = "Brave Rats"
	gameWindowWidth  = 1024
	gameWindowHeight = 768
)

// GID is GUI widget`s unique ID for certain dialogs, containers etc.
type GID uint8

const (
	GIDDialQuitConfirm GID = iota
	GIDDialCreateLobby
	GIDDialJoinLobby
	GIDDialLobby
)

type GUI struct {
	a       fyne.App
	w       fyne.Window
	dialogs map[GID]dialog.Dialog
}

func (gui *GUI) showDialog(gid GID) {
	dial, ok := gui.dialogs[gid]
	if !ok {
		log.Println("dialog not found: ", gid)
		return
	}
	if dial == nil {
		log.Println("dialog is nil: ", gid)
		return
	}

	dial.Show()
}

func (gui *GUI) sendNotification(title, message string) {
	gui.a.SendNotification(fyne.NewNotification(title, message))
}
