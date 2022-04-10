package protocol

// Ending is used to separate packets in the stream.
var Ending = []byte{0x00, 0xFF, 0xCC}

// TAG is command type for request or response.
type TAG string

// Requests
const (
	SetName      TAG = "SET_NAME"
	CreateLobby  TAG = "CREATE_LOBBY"
	JoinLobby    TAG = "JOIN_LOBBY"
	LeaveLobby   TAG = "LEAVE_LOBBY"
	SetReadiness TAG = "SET_READINESS"
	StartMatch   TAG = "START_MATCH"
)

// Responses
const (
	Ok              TAG = "OK"
	Err             TAG = "ERR"
	JoinedLobby     TAG = "JOINED_LOBBY"
	LeftLobby       TAG = "LEFT_LOBBY"
	PlayerReadiness TAG = "PLAYER_READINESS"
	MatchStarted    TAG = "MATCH_STARTED"
)

// TAGs is a set of all Brave Rats protocol
var TAGs = map[TAG]struct{}{
	SetName:      {},
	CreateLobby:  {},
	JoinLobby:    {},
	LeaveLobby:   {},
	SetReadiness: {},
	StartMatch:   {},

	Ok:              {},
	Err:             {},
	JoinedLobby:     {},
	LeftLobby:       {},
	PlayerReadiness: {},
	MatchStarted:    {},
}
