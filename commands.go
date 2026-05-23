package main

import (
	"fmt"
	"os"
	"strings"
)

func (c *Client) handleClientCommand(line string) {
	if strings.HasPrefix(line, "//") {
		c.sendToMud(line[1:])
		return
	}

	fields := strings.Fields(line)
	if len(fields) == 0 {
		return
	}

	switch fields[0] {
	case "/help":
		c.showHelp()

	case "/connect":
		if len(fields) != 3 {
			fmt.Println("Usage: /connect <host> <port>")
			return
		}
		c.connect(fields[1], fields[2])

	case "/alias":
		c.handleAliasCommand(fields)

	case "/log":
		c.toggleLog()

	case "/quit":
		c.close()
		fmt.Println("Goodbye.")
		os.Exit(0)

	default:
		fmt.Println("Unknown Muddle command. Type /help.")
	}
}

func (c *Client) showHelp() {
	fmt.Println("Muddle commands:")
	fmt.Println("  /connect <host> <port>  Connect to a MUD")
	fmt.Println("  /alias                  List aliases")
	fmt.Println("  /alias <name> <value>   Add or update alias")
	fmt.Println("  /log                    Toggle session logging")
	fmt.Println("  /quit                   Quit Muddle")
	fmt.Println("  //command               Send /command to the MUD")
}
