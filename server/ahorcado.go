package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var word, clue string

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

	for {
		fmt.Println("Listening in port 8080")
		conn, err := l.AcceptTCP()
		checkError(err)
		fmt.Println("Connection with", conn.RemoteAddr(), "stablished")

		go handler(conn)

	}

}

func handler(conn *net.TCPConn) {
	request := make([]byte, 128)

	readlen, err := conn.Read(request)
	checkError(err)
	fmt.Println("RECIBÍ", string(request[:readlen]), "LEN", len(request[:readlen]))

	if string(request[:readlen-1]) == "letr:" {
		fmt.Println(string(request[readlen-1]))
		if validate(string(request[readlen-1])) {
			conn.Write([]byte(clue))
			fmt.Println("Cliente descubrió", string(request[readlen-1]))
			return
		} else {
			conn.Write([]byte(clue))
			fmt.Println("Cliente falló", string(request[readlen-1]))
		}

	}

	switch string(request[:readlen]) {
	case "start":
		conn.Write([]byte("setlevel"))
		fmt.Println("GAME STARTED")
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

	conn.Close()
}

func startGame(lvl int, conn *net.TCPConn) {
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
	conn.CloseWrite()
}

func validate(c string) bool {
	if strings.ContainsAny(word, c) {
		for l := 0; l < len(word); l++ {
			if string(word[l]) == c {
				clue = clue[:l] + c + clue[l+1:]
			}
		}
		if !strings.ContainsAny(clue, "_") {
			clue = "GANASTE"
		}
		return true
	}
	return false
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//os.Exit(1)
	}
}
