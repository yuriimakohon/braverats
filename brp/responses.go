package brp

import (
	"bytes"
	"strconv"
)

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

func ParseRespLobby(packet Packet) (ready bool, nickname string, err error) {
	if packet.Tag != RespLobby {
		return ready, nickname, ErrPacketTagMismatch
	}

	args := bytes.Split(packet.Payload, []byte(" "))

	ready, err = strconv.ParseBool(string(args[0]))
	if err != nil {
		return ready, nickname, err
	}

	nickname = string(bytes.Join(args[1:], []byte(" ")))
	return ready, nickname, nil
}
