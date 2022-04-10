package brp

import "strconv"

func ReqSetName(name string) []byte {
	req := append([]byte(TagSetName), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqCreateLobby(name string) []byte {
	req := append([]byte(TagCreateLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqJoinLobby(name string) []byte {
	req := append([]byte(TagJoinLobby), ' ')
	req = append(req, []byte(name)...)
	return append(req, Ending...)
}

func ReqLeaveLobby() []byte {
	req := append([]byte(TagLeaveLobby), ' ')
	return append(req, Ending...)
}

func ReqSetReadiness(ready bool) []byte {
	req := append([]byte(TagSetReadiness), ' ')
	req = append(req, []byte(strconv.FormatBool(ready))...)
	return append(req, Ending...)
}

func ReqStartMatch() []byte {
	return append([]byte(TagStartMatch), Ending...)
}
