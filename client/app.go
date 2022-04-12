package client

import (
	"braverats/brp"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	conn      net.Conn        // connection to game server
	responses chan brp.Packet // channel for responses from server
	events    chan brp.Packet // channel for events from server
	a         fyne.App        // gui engine from fyne.io framework
	w         fyne.Window     // main window of gui
}

// NewApp creates a new Brave Rats game client application. addr parameter is game server address.
func NewApp(addr string) *App {
	a := app.New()
	w := a.NewWindow(gameClientTitle)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Server didn't start: %v", err)
	}
	log.Println("Connected to server", conn.RemoteAddr())

	return &App{
		conn:      conn,
		responses: make(chan brp.Packet),
		events:    make(chan brp.Packet),
		a:         a,
		w:         w,
	}
}

func (app *App) Start() {
	app.init()
	go app.handleIncomingPackets()
	app.w.ShowAndRun()
}

func (app *App) init() {
	app.w.Resize(fyne.NewSize(gameWindowWidth, gameWindowHeight))
	app.w.SetMaster()

	app.initGameMainMenu()
}

func (app *App) handleIncomingPackets() {
	go app.handleEvents()
	for {
		packet, err := brp.ReadPacket(app.conn)
		if err != nil {
			log.Println(err)
			continue
		}
		switch packet.Tag {
		case brp.RespErr, brp.RespOk:
			app.responses <- packet
		default:
			app.events <- packet
		}
	}
}

func (app *App) handleEvents() {
	var packet brp.Packet
	for {
		packet = <-app.events
		switch packet.Tag {
		case brp.EventJoinedLobby:
			app.JoinedLobby(string(packet.Payload))
		case brp.EventLeftLobby:
			app.LeftLobby(string(packet.Payload))
		case brp.EventPlayerReadiness:
			app.PlayerReadiness(string(packet.Payload))
		case brp.EventMatchStarted:
			app.MatchStarted()
		}
	}
}

func (app *App) receiveResponse() (brp.Packet, error) {
	select {
	case response := <-app.responses:
		return response, nil
	case <-time.Tick(time.Second * 5):
		return brp.Packet{}, errors.New("response timeout")
	}
}

func (app *App) receiveAndProcessResponse(tag brp.TAG, okMsg, errMsg string) {
	resp, err := app.receiveResponse()
	if err != nil {
		log.Printf("Error receiving %s request`s response: %v", tag, err)
		return
	}

	switch resp.Tag {
	case brp.RespErr:
		errMsg = fmt.Sprintf("%s :: %s : %s", tag, resp, errMsg)
		app.serverErrDialog(errMsg)
	case brp.RespOk:
		okMsg = fmt.Sprintf("%s :: %s : %s", tag, resp, okMsg)
		log.Println(okMsg)
	}
}
