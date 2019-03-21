package rcli

import (
	"fmt"
	"strings"
)

// AddHeader - add words to header
func (w *Window) AddHeader(header string) {
	w.clear(false)
	w.header = w.header + header
	w.lines = w.lines[1:]
	w.redraw()
}

// SetHeader - set new header
func (w *Window) SetHeader(header string) {
	w.clear(false)
	w.header = header
	w.lines = w.lines[1:]
	w.redraw()
}

// Println string line in window and move to next line
func (w *Window) Println(text string) {
	for _, s := range strings.Split(text, "\n") {
		w.println(s)
	}
}

// Print - and string to text in current window line
func (w *Window) Print(text string) {
	lines := strings.Split(text, "\n")

	cnt := len(lines)

	if cnt == 1 {
		w.print(text)
		return
	}

	for i := 0; i < cnt-1; i++ {
		w.println(lines[i])
	}
	w.print(lines[cnt-1])
}

// Input - print message and scan input
func (w *Window) Input(message string) (string, error) {
	message = strings.TrimSpace(message)
	if !strings.HasSuffix(message, " ") {
		message = message + " "
	}
	fmt.Print(message)
	var resp string
	_, err := fmt.Scanln(&resp)

	fmt.Print("\x1b[A")
	clearLine(len(message + resp))

	return resp, err
}

// Clear - clear window
func (w *Window) Clear() {
	w.clear(true)
	w.redraw()
}
