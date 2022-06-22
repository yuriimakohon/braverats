# BraveRats

GO implementation of cardboard game [BraveRats](https://boardgamegeek.com/boardgame/112373/braverats)

<p align="center">
  <img src="https://raw.githubusercontent.com/yuriimakohon/braverats/master/screenshot_1.png" alt="photo5443094174451740500" border="0">
</p>

## Server

Server accepts connections via tcp/ip and speaks with clients on BRP protocol.

#### Run:

$ `go run cmd/server/main.go`

Use `-port` flag for custom port (default is `3000`)

## Client

Client written on [fyne](https://github.com/fyne-io/fyne) GUI library

#### Run:

$ `go run cmd/client/main.go`

Use `-addr` flag to specify server address (default is `localhost:3000`)

## BRP protocol

This repository provides implementation of application level protocol **BRP** based on tcp/ip.
Name of this protocol stands for "Brave Rats Protocol". It was developed special for this game implementation.

### Protocol specification

BRP works with 3 types of protocol`s messages called 'Tag':

- Request - client send this message to server. Every request should expect appropriate response-tag.
- Response - server response for client`s request.
- Event - server send this messages to client when some events occured.

#### Message format:

BRP message is simply stream of bytes separated by ending as delimiter `\r\n`  
Every message starts with tag and followed by its payload data separated by spaces:  
`TAG` __ `data_1` __ `...` __ `data_n` `\r\n`

Some examples:  
`REQ_CREATE_LOBBY Super kingdom\r\n`  
`REQ_PUT_CARD 4\r\n`  
`RESP_INFO lobby with such name already exists\r\n`  
`EVENT_CARD_PUT true 2\r\n`

#### Responses:

| **TAG**    |  **Payload**   | **Description**                                                                                                                                             |
|------------|:--------------:|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| RESP_OK    |   `message`    | Request successful.<br/>`message` - additional info, simple string                                                                                          |
| RESP_ERR   |    `error`     | Request unsuccessful.<br/>`error` - error message, simple string                                                                                            |
| RESP_INFO  |     `info`     | Request successful, but `info` should be handled in special way on client.<br/>`info` is string                                                             |
| RESP_LOBBY | `ready` `name` | Provide info about lobby, that client just joined.<br/>`ready` is **Bool**. Is owner of the lobby ready to play match.<br/>`name` - nickname of lobby owner |

#### Requests:
Every request can be responded with **RESP_ERR**

| **TAG**           | **Payload** | **Response**             | **Description**                                              |
|-------------------|:-----------:|--------------------------|--------------------------------------------------------------|
| REQ_SET_NAME      |   `name`    | RESP_OK                  | Set own nickname as `name`                                   |
| REQ_CREATE_LOBBY  |   `name`    | RESP_OK<br/>RESP_INFO    | Create new lobby with specified `name`                       |
| REQ_JOIN_LOBBY    |   `name`    | RESP_LOBBY<br/>RESP_INFO | Join to lobby with `name`                                    |
| REQ_LEAVE_LOBBY   |             | RESP_OK                  | Leave current lobby.                                         |
| REQ_SET_READINESS |   `ready`   | RESP_OK                  | Set own readiness to play match.<br/>`ready` is **Bool**     |
| REQ_START_MATCH   |             | RESP_OK                  | Start match.                                                 |
| REQ_PUT_CARD      |    `id`     | RESP_OK                  | Put card from your hand to the table.<br/>`id` is **CardID** |

#### Events:

| **TAG**                |  **Payload**   | **Condition**                                                                                                                      |
|------------------------|:--------------:|------------------------------------------------------------------------------------------------------------------------------------|
| EVENT_JOINED_LOBBY     |     `name`     | Another player joined to your lobby.<br/>`name` is his nickname                                                                    |
| EVENT_LEFT_LOBBY       |     `name`     | Another player left your lobby.<br/>`name` is his nickname                                                                         |
| EVENT_LOBBY_CLOSED     |                | Owner of the lobby closed it.                                                                                                      |
| EVENT_PLAYER_READINESS |    `ready`     | Another player changed his readiness to play match.<br/>`ready` is **Bool**                                                        |
| EVENT_MATCH_STARTED    |                | Owner of the lobby started match.                                                                                                  |
| EVENT_CARD_PUT         | `face up` `id` | Another player put his card from hand to the table.<br/>`face up` is **Bool**<br/>`id` is **CardID** (only if `face up` is `true`) |
| EVENT_ROUND_ENDED      | `result` `id`  | All players put their card to the table.<br/>`result` is **RoundResult**<br/>`id` is **CardID**. This is another`s player card     |

#### Types:

| **Type**    | **Values**                                                                                                                                                   |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Bool        | `true` `false`                                                                                                                                               |
| CardID      | `0` - Unknown<br/>`1` - Musician<br/>`2` - Princess<br/>`3` - Spy<br/>`4` - Assassin<br/>`5` - Ambassador<br/>`6` - Wizard<br/>`7` - General<br/>`8` - Prince |
| RoundResult | `0` - Won round<br/>`1` - Loosed round<br/>`2` - Held round<br/>`4` - Won game<br/>`5` - Loosed game<br/>`6` - Draw game                                     |
