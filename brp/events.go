package brp

import "strconv"

func NewEventJoinedLobby(name string) []byte {
	resp := append([]byte(EventJoinedLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func NewEventLeftLobby(name string) []byte {
	resp := append([]byte(EventLeftLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func NewEventLobbyClosed() []byte {
	return append([]byte(EventLobbyClosed), Ending...)
}

func NewEventPlayerReadiness(ready bool) []byte {
	resp := append([]byte(EventPlayerReadiness), ' ')
	resp = append(resp, []byte(strconv.FormatBool(ready))...)
	return append(resp, Ending...)
}

func NewEventMatchStarted() []byte {
	return append([]byte(EventMatchStarted), Ending...)
}
