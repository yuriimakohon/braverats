package main

import (
	"braverats/server"
	"flag"
)

func main() {
	port := flag.String("port", "3000", "port to listen on")
	server.NewServer().Start(*port)
}
