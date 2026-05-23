package main

import (
	"fmt"
	"io"
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

	fmt.Println("Connected.")

	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				fmt.Println("\nDisconnected.")
				break
			}

			fmt.Println("\nRead error:", err)
			break
		}

		fmt.Print(string(buffer[:n]))
	}
}
