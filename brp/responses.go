package brp

func NewRespOk(message string) []byte {
	resp := append([]byte(RespOk), ' ')
	resp = append(resp, []byte(message)...)
	return append(resp, Ending...)
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
