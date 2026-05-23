package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn    net.Conn
	logFile *os.File
	logging bool
}

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
		fmt.Println("Muddle commands:")
		fmt.Println("  /connect <host> <port>  Connect to a MUD")
		fmt.Println("  /log                    Toggle session logging")
		fmt.Println("  /quit                   Quit Muddle")
		fmt.Println("  //command               Send /command to the MUD")

	case "/connect":
		if len(fields) != 3 {
			fmt.Println("Usage: /connect <host> <port>")
			return
		}
		c.connect(fields[1], fields[2])

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

	_, err := fmt.Fprintln(c.conn, line)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	if c.logging && c.logFile != nil {
		fmt.Fprintln(c.logFile, "\n> "+line)
	}
}

func (c *Client) toggleLog() {
	if c.logging {
		c.logging = false

		if c.logFile != nil {
			c.logFile.Close()
			c.logFile = nil
		}

		fmt.Println("Logging stopped.")
		return
	}

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		fmt.Println("Could not create logs directory:", err)
		return
	}

	filename := "logs/session-" + time.Now().Format("2006-01-02-150405") + ".log"

	logFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Could not create log file:", err)
		return
	}

	c.logFile = logFile
	c.logging = true

	fmt.Println("Logging started:", filename)
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
