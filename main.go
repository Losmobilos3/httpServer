package main

import (
	httpServer "httpServer/server"
	"net"
	"os"
)

func main() {
	serv := httpServer.CreateServer("localhost", 80)

	httpServer.AddHandler(serv, "/test", test)
	httpServer.AddHandler(serv, "/test2", test2)

	httpServer.StartServer(serv)
}

func test2(conn net.Conn) string {
	return "banan"
}

func test(conn net.Conn) string {
	contentBytes, err := os.ReadFile("test.html")

	contentStr := string(contentBytes)
	if err != nil {
		return "Error"
	}

	return contentStr
}
