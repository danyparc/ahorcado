package main

import (
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

	reply := make([]byte, 128)
	_, err = conn.Read(reply)
	checkError(err)

	fmt.Println("RESPONSE:", string(reply), "LEN:", len(reply))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
