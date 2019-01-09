package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		io.WriteString(conn, "\nTCP connection opened\n")
		fmt.Fprintln(conn, "Write using fmt" )

		conn.Close()
	}
}
