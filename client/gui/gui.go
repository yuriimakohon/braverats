package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

const (
	AssetsDir        = "~Desktop/braverats/client/assets"
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
	GIDDialMatchEnd
	GIDMainMenu
	GIDMatch
)

type GUI struct {
	A       fyne.App
	W       fyne.Window
	dialogs map[GID]dialog.Dialog
	scenes  map[GID]*fyne.Container
}

func NewGUI() GUI {
	a := app.New()
	w := a.NewWindow(gameClientTitle)
	w.Resize(fyne.NewSize(GameWindowWidth, GameWindowHeight))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.SetMaster()

	return GUI{
		A:       a,
		W:       w,
		dialogs: make(map[GID]dialog.Dialog),
		scenes:  make(map[GID]*fyne.Container),
	}
}

func (gui *GUI) AddDialog(gid GID, dialog dialog.Dialog) {
	gui.dialogs[gid] = dialog
}

func (gui *GUI) ShowDialog(gid GID) {
	if dial, ok := gui.dialogs[gid]; ok {
		dial.Show()
		return
	}
	log.Println("dialog not found: ", gid)
	return
}

func (gui *GUI) HideDialog(gid GID) {
	if dial, ok := gui.checkDialog(gid); ok {
		dial.Hide()
	}
}

func (gui *GUI) checkDialog(gid GID) (dialog.Dialog, bool) {
	if dial, ok := gui.dialogs[gid]; ok {
		return dial, true
	}
	log.Println("dialog not found: ", gid)
	return nil, false
}

func (gui *GUI) AddScene(gid GID, scene *fyne.Container) {
	gui.scenes[gid] = scene
}

func (gui *GUI) ShowScene(gid GID) {
	scene, ok := gui.checkScene(gid)
	if !ok {
		log.Println("scene not found: ", gid)
		return
	}
	// gui.W.Content().Hide()
	gui.W.SetContent(scene)
	// gui.W.Content().Show()
}

func (gui *GUI) checkScene(gid GID) (*fyne.Container, bool) {
	if scene, ok := gui.scenes[gid]; ok {
		return scene, true
	}
	log.Println("scene not found: ", gid)
	return nil, false
}

func (gui *GUI) SendNotification(title, message string) {
	gui.A.SendNotification(fyne.NewNotification(title, message))
}
