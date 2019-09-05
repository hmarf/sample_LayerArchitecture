package main

import (
	"flag"

	"github.com/hmarf/sample_layer/application/server"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":9000", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
