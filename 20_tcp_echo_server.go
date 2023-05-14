package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %#v", err)
			continue
		}
		go process(client)
	}
}

func process(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	for {
		buf, err := reader.ReadByte()
		if err != nil {
			break
		}

		_, err = connection.Write([]byte{buf})
		if err != nil {
			break
		}
	}
}
