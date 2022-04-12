package main

import (
	"braverats/client"
	"flag"
)

func main() {
	remoteAddr := flag.String("addr", "localhost", "")
	flag.Parse()

	client.NewApp(*remoteAddr + ":6077").Start()
}
