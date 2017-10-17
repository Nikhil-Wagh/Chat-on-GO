package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) bool {
	found := false
	if err != nil {
		fmt.Println(err)
		found = true
	}
	fmt.Printf("%s\n", message)
	return found
}

type Clients struct {
	name string
	conn net.Conn
}

type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

var clients = make(map[string]net.Conn)
var broadcast = make(chan Message)

func main() {

	go generateResponses()

	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is running.")

	for {
		conn, err := ln.Accept()
		//Register new connection
		//name := make([]byte, 1024)
		var name string
		if !check(err, "Accepted Connection") {
			conn.Write([]byte("Write your name: "))
			buf := bufio.NewReader(conn)
			var err error
			name, err = buf.ReadString('\n')
			check(err, "Name accepted.")
			name = name[:len([]rune(name))-2]
			// n, err := conn.Read(name)
			// check(err, "Connection Registered.")
			// fmt.Printf("%d bytes of name read", n)
			// var u_name string
			// u_name = string(name)
			conn.Write([]byte("Hello, " + name))
			clients[name] = conn
		}

		go func() {
			var msg Message

			buf := bufio.NewReader(conn)

			for {
				conn.Write([]byte("Write your message in format.\n"))
				msg.From, err = buf.ReadString(' ')
				msg.To, err = buf.ReadString(' ')
				msg.Text, err = buf.ReadString('\n')
				//fmt.Println(msg)
				if err != nil {
					fmt.Println("Client Disconnected.")
					delete(clients, string(name))
					break
				}
				msg.From = msg.From[:len([]rune(msg.From))-1]
				msg.To = msg.To[:len([]rune(msg.To))-1]
				fmt.Println(msg.From, msg.To)
				broadcast <- msg
			}
		}()
	}
}

func generateResponses() {
	for {
		clientJob := <-broadcast
		from := clientJob.From
		to := clientJob.To
		text := clientJob.Text
		fmt.Println(clientJob)
		for name, conn := range clients {
			fmt.Println(name, conn)
			if name == to {
				fmt.Println("OK")
			}
		}
		clients[to].Write([]byte(from + ": " + text))
	}
}
