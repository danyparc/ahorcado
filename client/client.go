package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8080")
	checkError(err)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("start"))
	checkError(err)
	//TODO: quitar espacios o declarar byte dinámico
	reply := make([]byte, 128)
	rl, err := conn.Read(reply)
	checkError(err)

	fmt.Println("RESPONSE:", string(reply[:rl]), "LEN:", len(reply[:rl]))

	handleResp(string(reply[:rl]), conn)

	reply = make([]byte, 20)
	_, err = conn.Read(reply)
	checkError(err)
	fmt.Println("RESPONSE:", string(reply), "LEN:", len(reply))

}

func handleResp(resp string, conn *net.TCPConn) {
	if resp == "setlevel" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Elige un nivel 1 | 2 | 3 : ")
		text, err := reader.ReadString('\n')
		checkError(err)
		fmt.Println(text[:1], len(text[:1]))
		conn.Write([]byte("level" + text[:1]))
		checkError(err)
		fmt.Println("enviado")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
