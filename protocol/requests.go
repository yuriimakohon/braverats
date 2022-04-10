package protocol

import "strconv"

func ReqSetName(name string) []byte {
	req := append([]byte(SetName), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqCreateLobby(name string) []byte {
	req := append([]byte(CreateLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqJoinLobby(name string) []byte {
	req := append([]byte(JoinLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqLeaveLobby() []byte {
	req := append([]byte(LeaveLobby), ' ')
	return append(req, Ending...)
}

func ReqSetReadiness(ready bool) []byte {
	req := append([]byte(SetReadiness), ' ')
	req = append(req, []byte(strconv.FormatBool(ready))...)
	return append(req, Ending...)
}

func ReqStartMatch() []byte {
	return append([]byte(StartMatch), Ending...)
}
