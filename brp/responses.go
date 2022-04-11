package brp

func NewRespOk() []byte {
	return append([]byte(RespOk), Ending...)
}

func NewRespErr(err error) []byte {
	resp := append([]byte(RespErr), ' ')
	resp = append(resp, []byte(err.Error())...)
	return append(resp, Ending...)
}
