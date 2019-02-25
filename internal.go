package rcli

import (
	"fmt"
	"strings"
)

func (w *Window) println(text string) {
	cur := len(w.lines) - 1
	fmt.Println(text)
	w.lines[cur] = w.lines[cur] + text
	w.lines = append(w.lines, "")
}

func (w *Window) print(text string) {
	cur := len(w.lines) - 1
	fmt.Print(text)
	w.lines[cur] = w.lines[cur] + text
}

func (w *Window) redraw() {
	count := len(w.lines)

	if count == 0 {
		w.lines = append(w.lines, w.header)
		w.lines = append(w.lines, "")
		count = 2
	} else {
		if w.lines[0] != w.header {
			w.lines = append([]string{w.header}, w.lines...)
		}
		count = len(w.lines)
	}

	for i := 0; i < count; i++ {
		fmt.Print(w.lines[i])
		if i != count-1 {
			fmt.Println()
		}
	}
}

func (w *Window) clear(delete bool) {
	count := len(w.lines)

	if count == 0 {
		return
	}

	if count == 1 {
		clearLine(len(w.lines[count-1]))
		if delete {
			w.lines = nil
		}
		return
	}

	clearLine(len(w.lines[count-1]))
	count--

	for count > 0 {
		fmt.Print("\x1b[A")
		clearLine(len(w.lines[count-1]))
		count--
	}
	if delete {
		w.lines = nil
	}
}

func clearLine(length int) {
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", length))
	fmt.Print(strings.Repeat("\b", length))
}
