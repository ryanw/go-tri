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
	t.DisableCanonical()
	t.DisableFlowControl()
	t.DisableInputProcessing()
	t.DisableCRtoNL()
}

func (t *Terminal) NormalMode() {
	t.EnableCRtoNL()
	t.EnableInputProcessing()
	t.EnableFlowControl()
	t.EnableCanonical()
	t.EnableEcho()
}

func (t *Terminal) SetLFlag(flag uint32) {
	termios := t.TermIOs()
	termios.Lflag |= flag
	t.SetTermIOs(termios)
}

func (t *Terminal) UnsetLFlag(flag uint32) {
	termios := t.TermIOs()
	termios.Lflag &^= flag
	t.SetTermIOs(termios)
}

func (t *Terminal) SetIFlag(flag uint32) {
	termios := t.TermIOs()
	termios.Iflag |= flag
	t.SetTermIOs(termios)
}

func (t *Terminal) UnsetIFlag(flag uint32) {
	termios := t.TermIOs()
	termios.Iflag &^= flag
	t.SetTermIOs(termios)
}

func (t *Terminal) EnableFlowControl() {
	t.SetIFlag(syscall.IXON)
}

func (t *Terminal) DisableFlowControl() {
	t.UnsetIFlag(syscall.IXON)
}

func (t *Terminal) EnableCRtoNL() {
	t.SetIFlag(syscall.ICRNL)
}

func (t *Terminal) DisableCRtoNL() {
	t.UnsetIFlag(syscall.ICRNL)
}

func (t *Terminal) EnableCanonical() {
	t.SetLFlag(syscall.ICANON)
}

func (t *Terminal) DisableCanonical() {
	t.UnsetLFlag(syscall.ICANON)
}

func (t *Terminal) EnableInputProcessing() {
	t.SetLFlag(syscall.IEXTEN)
}

func (t *Terminal) DisableInputProcessing() {
	t.UnsetLFlag(syscall.IEXTEN)
}

func (t *Terminal) EnableCtrlC() {
	t.SetLFlag(syscall.ISIG)
}

func (t *Terminal) DisableCtrlC() {
	t.UnsetLFlag(syscall.ISIG)
}

func (t *Terminal) EnableEcho() {
	t.SetLFlag(syscall.ECHO)
}

func (t *Terminal) DisableEcho() {
	t.UnsetLFlag(syscall.ECHO)
}
