package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

func (t *Terminal) TermIOs() syscall.Termios {
	var termios syscall.Termios
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&termios)),
	)
	return termios
}

func (t *Terminal) SetTermIOs(termios syscall.Termios) {
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(&termios)),
	)
}

func (t *Terminal) RawMode() {
	t.DisableEcho()
}

func (t *Terminal) NormalMode() {
	t.EnableEcho()
}

func (t *Terminal) DisableEcho() {
	termios := t.TermIOs()
	termios.Lflag &^= syscall.ECHO
	t.SetTermIOs(termios)
}

func (t *Terminal) EnableEcho() {
	termios := t.TermIOs()
	termios.Lflag += syscall.ECHO
	t.SetTermIOs(termios)
}
