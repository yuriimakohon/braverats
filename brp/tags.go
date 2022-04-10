package brp

// TAG is type of request or response in BRP protocol.
type TAG string

// Request BRP TAGs.
const (
	TagSetName      TAG = "SET_NAME"
	TagCreateLobby  TAG = "CREATE_LOBBY"
	TagJoinLobby    TAG = "JOIN_LOBBY"
	TagLeaveLobby   TAG = "LEAVE_LOBBY"
	TagSetReadiness TAG = "SET_READINESS"
	TagStartMatch   TAG = "START_MATCH"
)

// Response BRP TAGs.
const (
	TagOk              TAG = "OK"
	TagErr             TAG = "ERR"
	TagJoinedLobby     TAG = "JOINED_LOBBY"
	LeftLobby          TAG = "LEFT_LOBBY"
	TagPlayerReadiness TAG = "PLAYER_READINESS"
	TagMatchStarted    TAG = "MATCH_STARTED"
)

// tags is a set of all BRP tags
var tags = map[TAG]struct{}{
	TagSetName:      {},
	TagCreateLobby:  {},
	TagJoinLobby:    {},
	TagLeaveLobby:   {},
	TagSetReadiness: {},
	TagStartMatch:   {},

	TagOk:              {},
	TagErr:             {},
	TagJoinedLobby:     {},
	LeftLobby:          {},
	TagPlayerReadiness: {},
	TagMatchStarted:    {},
}
