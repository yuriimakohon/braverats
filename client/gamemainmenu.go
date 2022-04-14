package client

import (
	"braverats/brp"
	"braverats/client/gui"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *App) SetNickname(name string) {
	_, err := app.conn.Write(brp.NewReqSetName(name))
	app.gui.SendErrDialog(brp.ReqSetName, err)
	if ok, _ := app.receiveAndProcessResponse(brp.ReqSetName, "Nickname"); ok {
		app.gui.SendNotification("Nickname saved", name)
	}
}

func (app *App) initGameMainMenu() {
	app.initLobby()

	dialogQuitConfirm := dialog.NewConfirm(
		"Quit", "Are you sure you want to quit the game?",
		func(ok bool) {
			if ok {
				app.gui.A.Quit()
			}
		},
		app.gui.W,
	)
	app.gui.AddDialog(gui.GIDDialQuitConfirm, dialogQuitConfirm)

	createLobbyBtn := widget.NewButton("Create Lobby", func() {
		app.gui.ShowDialog(gui.GIDDialCreateLobby)
	})
	joinLobbyBtn := widget.NewButton("Join Lobby", func() {
		app.gui.ShowDialog(gui.GIDDialJoinLobby)
	})
	gameQuitBtn := widget.NewButton("Quit", func() {
		app.gui.ShowDialog(gui.GIDDialQuitConfirm)
	})
	btnMinSize := fyne.NewSize(200, 50)

	nicknameEntry := widget.NewEntry()
	saveNicknameBtn := widget.NewButton("Save", func() {
		app.SetNickname(nicknameEntry.Text)
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
	nicknameEntry.Move(fyne.NewPos(gui.GameWindowWidth/2-btnMinSize.Width/2, gui.GameWindowHeight/3-btnMinSize.Height/2))
	saveNicknameBtn.Resize(fyne.NewSize(120, 35))
	saveNicknameBtn.Move(fyne.NewPos(gui.GameWindowWidth/2-120/2, nicknameEntry.Position().Y+nicknameEntry.Size().Height+10))

	createLobbyBtn.Resize(btnMinSize)
	createLobbyBtn.Move(fyne.NewPos(gui.GameWindowWidth/2-btnMinSize.Width/2, gui.GameWindowHeight/2-btnMinSize.Height/2))
	joinLobbyBtn.Resize(btnMinSize)
	joinLobbyBtn.Move(fyne.NewPos(gui.GameWindowWidth/2-btnMinSize.Width/2, gui.GameWindowHeight/2-btnMinSize.Height/2+70))
	gameQuitBtn.Resize(btnMinSize)
	gameQuitBtn.Move(fyne.NewPos(gui.GameWindowWidth/2-btnMinSize.Width/2, gui.GameWindowHeight/1.2-btnMinSize.Height/2))

	mainGameMenuVBox := container.NewWithoutLayout()
	mainGameMenuVBox.Add(nicknameEntry)
	mainGameMenuVBox.Add(saveNicknameBtn)
	mainGameMenuVBox.Add(createLobbyBtn)
	mainGameMenuVBox.Add(joinLobbyBtn)
	mainGameMenuVBox.Add(gameQuitBtn)

	app.gui.W.SetContent(mainGameMenuVBox)
}
