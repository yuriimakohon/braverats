package main

import (
	"braverats/client"
	"flag"
)

func main() {
	remoteAddr := flag.String("addr", "localhost:3000", "")
	flag.Parse()

	client.NewApp(*remoteAddr).Start()
}
