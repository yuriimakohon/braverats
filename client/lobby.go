package client

import (
	"braverats/brp"
	"braverats/client/gui"
	"bytes"
	"fmt"
	"strconv"
)

func (app *App) CreateLobby(name string) {
	_, err := app.conn.Write(brp.NewReqCreateLobby(name))
	app.gui.SendErrDialog(brp.ReqCreateLobby, err)
	if app.receiveAndProcessResponse(brp.ReqCreateLobby, "Lobby") {
		err = app.lobby.gui.Reset(name)
		app.gui.ApplicationErrDialog(err)
		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) JoinLobby(name string) {
	_, err := app.conn.Write(brp.NewReqJoinLobby(name))
	app.gui.SendErrDialog(brp.ReqJoinLobby, err)
	if app.receiveAndProcessResponse(brp.ReqJoinLobby, "Lobby") {
		err = app.lobby.gui.Name.Set(name)
		app.gui.ApplicationErrDialog(err)
		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) LeaveLobby() {
	_, err := app.conn.Write(brp.NewReqLeaveLobby())
	app.gui.SendErrDialog(brp.ReqLeaveLobby, err)
	app.receiveAndProcessResponse(brp.ReqLeaveLobby, "Lobby")
}

func (app *App) SetReadiness(ready bool) {
	_, err := app.conn.Write(brp.NewReqSetReadiness(ready))
	app.gui.SendErrDialog(brp.ReqSetReadiness, err)
	if app.receiveAndProcessResponse(brp.ReqSetReadiness, "Lobby") {
		err = app.lobby.gui.FirstPlayer.Ready.Set(ready)
		app.gui.ApplicationErrDialog(err)
	}
}

func (app *App) JoinedLobby(name string) {
	app.gui.SendNotification("Lobby", fmt.Sprintf("%s joined the lobby", name))
	err := app.lobby.gui.SecondPlayer.Name.Set(name)
	app.gui.ApplicationErrDialog(err)
}

func (app *App) LeftLobby(name string) {
	app.gui.SendNotification("Lobby", fmt.Sprintf("%s left the lobby", name))
	err := app.lobby.gui.ResetSecondPlayer()
	app.gui.ApplicationErrDialog(err)
}

func (app *App) LobbyClosed() {
	app.gui.SendNotification("Lobby", "Owner left lobby")
	app.gui.HideDialog(gui.GIDDialLobby)
	app.gui.ApplicationInfoDialog("Lobby closed", "The lobby owner has left the lobby")
}

func (app *App) PlayerReadiness(ready string) {
	r, err := strconv.ParseBool(ready)
	app.gui.ApplicationErrDialog(err)

	err = app.lobby.gui.SecondPlayer.Ready.Set(r)
	app.gui.ApplicationErrDialog(err)
}

func (app *App) MatchStarted() {
	app.gui.SendNotification("Match", "implement MatchStarted")
}

type lobby struct {
	app *App // parent App
	gui *gui.Lobby
}

func newLobby(parentApp *App) *lobby {
	return &lobby{
		app: parentApp,
		gui: gui.NewLobby(),
	}
}

func (l *lobby) RespLobby(payload []byte) {
	args := bytes.Split(payload, []byte(" "))
	ready, err := strconv.ParseBool(string(args[0]))
	l.app.gui.ApplicationErrDialog(err)
	name := string(bytes.Join(args[1:], []byte(" ")))

	err = l.gui.Reset("")
	l.app.gui.ApplicationErrDialog(err)
	err = l.gui.SecondPlayer.Ready.Set(ready)
	l.app.gui.ApplicationErrDialog(err)
	err = l.gui.SecondPlayer.Name.Set(name)
	l.app.gui.ApplicationErrDialog(err)
}

func (app *App) initLobby() {
	dialogCreateLobby := gui.NewLobbyCreatorDialog("Create lobby", "Create",
		func(name string) { app.CreateLobby(name) }, app.gui.W)
	app.gui.AddDialog(gui.GIDDialCreateLobby, dialogCreateLobby)

	dialogJoinLobby := gui.NewLobbyCreatorDialog("Join lobby", "Join",
		func(name string) { app.JoinLobby(name) }, app.gui.W)
	app.gui.AddDialog(gui.GIDDialJoinLobby, dialogJoinLobby)

	dialogLobby := gui.NewLobbyDialog(
		func(ready bool) { app.SetReadiness(ready) },
		func() { app.LeaveLobby() },
		*app.lobby.gui,
		app.gui.W,
	)

	app.gui.AddDialog(gui.GIDDialLobby, dialogLobby)
}
