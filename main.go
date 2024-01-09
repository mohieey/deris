package main

import (
	"fmt"
	"net"
)

const PORT = ":6379"

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("+OK\r\n"))
}
