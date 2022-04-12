package client

import (
	"braverats/brp"
	"errors"
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *App) CreateLobby(name string) {
	_, err := app.conn.Write(brp.NewReqCreateLobby(name))
	app.gui.processSendErr(brp.ReqCreateLobby, err)
	if app.receiveAndProcessResponse(brp.ReqCreateLobby, "Lobby") {
		app.gui.showDialog(GIDDialLobby)
	}
}

func (app *App) JoinLobby(name string) {
	_, err := app.conn.Write(brp.NewReqJoinLobby(name))
	app.gui.processSendErr(brp.ReqJoinLobby, err)
	if app.receiveAndProcessResponse(brp.ReqJoinLobby, "Lobby") {
		app.gui.showDialog(GIDDialLobby)
	}
}

func (app *App) LeaveLobby() {
	_, err := app.conn.Write(brp.NewReqLeaveLobby())
	app.gui.processSendErr(brp.ReqLeaveLobby, err)
	app.receiveAndProcessResponse(brp.ReqLeaveLobby, "Lobby")
}

func (app *App) SetReadiness(ready bool) {
	_, err := app.conn.Write(brp.NewReqSetReadiness(ready))
	app.gui.processSendErr(brp.ReqSetReadiness, err)
	app.receiveAndProcessResponse(brp.ReqSetReadiness, "Lobby")
}

func (app *App) JoinedLobby(name string) {
	app.gui.sendNotification("Lobby", fmt.Sprintf("%s joined the lobby", name))
}

func (app *App) LeftLobby(name string) {
	app.gui.sendNotification("Lobby", fmt.Sprintf("%s left the lobby", name))
}

func (app App) LobbyClosed() {
	app.gui.sendNotification("Lobby", "Owner left lobby")
	app.gui.applicationInfoDialog("Lobby closed", "The lobby owner has left the lobby")
}

var anotherPlayerReady = binding.NewBool()

func (app *App) PlayerReadiness(ready string) {
	r, err := strconv.ParseBool(ready)
	if err != nil {
		app.gui.applicationErrDialog("failed to parse readiness: " + err.Error())
		return
	}

	err = anotherPlayerReady.Set(r)
	if err != nil {
		app.gui.applicationErrDialog("failed to set another player readiness: " + err.Error())
		return
	}
	if r {
		app.gui.sendNotification("Lobby", "Another player is ready")
	} else {
		app.gui.sendNotification("Lobby", "Another player is not ready")
	}
}

func (app *App) MatchStarted() {
	app.gui.sendNotification("Match", "implement MatchStarted")
}

func (app *App) initLobby() {
	lobbyNameValidator := func(s string) error {
		if len(s) == 0 || len(s) > brp.MaxLobbyNameLen {
			return errors.New("")
		}
		return nil
	}
	dialogSize := fyne.NewSize(300, 100)

	lobbyCreateNameString := binding.NewString()
	lobbyCreateNameEntry := widget.NewEntryWithData(lobbyCreateNameString)
	lobbyCreateNameEntry.Validator = lobbyNameValidator
	lobbyNameFormItem := widget.NewFormItem("Lobby name", lobbyCreateNameEntry)
	dialogCreateLobby := dialog.NewForm(
		"Create lobby",
		"Create",
		"Cancel",
		[]*widget.FormItem{lobbyNameFormItem},
		func(confirm bool) {
			if !confirm {
				return
			}
			name, err := lobbyCreateNameString.Get()
			if err != nil {
				log.Println(err)
				return
			}
			app.CreateLobby(name)
		},
		app.gui.w,
	)
	dialogCreateLobby.Resize(dialogSize)
	app.gui.dialogs[GIDDialCreateLobby] = dialogCreateLobby

	lobbyJoinNameString := binding.NewString()
	lobbyJoinNameEntry := widget.NewEntryWithData(lobbyJoinNameString)
	lobbyJoinNameEntry.Validator = lobbyNameValidator
	lobbyJoinFormItem := widget.NewFormItem("Lobby name", lobbyJoinNameEntry)
	dialogJoinLobby := dialog.NewForm(
		"Join lobby",
		"Join",
		"Cancel",
		[]*widget.FormItem{lobbyJoinFormItem},
		func(confirm bool) {
			if !confirm {
				return
			}
			name, err := lobbyJoinNameString.Get()
			if err != nil {
				log.Println(err)
				return
			}
			if len(name) <= 0 {
				app.gui.applicationInfoDialog("Lobby", "Lobby name cannot be empty")
				return
			}
			app.JoinLobby(name)
		},
		app.gui.w,
	)
	dialogJoinLobby.Resize(dialogSize)
	app.gui.dialogs[GIDDialJoinLobby] = dialogJoinLobby

	ownReadinessCheck := widget.NewCheck("Ready", func(ready bool) {
		app.SetReadiness(ready)
	})
	anotherPlayerReadinessCheck := &widget.Check{
		Text:    "Another player is ready",
		Checked: false,
		OnChanged: func(ready bool) {
			app.PlayerReadiness(strconv.FormatBool(ready))
		},
	}
	anotherPlayerReadinessCheck.Disable()
	anotherPlayerReadinessCheck.Bind(anotherPlayerReady)
	vbox := container.NewVBox(ownReadinessCheck, anotherPlayerReadinessCheck)
	lobbyDialog := dialog.NewCustom("Lobby", "Leave", vbox, app.gui.w)
	lobbyDialog.SetOnClosed(func() {
		app.LeaveLobby()
	})
	app.gui.dialogs[GIDDialLobby] = lobbyDialog
}
