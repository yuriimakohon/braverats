package client

import (
	"braverats/brp"
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *App) CreateLobby(name string) {
	_, err := app.conn.Write(brp.NewReqCreateLobby(name))
	app.processSendErr(brp.ReqCreateLobby, err)
	app.receiveAndProcessResponse(brp.ReqCreateLobby,
		fmt.Sprintf("Lobby %q successfully created", name),
		fmt.Sprintf("Failed to create lobby %q", name))
}

func (app *App) JoinLobby(name string) {
	_, err := app.conn.Write(brp.NewReqJoinLobby(name))
	app.processSendErr(brp.ReqJoinLobby, err)
	app.receiveAndProcessResponse(brp.ReqJoinLobby,
		fmt.Sprintf("Successfully joined to %q lobby", name),
		fmt.Sprintf("Failed to join %q lobby", name))
}

func (app *App) LeaveLobby() {
	_, err := app.conn.Write(brp.NewReqLeaveLobby())
	app.processSendErr(brp.ReqLeaveLobby, err)
	app.receiveAndProcessResponse(brp.ReqLeaveLobby,
		"Successfully left lobby",
		"Failed to leave lobby")
}

func (app *App) JoinedLobby(name string) {
	app.sendNotification("Lobby", fmt.Sprintf("%s joined the lobby", name))
}

func (app *App) LeftLobby(name string) {
	app.sendNotification("Lobby", fmt.Sprintf("%s left the lobby", name))
}

func (app App) LobbyClosed() {
	app.sendNotification("Lobby", "Owner left lobby")
}

func (app *App) PlayerReadiness(ready string) {
	r, err := strconv.ParseBool(ready)
	if err != nil {
		log.Println("Failed to parse readiness: ", err)
		return
	}

	if r {
		app.sendNotification("Lobby", "Another player is ready")
	} else {
		app.sendNotification("Lobby", "Another player is not ready")
	}
}

func (app *App) MatchStarted() {
	app.sendNotification("Match", "implement MatchStarted")
}

var (
	createLobbyDialog dialog.Dialog
	joinLobbyDialog   dialog.Dialog
)

func (app *App) initLobby() {
	lobbyCreateNameString := binding.NewString()
	lobbyNameFormItem := widget.NewFormItem("Lobby name", widget.NewEntryWithData(lobbyCreateNameString))
	createLobbyDialog = dialog.NewForm(
		"Create lobby",
		"Create",
		"Cancel",
		[]*widget.FormItem{lobbyNameFormItem},
		func(confirm bool) {
			name, err := lobbyCreateNameString.Get()
			if err != nil {
				log.Println(err)
				return
			}
			if confirm && len(name) > 0 {
				app.CreateLobby(name)
			}
		},
		app.w,
	)

	lobbyJoinNameString := binding.NewString()
	lobbyJoinFormItem := widget.NewFormItem("Lobby name", widget.NewEntryWithData(lobbyJoinNameString))
	joinLobbyDialog = dialog.NewForm(
		"Join lobby",
		"Join",
		"Cancel",
		[]*widget.FormItem{lobbyJoinFormItem},
		func(confirm bool) {
			name, err := lobbyJoinNameString.Get()
			if err != nil {
				log.Println(err)
				return
			}
			if confirm && len(name) > 0 {
				app.JoinLobby(name)
			}
		},
		app.w,
	)
}
