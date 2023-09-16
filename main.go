package main

import (
	"github.com/xenedium/xhomer/server"
)

func main() {
	server := server.NewServer(":8080")
	server.Run()
}
