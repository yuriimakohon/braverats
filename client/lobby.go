package client

import (
	"braverats/brp"
	"braverats/client/gui"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *App) CreateLobby(name string) {
	_, err := app.conn.Write(brp.NewReqCreateLobby(name))
	app.gui.ProcessSendErr(brp.ReqCreateLobby, err)
	if app.receiveAndProcessResponse(brp.ReqCreateLobby, "Lobby") {
		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) JoinLobby(name string) {
	_, err := app.conn.Write(brp.NewReqJoinLobby(name))
	app.gui.ProcessSendErr(brp.ReqJoinLobby, err)
	if app.receiveAndProcessResponse(brp.ReqJoinLobby, "Lobby") {
		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) LeaveLobby() {
	_, err := app.conn.Write(brp.NewReqLeaveLobby())
	app.gui.ProcessSendErr(brp.ReqLeaveLobby, err)
	app.receiveAndProcessResponse(brp.ReqLeaveLobby, "Lobby")
}

func (app *App) SetReadiness(ready bool) {
	_, err := app.conn.Write(brp.NewReqSetReadiness(ready))
	app.gui.ProcessSendErr(brp.ReqSetReadiness, err)
	app.receiveAndProcessResponse(brp.ReqSetReadiness, "Lobby")
}

func (app *App) JoinedLobby(name string) {
	app.gui.SendNotification("Lobby", fmt.Sprintf("%s joined the lobby", name))
}

func (app *App) LeftLobby(name string) {
	app.gui.SendNotification("Lobby", fmt.Sprintf("%s left the lobby", name))
}

func (app App) LobbyClosed() {
	app.gui.SendNotification("Lobby", "Owner left lobby")
	app.gui.ApplicationInfoDialog("Lobby closed", "The lobby owner has left the lobby")
}

var anotherPlayerReady = binding.NewBool()

func (app *App) PlayerReadiness(ready string) {
	r, err := strconv.ParseBool(ready)
	if err != nil {
		app.gui.ApplicationErrDialog("failed to parse readiness: " + err.Error())
		return
	}

	err = anotherPlayerReady.Set(r)
	if err != nil {
		app.gui.ApplicationErrDialog("failed to set another player readiness: " + err.Error())
		return
	}
	if r {
		app.gui.SendNotification("Lobby", "Another player is ready")
	} else {
		app.gui.SendNotification("Lobby", "Another player is not ready")
	}
}

func (app *App) MatchStarted() {
	app.gui.SendNotification("Match", "implement MatchStarted")
}

type lobby struct {
	app  *App // parent App
	name string
}

func newLobby(parentApp *App) *lobby {
	return &lobby{app: parentApp}
}

func (app *App) initLobby() {
	dialogCreateLobby := gui.NewLobbyCreatorDialog("Create lobby", "Create",
		func(name string) { app.CreateLobby(name) }, app.gui.W)
	app.gui.AddDialog(gui.GIDDialCreateLobby, dialogCreateLobby)

	dialogJoinLobby := gui.NewLobbyCreatorDialog("Join lobby", "Join",
		func(name string) { app.JoinLobby(name) }, app.gui.W)
	app.gui.AddDialog(gui.GIDDialJoinLobby, dialogJoinLobby)

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
	lobbyDialog := dialog.NewCustom("Lobby", "Leave", vbox, app.gui.W)
	lobbyDialog.SetOnClosed(func() {
		app.LeaveLobby()
	})
	app.gui.AddDialog(gui.GIDDialLobby, lobbyDialog)
}
