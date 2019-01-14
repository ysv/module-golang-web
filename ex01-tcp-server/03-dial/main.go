package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, "I dialed you bro\n")

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

