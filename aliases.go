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
		c.ui.Println("Usage: /alias <name> <value>")
		return
	}

	name := fields[1]
	value := strings.Join(fields[2:], " ")

	c.profile.Aliases[name] = value

	err := saveProfile("profiles/default.yaml", c.profile)
	if err != nil {
		c.ui.Println("Alias saved in memory, but profile save failed: " + err.Error())
		return
	}

	c.ui.Println(fmt.Sprintf("Alias added: %s -> %s", name, value))
}

func (c *Client) listAliases() {
	if len(c.profile.Aliases) == 0 {
		c.ui.Println("No aliases defined.")
		return
	}

	c.ui.Println("Aliases:")

	for name, value := range c.profile.Aliases {
		c.ui.Println(fmt.Sprintf("  %s -> %s", name, value))
	}
}
