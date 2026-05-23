package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	profile, err := loadProfile("profiles/default.yaml")
	if err != nil {
		fmt.Println("Profile error:", err)
		os.Exit(1)
	}

	client := &Client{
		profile: profile,
	}

	fmt.Println("Muddle starting. Type /help for commands.")
	fmt.Println("Loaded profile:", profile.Name)

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
