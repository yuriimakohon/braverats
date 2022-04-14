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
	if ok, _ := app.receiveAndProcessResponse(brp.ReqCreateLobby, "Lobby"); ok {
		app.lobby.playerIn = true

		err = app.lobby.gui.Reset(name, true)
		app.gui.ApplicationErrDialog(err)

		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) JoinLobby(lobbyName string) {
	_, err := app.conn.Write(brp.NewReqJoinLobby(lobbyName))
	app.gui.SendErrDialog(brp.ReqJoinLobby, err)
	if ok, resp := app.receiveAndProcessResponse(brp.ReqJoinLobby, "Lobby"); ok {
		app.lobby.playerIn = true

		err = app.lobby.gui.Reset(lobbyName, false)
		app.gui.ApplicationErrDialog(err)

		ready, opponentName := app.RespLobby(resp.Payload)
		err = app.lobby.gui.SecondPlayer.Ready.Set(ready)
		app.gui.ApplicationErrDialog(err)
		err = app.lobby.gui.SecondPlayer.Name.Set(opponentName)
		app.gui.ApplicationErrDialog(err)

		app.gui.ShowDialog(gui.GIDDialLobby)
	}
}

func (app *App) LeaveLobby() {
	_, err := app.conn.Write(brp.NewReqLeaveLobby())
	app.gui.SendErrDialog(brp.ReqLeaveLobby, err)
	if ok, _ := app.receiveAndProcessResponse(brp.ReqLeaveLobby, "Lobby"); ok {
		app.lobby.playerIn = false
	}
}

func (app *App) SetReadiness(ready bool) {
	_, err := app.conn.Write(brp.NewReqSetReadiness(ready))
	app.gui.SendErrDialog(brp.ReqSetReadiness, err)
	if ok, _ := app.receiveAndProcessResponse(brp.ReqSetReadiness, "Lobby"); ok {
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
	app.lobby.playerIn = false
	app.gui.HideDialog(gui.GIDDialLobby)
	app.gui.ApplicationInfoDialog("Lobby closed", "The lobby owner has left the lobby")
}

func (app *App) PlayerReadiness(ready string) {
	r, err := strconv.ParseBool(ready)
	app.gui.ApplicationErrDialog(err)

	err = app.lobby.gui.SecondPlayer.Ready.Set(r)
	app.gui.ApplicationErrDialog(err)
}

func (app *App) StartMatch() {
	_, err := app.conn.Write(brp.NewReqStartMatch())
	app.gui.SendErrDialog(brp.ReqStartMatch, err)
	app.receiveAndProcessResponse(brp.ReqStartMatch, "Match")
}

func (app *App) MatchStarted() {
	app.gui.ApplicationInfoDialog("Match started", "implement MatchStarted")
}

type lobby struct {
	app      *App // parent App
	gui      *gui.Lobby
	playerIn bool
}

func newLobby(parentApp *App) *lobby {
	return &lobby{
		app: parentApp,
		gui: gui.NewLobby(),
	}
}

func (app *App) RespLobby(payload []byte) (bool, string) {
	args := bytes.Split(payload, []byte(" "))

	ready, err := strconv.ParseBool(string(args[0]))
	app.gui.ApplicationErrDialog(err)

	return ready, string(bytes.Join(args[1:], []byte(" ")))
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
		func() {
			if app.lobby.playerIn {
				app.LeaveLobby()
			}
		},
		func() { app.StartMatch() },
		*app.lobby.gui,
		app.gui.W,
	)

	app.gui.AddDialog(gui.GIDDialLobby, dialogLobby)
}
