package server

import (
	"braverats/brp"
	"log"
)

func (c *client) logResponse(tag brp.TAG, message string) {
	log.Printf("[%s] reveives %s: %s\n", c.id, tag, message)
}

func (c *client) respOk(message string) {
	c.logResponse(brp.RespOk, message)
	_, err := c.conn.Write(brp.NewRespOk(message))
	c.handleWriteErr(brp.RespOk, err)
}

func (c *client) respErr(err error) {
	c.logResponse(brp.RespErr, err.Error())
	_, err = c.conn.Write(brp.NewRespErr(err))
	c.handleWriteErr(brp.RespErr, err)
}

func (c *client) respInfo(info string) {
	c.logResponse(brp.RespInfo, info)
	_, err := c.conn.Write(brp.NewRespInfo(info))
	c.handleWriteErr(brp.RespInfo, err)
}

func (c *client) respLobby(ready bool, name string) {
	resp := brp.NewRespLobby(ready, name)
	c.logResponse(brp.RespLobby, string(resp[:len(resp)-3]))
	_, err := c.conn.Write(resp)
	c.handleWriteErr(brp.RespLobby, err)
}
