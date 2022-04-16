package brp

// TAG is type of request, response and events in BRP protocol.
type TAG string

// Client side request BRP TAGs.
const (
	ReqSetName      TAG = "REQ_SET_NAME"
	ReqCreateLobby  TAG = "REQ_CREATE_LOBBY"
	ReqJoinLobby    TAG = "REQ_JOIN_LOBBY"
	ReqLeaveLobby   TAG = "REQ_LEAVE_LOBBY"
	ReqSetReadiness TAG = "REQ_SET_READINESS"
	ReqStartMatch   TAG = "REQ_START_MATCH"
)

// Server side response BRP TAGs.
const (
	RespOk    TAG = "RESP_OK"
	RespErr   TAG = "RESP_ERR"
	RespInfo  TAG = "RESP_INFO"
	RespLobby TAG = "RESP_LOBBY"
)

// Server side event BRP TAGs.
const (
	EventJoinedLobby     TAG = "EVENT_JOINED_LOBBY"
	EventLeftLobby       TAG = "EVENT_LEFT_LOBBY"
	EventLobbyClosed     TAG = "EVENT_LOBBY_CLOSED"
	EventPlayerReadiness TAG = "EVENT_PLAYER_READINESS"
	EventMatchStarted    TAG = "EVENT_MATCH_STARTED"
)

func IsTAG(tag TAG) bool {
	_, ok := tags[tag]
	return ok
}

// tags is a set of all BRP tags
var tags = map[TAG]struct{}{
	ReqSetName:      {},
	ReqCreateLobby:  {},
	ReqJoinLobby:    {},
	ReqLeaveLobby:   {},
	ReqSetReadiness: {},
	ReqStartMatch:   {},

	RespOk:    {},
	RespErr:   {},
	RespInfo:  {},
	RespLobby: {},

	EventJoinedLobby:     {},
	EventLeftLobby:       {},
	EventLobbyClosed:     {},
	EventPlayerReadiness: {},
	EventMatchStarted:    {},
}
