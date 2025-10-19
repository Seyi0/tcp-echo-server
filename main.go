package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// port used will be 9090
	
	if len(os.Args) <2{
		fmt.Println("for usage: go run main.go <port>")
		os.Exit(1)
	}

	// getting port from command line argument
	port := fmt.Sprintf(":%s", os.Args[1])

	net.Listen("tcp", port)
	listener, err := net.Listen("tcp", ":8080")

	if err != nil{
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("TCP server listening on port", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil{
			if err != io.EOF {
			fmt.Println("Error reading from connection:", err)
		}
			return
		}

		fmt.Printf("request: %s", bytes)
		line := fmt.Sprintf("Echo: %s", bytes)

		fmt.Printf("response: %s", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}