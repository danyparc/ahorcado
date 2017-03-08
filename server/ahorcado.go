package main

import (
	"fmt"
	"log"
	"net"
)

//Server
func main() {
	//var word string

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println("Listening in port 8080")
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		request := make([]byte, 128)
		defer conn.Close()

		fmt.Println("Connection with", conn.RemoteAddr(), "stablished")

		readlen, err := conn.Read(request)
		if err != nil {
			log.Fatal(err)
		}

		switch string(request[:readlen]) {
		case "start":
			fmt.Println("STARTED")
			conn.Write([]byte("setlevel"))
		case "level1":
			startGame(1, conn)
		case "level2":
			startGame()
		case "level3":
			startGame()
		default:
			fmt.Println("Invalid Message")
		}

	}

}

func startGame(lvl int, conn *net.TCPConn) {
	var word string
	switch lvl {
	case 1:
		word = "perro"
	case 2:
		word = "biblioteca"
	case 3:
		word = "peritonitis"
	default:
		word = perro
	}
}
