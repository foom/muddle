package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Could not enter raw terminal mode:", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	profile, err := loadProfile("profiles/default.yaml")
	if err != nil {
		fmt.Println("Profile error:", err)
		os.Exit(1)
	}

	ui := NewTerminalUI("> ")

	client := &Client{
		profile: profile,
		ui:      ui,
	}

	ui.Println("Muddle starting. Type /help for commands.")
	ui.Println("Loaded profile: " + profile.Name)
	ui.RedrawInput()

	reader := bufio.NewReader(os.Stdin)

	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			ui.Println("Input error: " + err.Error())
			break
		}

		switch ch {
		case '\r', '\n':
			line := ui.SubmitInput()

			if strings.HasPrefix(line, "/") {
				client.handleClientCommand(line)
			} else {
				client.sendToMud(line)
			}

		case 3: // Ctrl+C
			client.close()
			ui.Println("Goodbye.")
			return

		case 8, 127: // Backspace
			ui.Backspace()

		default:
			if ch >= 32 {
				ui.AppendRune(ch)
			}
		}
	}
}
