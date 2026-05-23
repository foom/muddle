package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	conn    net.Conn
	logFile *os.File
	logging bool
	profile Profile
}

func (c *Client) connect(host, port string) {
	if c.conn != nil {
		fmt.Println("Already connected. Use /quit for now.")
		return
	}

	address := host + ":" + port
	fmt.Printf("Connecting to %s...\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}

	c.conn = conn
	fmt.Println("Connected.")

	go c.readFromMud()
}

func (c *Client) readFromMud() {
	buffer := make([]byte, 4096)

	for {
		n, err := c.conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nDisconnected.")
			} else {
				fmt.Println("\nRead error:", err)
			}

			c.conn = nil
			return
		}

		data := buffer[:n]

		fmt.Print(string(data))

		if c.logging && c.logFile != nil {
			c.logFile.Write(data)
		}
	}
}

func (c *Client) sendToMud(line string) {
	if c.conn == nil {
		fmt.Println("Not connected. Use /connect <host> <port>.")
		return
	}

	line = c.expandAlias(line)

	_, err := fmt.Fprintln(c.conn, line)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	if c.logging && c.logFile != nil {
		fmt.Fprintln(c.logFile, "\n> "+line)
	}
}

func (c *Client) close() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}

	if c.logFile != nil {
		c.logFile.Close()
		c.logFile = nil
	}

	c.logging = false
}
