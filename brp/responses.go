package brp

import "strconv"

func NewRespOk(message string) []byte {
	resp := append([]byte(RespOk), ' ')
	resp = append(resp, message...)
	return append(resp, Ending...)
}

func NewRespErr(err error) []byte {
	resp := append([]byte(RespErr), ' ')
	resp = append(resp, err.Error()...)
	return append(resp, Ending...)
}

func NewRespInfo(info string) []byte {
	resp := append([]byte(RespInfo), ' ')
	resp = append(resp, info...)
	return append(resp, Ending...)
}

func NewRespLobby(ready bool, name string) []byte {
	resp := append([]byte(RespLobby), ' ')
	resp = append(resp, []byte(strconv.FormatBool(ready))...)
	resp = append(resp, ' ')
	resp = append(resp, []byte(name)...)
	resp = append(resp, Ending...)
	return resp
}
