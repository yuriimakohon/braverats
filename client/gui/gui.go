package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

const (
	gameClientTitle  = "Brave Rats"
	GameWindowWidth  = 1024
	GameWindowHeight = 768
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
	A       fyne.App
	W       fyne.Window
	dialogs map[GID]dialog.Dialog
}

func NewGUI() GUI {
	a := app.New()
	w := a.NewWindow(gameClientTitle)
	w.Resize(fyne.NewSize(GameWindowWidth, GameWindowHeight))
	w.CenterOnScreen()
	w.SetMaster()

	return GUI{
		A:       a,
		W:       w,
		dialogs: make(map[GID]dialog.Dialog),
	}
}

func (gui *GUI) AddDialog(gid GID, d dialog.Dialog) {
	gui.dialogs[gid] = d
}

func (gui *GUI) ShowDialog(gid GID) {
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

func (gui *GUI) SendNotification(title, message string) {
	gui.A.SendNotification(fyne.NewNotification(title, message))
}
