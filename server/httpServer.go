package httpServer

import (
	"fmt"
	"net"
	"strings"
)

// Max size of request
var maxRequestSize int = 8196 // bytes

// Request types
const (
	GET     int = 0
	POST        = 1
	PUT         = 2
	DELETE      = 3
	HEAD        = 4
	CONNECT     = 5
	OPTIONS     = 6
	TRACE       = 7
	PATCH       = 8
)

type httpServer struct {
	address  string
	port     int
	listener net.Listener
	handlers map[string]func(net.Conn)
}

func createServer(address string, port int) httpServer {
	listener, err := net.Listen(
		"tcp",
		"localhost:80",
	)
	if err != nil {
		fmt.Println(err)
	}

	handlerMap := map[string]func(net.Conn){}

	return httpServer{address: address, port: port, listener: listener, handlers: handlerMap}
}

func addHandler(server httpServer, endpoint string, function func(net.Conn)) {
	server.handlers[endpoint] = function
}

func startServer(server httpServer) {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		handleRequest(server, conn)
	}

}

func handleRequest(server httpServer, conn net.Conn) {
	msg := make([]byte, maxRequestSize)

	conn.Read(msg)

	httpStr := string(msg[:])

	_, endpoint := getEndpointInformation(httpStr)

	handler := server.handlers[endpoint]

	handler(conn)
}

func getEndpointInformation(httpStr string) (int, string) {
	lines := strings.Split(httpStr, "\n")
	firstLineWords := strings.Split(lines[0], " ")
	requestType := mapStringToRequestType(firstLineWords[0])
	endpoint := firstLineWords[1]

	return requestType, endpoint
}

func mapStringToRequestType(str string) int {
	switch str {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	}
	return -1
}
