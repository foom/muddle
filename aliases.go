package main

import "strings"

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
