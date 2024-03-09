package main

import (
	"fmt"
	"time"
)

type ChannelServer struct {
	quitch chan struct {}
	msgch chan string
}


func (s *ChannelServer) FetchData() string {
	time.Sleep(time.Second * 4)
	return "return value"
}

func (s *ChannelServer) CreateMessage(msg string) {
	s.msgch <- msg
}

func NewServer() *ChannelServer {
	return &ChannelServer{
		msgch: make(chan string),
	}
}

func (s *ChannelServer) HandleMessage(message string) {
	fmt.Println("message received:", message)
}

func (s *ChannelServer) Start() {
	s.loop()
}

func (s *ChannelServer) loop() {
	mainloop:for {
		select {
		case <- s.quitch:
			break mainloop
		case msg:= <-s.msgch:
			s.HandleMessage(msg)
		default:
		}
	}
}


func ChannelMainFunc () string {
	server := NewServer()
	go server.Start()
	
	server.CreateMessage("Hello this is an other message in message channel!")
	time.Sleep(time.Second * 10)

	return "Done"
}
