package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main () {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	addr := ln.Addr()
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("connection from:", addr, "\n"))
		io.WriteString(conn, fmt.Sprint("Hello Andy\n", time.Now(), "\n"))

		conn.Close()
	}
}
