package brp

func NewRespOk() []byte {
	return append([]byte(RespOk), Ending...)
}

func NewRespErr(err error) []byte {
	resp := append([]byte(RespErr), ' ')
	resp = append(resp, []byte(err.Error())...)
	return append(resp, Ending...)
}

func NewRespInfo(info string) []byte {
	resp := append([]byte(RespInfo), ' ')
	resp = append(resp, []byte(info)...)
	return append(resp, Ending...)
}
