package main

import "fmt"

/**
 * https://www.bytesizego.com/blog/functional-options-pattern-golang
 */
type Server struct {
	host     string
	port     int
	protocol string
}

func NewServer(host string, opts ...ServerOption) *Server {
	server := &Server{
		host:     host,
		port:     443,
		protocol: "https",
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

type ServerOption func(*Server)

func withPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func witProtocol(protocol string) ServerOption {
	return func(s *Server) {
		s.protocol = protocol
	}
}

func main() {

	server1 := NewServer("localhost")

	server2 := NewServer("localhsot", withPort(8080), witProtocol("http"))

	fmt.Printf("server is: %+v\n", server1)
	fmt.Printf("server is: %+v\n", server2)
}
