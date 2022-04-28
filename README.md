# BraveRats

GO implementation of cardboard game [BraveRats](https://boardgamegeek.com/boardgame/112373/braverats)

<p align="center">
  <img src="https://raw.githubusercontent.com/yuriimakohon/braverats/master/screenshot_1.png" alt="photo5443094174451740500" border="0">
</p>

## BRP protocol
This repository provides implementation of application level protocol **BRP** based on tcp/ip

## Server
Server accepts connections via tcp/ip and speaks with clients on BRP protocol.
#### Run:
`go run cmd/server/main.go`

Use `-port` flag for custom port (default is `3000`)

## Client
Client written on [fyne](https://github.com/fyne-io/fyne) GUI library
#### Run:
`go run cmd/client/main.go`

Use `-addr` flag to specify server address (default is `localhost:3000`)
