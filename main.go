package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: muddle <host> <port>")
		fmt.Println("Example: muddle torilmud.com 9999")
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]

	fmt.Printf("Muddle connecting to %s:%s...\n", host, port)
}
