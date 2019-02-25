package rcli

import (
	"fmt"
	"log"
	"os"
)

// Window - represent application terminal window
type Window struct {
	header string

	lines []string
}

// NewWindow - create new window object
func NewWindow(header string) *Window {
	fmt.Print("\x1B[2J\x1B[H")

	w := &Window{
		header: header,
		lines:  make([]string, 2, 2),
	}

	if w.header == "" {
		w.header = os.Args[0]
	}

	w.lines[0] = w.header

	w.redraw()

	return w
}

// Close - close window
func (w *Window) Close() {
	w.clear(true)
	w = nil
}

// Fatal - fatal end execusion
func (w *Window) Fatal(v ...interface{}) {
	w.clear(true)
	log.Fatal(v...)
}
