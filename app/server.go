package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	//
	conn, con_err := l.Accept()
	if con_err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	response := []byte("HTTP/1.1 200 OK\r\n\r\n")
	_, res_err := conn.Write(response)
	if res_err != nil {
		fmt.Println("Error sending data", res_err.Error())
	}

}
