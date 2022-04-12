package client

import (
	"braverats/brp"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *App) SaveName(name string) {
	_, err := app.conn.Write(brp.NewReqSetName(name))
	app.gui.processSendErr(brp.ReqSetName, err)
	app.receiveAndProcessResponse(brp.ReqSetName,
		fmt.Sprintf("Name %q sucessfuly set", name),
		fmt.Sprintf("Failed to set name %q", name))
}

func (app *App) initGameMainMenu() {
	app.initLobby()

	dialogQuitConfirm := dialog.NewConfirm(
		"Quit", "Are you sure you want to quit the game?",
		func(ok bool) {
			if ok {
				app.gui.a.Quit()
			}
		},
		app.gui.w,
	)
	app.gui.dialogs[GIDDialQuitConfirm] = dialogQuitConfirm

	createLobbyBtn := widget.NewButton("Create Lobby", func() {
		app.gui.showDialog(GIDDialCreateLobby)
	})
	joinLobbyBtn := widget.NewButton("Join Lobby", func() {
		app.gui.showDialog(GIDDialJoinLobby)
	})
	gameQuitBtn := widget.NewButton("Quit", func() {
		app.gui.showDialog(GIDDialQuitConfirm)
	})
	btnMinSize := fyne.NewSize(200, 50)

	nicknameEntry := widget.NewEntry()
	saveNicknameBtn := widget.NewButton("Save", func() {
		app.SaveName(nicknameEntry.Text)
	})
	nicknameEntry.Validator = func(s string) error {
		if len(s) == 0 || len(s) > brp.MaxPlayerNameLen {
			if saveNicknameBtn != nil {
				saveNicknameBtn.Disable()
			}
			return errors.New("")
		}
		saveNicknameBtn.Enable()
		return nil
	}
	nicknameEntry.SetPlaceHolder("Enter your nickname")
	nicknameEntry.Resize(fyne.NewSize(200, nicknameEntry.MinSize().Height))
	nicknameEntry.Move(fyne.NewPos(gameWindowWidth/2-btnMinSize.Width/2, gameWindowHeight/3-btnMinSize.Height/2))
	saveNicknameBtn.Resize(fyne.NewSize(120, 35))
	saveNicknameBtn.Move(fyne.NewPos(gameWindowWidth/2-120/2, nicknameEntry.Position().Y+nicknameEntry.Size().Height+10))

	createLobbyBtn.Resize(btnMinSize)
	createLobbyBtn.Move(fyne.NewPos(gameWindowWidth/2-btnMinSize.Width/2, gameWindowHeight/2-btnMinSize.Height/2))
	joinLobbyBtn.Resize(btnMinSize)
	joinLobbyBtn.Move(fyne.NewPos(gameWindowWidth/2-btnMinSize.Width/2, gameWindowHeight/2-btnMinSize.Height/2+70))
	gameQuitBtn.Resize(btnMinSize)
	gameQuitBtn.Move(fyne.NewPos(gameWindowWidth/2-btnMinSize.Width/2, gameWindowHeight/1.2-btnMinSize.Height/2))

	mainGameMenuVBox := container.NewWithoutLayout()
	mainGameMenuVBox.Add(nicknameEntry)
	mainGameMenuVBox.Add(saveNicknameBtn)
	mainGameMenuVBox.Add(createLobbyBtn)
	mainGameMenuVBox.Add(joinLobbyBtn)
	mainGameMenuVBox.Add(gameQuitBtn)

	app.gui.w.SetContent(mainGameMenuVBox)
}
