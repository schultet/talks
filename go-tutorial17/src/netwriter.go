package main

import (
	"fmt"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", "localhost:3333")
	defer l.Close()

	for {
		conn, _ := l.Accept()
		fmt.Fprintln(conn, "Hello!") // HL
		conn.Close()
	}
}
