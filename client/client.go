package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var tcpAddr *net.TCPAddr
var clue string

func main() {
	var err error
	tcpAddr, err = net.ResolveTCPAddr("tcp4", ":8080")
	checkError(err)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("start"))
	checkError(err)
	reply := make([]byte, 128)
	rl, err := conn.Read(reply)
	checkError(err)
	conn.Close()

	fmt.Println("RESPONSE:", string(reply[:rl]), "LEN:", len(reply[:rl]))

	handleResp(string(reply[:rl]))

}

func sentTo(address *net.TCPAddr, content string) string {
	conn, err := net.DialTCP("tcp4", nil, address)
	checkError(err)
	_, err = conn.Write([]byte(content))
	checkError(err)
	reply := make([]byte, 128)
	rl, err := conn.Read(reply)
	checkError(err)
	conn.CloseRead()
	conn.Close()
	return string(reply[:rl])
}

func handleResp(resp string) {

	if resp == "setlevel" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Elige un nivel 1 | 2 | 3 : ")
		text, err := reader.ReadString('\n')
		checkError(err)
		reply := sentTo(tcpAddr, "level"+text[:1])
		fmt.Println("\nLa pista es: ", reply)
		clue = reply
	}
	if resp == "GANASTE" {
		fmt.Println("HAS GANADO")
	}
	runGame()

}
func runGame() {
	for strings.Contains(clue, "_") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEscribe una letra: ")
		letr, err := reader.ReadString('\n')
		checkError(err)
		reply := sentTo(tcpAddr, "letr:"+letr[:1])
		fmt.Println(reply)

	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
