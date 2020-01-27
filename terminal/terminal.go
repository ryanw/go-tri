package terminal

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type WinSize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

type Terminal struct {
	width, height int
	stdout        bufio.Writer
	stdin         bufio.Reader
}

func NewTerminal() Terminal {
	term := Terminal{
		width:  16,
		height: 16,
		stdout: *bufio.NewWriterSize(os.Stdout, 4096),
		stdin:  *bufio.NewReaderSize(os.Stdin, 64),
	}
	term.UpdateSize()
	return term
}

// Update our size to match the real TTY session
func (t *Terminal) UpdateSize() {
	var winSize WinSize
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(&winSize)),
	)

	t.width = int(winSize.cols)
	t.height = int(winSize.rows)
}

func (t *Terminal) Size() (int, int) {
	return t.width, t.height
}

func (t *Terminal) Width() int {
	return t.width
}

func (t *Terminal) Height() int {
	return t.height
}

func (t *Terminal) Draw(callback func()) {
	callback()
	t.Flush()
}

func (t *Terminal) Flush() {
	t.stdout.Flush()
}

func (t *Terminal) Write(format string, a ...interface{}) {
	fmt.Fprintf(&t.stdout, format, a...)
}

func (t *Terminal) WriteRune(char rune) {
	fmt.Fprint(&t.stdout, string(char))
}
