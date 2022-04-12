package brp

// TAG is type of request, response and events in BRP protocol.
type TAG string

// Client side request BRP TAGs.
const (
	ReqSetName      TAG = "SET_NAME"
	ReqCreateLobby  TAG = "CREATE_LOBBY"
	ReqJoinLobby    TAG = "JOIN_LOBBY"
	ReqLeaveLobby   TAG = "LEAVE_LOBBY"
	ReqSetReadiness TAG = "SET_READINESS"
	ReqStartMatch   TAG = "START_MATCH"
)

// Server side response BRP TAGs.
const (
	RespOk   TAG = "OK"
	RespErr  TAG = "ERR"
	RespInfo TAG = "INFO"
)

// Server side event BRP TAGs.
const (
	EventJoinedLobby     TAG = "JOINED_LOBBY"
	EventLeftLobby       TAG = "LEFT_LOBBY"
	EventLobbyClosed     TAG = "LOBBY_CLOSED"
	EventPlayerReadiness TAG = "PLAYER_READINESS"
	EventMatchStarted    TAG = "MATCH_STARTED"
)

// tags is a set of all BRP tags
var tags = map[TAG]struct{}{
	ReqSetName:      {},
	ReqCreateLobby:  {},
	ReqJoinLobby:    {},
	ReqLeaveLobby:   {},
	ReqSetReadiness: {},
	ReqStartMatch:   {},

	RespOk:   {},
	RespErr:  {},
	RespInfo: {},

	EventJoinedLobby:     {},
	EventLeftLobby:       {},
	EventLobbyClosed:     {},
	EventPlayerReadiness: {},
	EventMatchStarted:    {},
}
