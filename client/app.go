package client

import (
	"braverats/brp"
	"braverats/client/gui"
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
	"unicode"
	"unicode/utf8"
)

type App struct {
	conn      net.Conn        // connection to game server
	responses chan brp.Packet // channel for responses from server
	events    chan brp.Packet // channel for events from server
	gui       gui.GUI
	lobby     *lobby
}

// NewApp creates a new Brave Rats game client application. addr parameter is game server address.
func NewApp(addr string) *App {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Server didn't start: %v", err)
	}
	log.Println("Connected to server ", conn.RemoteAddr())

	app := &App{
		conn:      conn,
		responses: make(chan brp.Packet),
		events:    make(chan brp.Packet),
		gui:       gui.NewGUI(),
	}
	app.lobby = newLobby(app)

	return app
}

func (app *App) Start() {
	app.init()
	go app.handleIncomingPackets()
	app.gui.W.ShowAndRun()
}

func (app *App) init() {
	app.initGameMainMenu()
}

func (app *App) handleIncomingPackets() {
	go app.handleEvents()
	reader := bufio.NewReader(app.conn)
	scanner := bufio.NewScanner(reader)
	scanner.Split(brp.ScanCRLF)
	for scanner.Scan() {
		packet, err := brp.ParsePacket(scanner.Bytes())
		if err != nil {
			log.Println("Error parsing packet: ", err)
			continue
		}
		switch packet.Type {
		case brp.TypeResp:
			app.responses <- packet
		case brp.TypeEvent:
			app.events <- packet
		default:
			log.Printf("Client can't handle packet %s with %s type", packet, packet.Type)
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("lost connection with server")
	}
	app.gui.A.Quit()
}

func (app *App) handleEvents() {
	var packet brp.Packet
	for {
		packet = <-app.events
		log.Printf("%s : %s\n", packet.Tag, toUpperFirstLetter(string(packet.Payload)))
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

// receiveAndProcessResponse waits for a response from the server, then logs it.
//
// If the response is unsuccessful (RESP_ERR or RESP_INFO), displays appropriate dial.
//
// Returns success and the package with response
func (app *App) receiveAndProcessResponse(tag brp.TAG, title string) (bool, brp.Packet) {
	resp, err := app.receiveResponse()
	if err != nil {
		log.Printf("Error receiving %s request`s response: %v", tag, err)
		return false, resp
	}

	msg := toUpperFirstLetter(string(resp.Payload))

	switch resp.Tag {
	case brp.RespErr:
		app.gui.ServerErrDialog(fmt.Sprintf("%s :: %s : %s", tag, resp.Tag, msg))
		return false, resp
	case brp.RespInfo:
		log.Printf("%s :: %s : %s\n", tag, resp.Tag, msg)
		app.gui.ApplicationInfoDialog(title, msg)
		return false, resp
	default:
		log.Printf("%s :: %s : %s\n", tag, resp.Tag, msg)
		return true, resp
	}
}

// toUpperFirstLetter returns a string with the first letter capitalized.
func toUpperFirstLetter(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
