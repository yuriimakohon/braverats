package server

import (
	"braverats/brp"
	"log"
)

func (c *client) err(err error) {
	log.Printf("client %s caused error: %v\n", c.id, err)
	_, err = c.conn.Write(brp.NewRespErr(err))
	c.handleWriteErr(brp.RespErr, err)
}

func (c *client) ok() {
	_, err := c.conn.Write(brp.NewRespOk())
	c.handleWriteErr(brp.RespOk, err)
}
