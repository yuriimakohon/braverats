package server

import (
	"braverats/brp"
	"log"
)

func (c *client) ok(message string) {
	log.Printf("client %s receives OK: %s\n", c.id, message)
	_, err := c.conn.Write(brp.NewRespOk(message))
	c.handleWriteErr(brp.RespOk, err)
}

func (c *client) err(err error) {
	log.Printf("client %s receives ERR: %v\n", c.id, err)
	_, err = c.conn.Write(brp.NewRespErr(err))
	c.handleWriteErr(brp.RespErr, err)
}

func (c *client) info(info string) {
	log.Printf("client %s receives INFO: %s\n", c.id, info)
	_, err := c.conn.Write(brp.NewRespInfo(info))
	c.handleWriteErr(brp.RespInfo, err)
}
