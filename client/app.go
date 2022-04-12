package client

import (
	"braverats/brp"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

type App struct {
	conn      net.Conn        // connection to game server
	responses chan brp.Packet // channel for responses from server
	events    chan brp.Packet // channel for events from server
	gui       GUI
}

// NewApp creates a new Brave Rats game client application. addr parameter is game server address.
func NewApp(addr string) *App {
	a := app.New()
	w := a.NewWindow(gameClientTitle)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Server didn't start: %v", err)
	}
	log.Println("Connected to server ", conn.RemoteAddr())

	return &App{
		conn:      conn,
		responses: make(chan brp.Packet),
		events:    make(chan brp.Packet),
		gui: GUI{
			a:       a,
			w:       w,
			dialogs: make(map[GID]dialog.Dialog),
		},
	}
}

func (app *App) Start() {
	app.init()
	go app.handleIncomingPackets()
	app.gui.w.ShowAndRun()
}

func (app *App) init() {
	app.gui.w.Resize(fyne.NewSize(gameWindowWidth, gameWindowHeight))
	app.gui.w.SetMaster()

	app.initGameMainMenu()
}

func (app *App) handleIncomingPackets() {
	go app.handleEvents()
	for {
		packet, err := brp.ReadPacket(app.conn)
		if err == io.EOF {
			log.Println("lost connection with server")
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		switch packet.Tag {
		case brp.RespOk, brp.RespErr, brp.RespInfo:
			app.responses <- packet
		default:
			app.events <- packet
		}
	}
	app.gui.a.Quit()
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
		case brp.EventLobbyClosed:
			app.LobbyClosed()
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

func (app *App) receiveAndProcessResponse(tag brp.TAG, title string) {
	resp, err := app.receiveResponse()
	if err != nil {
		log.Printf("Error receiving %s request`s response: %v", tag, err)
		return
	}

	// Capitalize first letter of response message
	msg := string(append(bytes.ToUpper(resp.Payload[0:1]), resp.Payload[1:]...))

	switch resp.Tag {
	case brp.RespOk:
		log.Printf("%s :: %s : %s\n", tag, resp.Tag, msg)
	case brp.RespErr:
		app.gui.serverErrDialog(fmt.Sprintf("%s :: %s : %s", tag, resp.Tag, msg))
	case brp.RespInfo:
		log.Printf("%s :: %s : %s\n", tag, resp.Tag, msg)
		app.gui.applicationInfoDialog(title, msg)
	}
}
