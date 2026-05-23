package main

import (
	"fmt"
	"strings"
)

func (c *Client) expandAlias(line string) string {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return line
	}

	firstWord := fields[0]

	replacement, found := c.profile.Aliases[firstWord]
	if !found {
		return line
	}

	if len(fields) == 1 {
		return replacement
	}

	rest := strings.Join(fields[1:], " ")

	return replacement + " " + rest
}

func (c *Client) handleAliasCommand(fields []string) {
	if len(fields) == 1 {
		c.listAliases()
		return
	}

	if len(fields) < 3 {
		fmt.Println("Usage: /alias <name> <value>")
		return
	}

	name := fields[1]
	value := strings.Join(fields[2:], " ")

	c.profile.Aliases[name] = value

	err := saveProfile("profiles/default.yaml", c.profile)
	if err != nil {
		fmt.Println("Alias saved in memory, but profile save failed:", err)
		return
	}

	fmt.Printf("Alias added: %s -> %s\n", name, value)
}

func (c *Client) listAliases() {
	if len(c.profile.Aliases) == 0 {
		fmt.Println("No aliases defined.")
		return
	}

	fmt.Println("Aliases:")

	for name, value := range c.profile.Aliases {
		fmt.Printf("  %s -> %s\n", name, value)
	}
}
