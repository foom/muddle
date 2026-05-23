package main

import (
	"bufio"
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

	address := os.Args[1] + ":" + os.Args[2]

	fmt.Printf("Connecting to %s...\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected.")

	go readFromMud(conn)

	readFromKeyboard(conn)
}

func readFromMud(conn net.Conn) {
	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nDisconnected.")
			} else {
				fmt.Println("\nRead error:", err)
			}
			os.Exit(0)
		}

		fmt.Print(string(buffer[:n]))
	}
}

func readFromKeyboard(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Fprintln(conn, line)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}
