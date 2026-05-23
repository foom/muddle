package main

import (
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
			c.ui.Println("Usage: /connect <host> <port>")
			return
		}
		c.connect(fields[1], fields[2])

	case "/alias":
		c.handleAliasCommand(fields)

	case "/log":
		c.toggleLog()

	case "/quit":
		c.close()
		c.ui.Println("Goodbye.")
		os.Exit(0)

	default:
		c.ui.Println("Unknown Muddle command. Type /help.")
	}
}

func (c *Client) showHelp() {
	c.ui.Println("Muddle commands:")
	c.ui.Println("  /connect <host> <port>  Connect to a MUD")
	c.ui.Println("  /alias                  List aliases")
	c.ui.Println("  /alias <name> <value>   Add or update alias")
	c.ui.Println("  /log                    Toggle session logging")
	c.ui.Println("  /quit                   Quit Muddle")
	c.ui.Println("  //command               Send /command to the MUD")
}
