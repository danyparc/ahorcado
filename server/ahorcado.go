package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
		checkError(err)
		request := make([]byte, 128)
		defer conn.Close()

		fmt.Println("Connection with", conn.RemoteAddr(), "stablished")

		readlen, err := conn.Read(request)
		checkError(err)
		fmt.Println("RECIBÍ", string(request[:readlen]))
		switch string(request[:readlen]) {
		case "start":
			fmt.Println("STARTED")
			conn.Write([]byte("setlevel"))
		case "level1":
			fmt.Println("level1")
			startGame(1, conn)
		case "level2":
			fmt.Println("level2")
			startGame(2, conn)
		case "level3":
			fmt.Println("level3")
			startGame(3, conn)
		default:
			fmt.Println("Invalid Message")
		}

	}

}

func startGame(lvl int, conn *net.TCPConn) {
	var word, clue string
	switch lvl {
	case 1:
		word = "perro"
	case 2:
		word = "biblioteca"
	case 3:
		word = "peritonitis"
	default:
		word = "perro"
	}

	for i := 0; i < len(word); i++ {
		clue = clue + "_"
	}

	_, err := conn.Write([]byte(clue))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
