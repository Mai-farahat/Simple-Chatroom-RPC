package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

func (c *ChatServer) SendMessage(msg string, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.messages = append(c.messages, msg)
	*reply = append([]string(nil), c.messages...)

	fmt.Printf("📩 New message received: %s\n", msg)
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	fmt.Println("🚀 Chat server running on port 1234...")
	fmt.Println("💬 Waiting for friends to join the chat...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
