package main

import (
	"os"
	"time"
)

func (c *Client) toggleLog() {
	if c.logging {
		c.logging = false

		if c.logFile != nil {
			c.logFile.Close()
			c.logFile = nil
		}

		c.ui.Println("Logging stopped.")
		return
	}

	err := os.MkdirAll("logs", 0755)
	if err != nil {
		c.ui.Println("Could not create logs directory: " + err.Error())
		return
	}

	filename := "logs/session-" + time.Now().Format("2006-01-02-150405") + ".log"

	logFile, err := os.Create(filename)
	if err != nil {
		c.ui.Println("Could not create log file: " + err.Error())
		return
	}

	c.logFile = logFile
	c.logging = true

	c.ui.Println("Logging started: " + filename)
}
