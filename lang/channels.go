package main

import "fmt"

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) Start() {
	fmt.Println("Server Starting...")
	s.loop()
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Got message", msg)

func main() {

	server := newServer()
	go server.Start()
	for i := 0; i < 100; i++ {
		server.sendMessage(fmt.Sprintf("handle this number %d", i))
	}
	time.Sleep(time.Second * 5)
}
