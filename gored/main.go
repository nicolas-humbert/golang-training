package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// #region Conn
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

	// -------------------- TEMP -------------------- 
	input := "$15\r\nNicolas Humbert\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()
	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	var strSizeStr string
	strSizeBuff := bytes.NewBufferString(strSizeStr)
	
	for {
		bInt, _ := reader.ReadByte()

		if bInt == '\r' {
			reader.ReadByte()
			break;
		} else {
			_, err := strconv.ParseInt(string(bInt), 10, 64)
			if err != nil {
				fmt.Println("Invalid format, expecting size of the input")
			} else {
				strSizeBuff.WriteByte(bInt)
			}
		}
	}

	strSize, err := strconv.ParseInt(strSizeBuff.String(), 10, 64)
	if err != nil {
		fmt.Println("Invalid format for the input size")
	}

	// size, _ := reader.ReadByte()
	// strSize, err := strconv.ParseInt(string(size), 10, 64)
	// if err != nil {
		// fmt.Println("Invalid format, expecting size of the entry")
	// }

	// // Consume \r\n
	// reader.ReadByte()
	// reader.ReadByte()

	// Consume name
	name :=  make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))

	// #region Loop
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