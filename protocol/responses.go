package protocol

import "strconv"

func RespOk() []byte {
	return append([]byte(Ok), Ending...)
}

func RespErr(err error) []byte {
	resp := append([]byte(Err), ' ')
	resp = append(resp, []byte(err.Error())...)
	return append(resp, Ending...)
}

func RespJoinedLobby(name string) []byte {
	resp := append([]byte(JoinedLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func RespLeftLobby(name string) []byte {
	resp := append([]byte(LeftLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func RespPlayerReadiness(ready bool) []byte {
	resp := append([]byte(PlayerReadiness), ' ')
	resp = append(resp, []byte(strconv.FormatBool(ready))...)
	return append(resp, Ending...)
}

func RespMatchStarted() []byte {
	return append([]byte(MatchStarted), Ending...)
}
