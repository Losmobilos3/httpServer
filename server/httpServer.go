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
	handlers map[string](func(net.Conn) string)
}

func CreateServer(address string, port int) httpServer {
	listener, err := net.Listen(
		"tcp",
		"localhost:80",
	)
	if err != nil {
		fmt.Println(err)
	}

	handlerMap := map[string](func(net.Conn) string){}

	return httpServer{address: address, port: port, listener: listener, handlers: handlerMap}
}

func AddHandler(server httpServer, endpoint string, function func(net.Conn) string) {
	server.handlers[endpoint] = function
}

func StartServer(server httpServer) {
	fmt.Println("👌 Started http-server.")
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		handleRequest(server, conn)

		conn.Close()
	}

}

func handleRequest(server httpServer, conn net.Conn) {
	msg := make([]byte, maxRequestSize)

	conn.Read(msg)

	httpStr := string(msg[:])

	_, endpoint := getEndpointInformation(httpStr)

	handler, ok := server.handlers[endpoint]

	if !ok {
		fmt.Println("Someone tried requesting the following endpoint, which is not defined:", endpoint)
		return
	}

	responseStr := handler(conn)

	responseBytes := []byte(responseStr)

	conn.Write(responseBytes)
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
