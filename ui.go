package main

import (
	"fmt"
	"sync"
)

type TerminalUI struct {
	mu     sync.Mutex
	prompt string
	input  []rune
}

func NewTerminalUI(prompt string) *TerminalUI {
	return &TerminalUI{
		prompt: prompt,
		input:  []rune{},
	}
}

func (ui *TerminalUI) Print(text string) {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	ui.clearInputLine()
	fmt.Print(text)

	if len(text) > 0 && text[len(text)-1] != '\n' && text[len(text)-1] != '\r' {
		fmt.Print("\r\n")
	}

	ui.redrawInput()
}

func (ui *TerminalUI) Println(text string) {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	ui.clearInputLine()
	fmt.Print(text)
	fmt.Print("\r\n")
	ui.redrawInput()
}

func (ui *TerminalUI) AppendRune(ch rune) {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	ui.input = append(ui.input, ch)
	ui.redrawInput()
}

func (ui *TerminalUI) Backspace() {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	if len(ui.input) > 0 {
		ui.input = ui.input[:len(ui.input)-1]
	}

	ui.redrawInput()
}

func (ui *TerminalUI) SubmitInput() string {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	line := string(ui.input)
	ui.input = []rune{}

	ui.clearInputLine()
	fmt.Print("\r\n")

	return line
}

func (ui *TerminalUI) RedrawInput() {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	ui.redrawInput()
}

func (ui *TerminalUI) clearInputLine() {
	fmt.Print("\r\033[2K")
}

func (ui *TerminalUI) redrawInput() {
	ui.clearInputLine()
	fmt.Print(ui.prompt)
	fmt.Print(string(ui.input))
}
