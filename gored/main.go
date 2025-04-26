package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Listen to the port 6379
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port 6379")

	// Accept connexion with port
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to port 6379")

	// Loop
	for {
		buff := make([]byte, 1024)

		// Read msg from client
		_, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		// Send back default response for now
		conn.Write([]byte("+OK\r\n"))
	}

	// Close connexion anyway once finished
	defer conn.Close()
}