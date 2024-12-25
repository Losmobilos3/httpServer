package main

import (
	httpServer "httpServer/server"
	"net"
)

func main() {
	serv := httpServer.CreateServer("localhost", 80)

	httpServer.AddHandler(serv, "/test", test)

	httpServer.StartServer(serv)
}

func test(conn net.Conn) string {
	return "Test"
}
