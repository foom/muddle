package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: muddle <host> <port>")
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]

	address := host + ":" + port

	fmt.Printf("Connecting to %s...\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection error:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connected successfully.")
}
