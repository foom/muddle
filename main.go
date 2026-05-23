package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	client := &Client{}

	fmt.Println("Muddle starting. Type /help for commands.")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "/") {
			client.handleClientCommand(line)
		} else {
			client.sendToMud(line)
		}
	}
}
