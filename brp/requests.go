package brp

import "strconv"

func NewReqSetName(name string) []byte {
	req := append([]byte(ReqSetName), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func NewReqCreateLobby(name string) []byte {
	req := append([]byte(ReqCreateLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func NewReqJoinLobby(name string) []byte {
	req := append([]byte(ReqJoinLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func NewReqLeaveLobby() []byte {
	req := append([]byte(ReqLeaveLobby), ' ')
	return append(req, Ending...)
}

func NewReqSetReadiness(ready bool) []byte {
	req := append([]byte(ReqSetReadiness), ' ')
	req = append(req, []byte(strconv.FormatBool(ready))...)
	return append(req, Ending...)
}

func NewReqStartMatch() []byte {
	return append([]byte(ReqStartMatch), Ending...)
}
