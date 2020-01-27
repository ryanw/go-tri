package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type InputEventType uint8
type MouseAction uint8
type MouseButton uint8

const (
	KeyEvent InputEventType = iota
	MouseEvent
)

const (
	MouseLeft MouseButton = iota
	MouseRight
	MouseMiddle
	MouseScrollUp
	MouseScrollDown
)

const (
	MouseDown MouseAction = iota
	MouseUp
	MouseDrag
	MouseMove
)

type InputEvent struct {
	EventType   InputEventType
	Key         rune
	MouseButton MouseButton
	MouseAction MouseAction
	MouseX      int
	MouseY      int
}

// Waits for the next user input event. Either mouse or keyboard
func (t *Terminal) NextEvent() InputEvent {
	b := nextByte(&t.stdin)
	if t.stdin.Buffered() > 0 && b == '\x1b' && isCtrlChar(peekByte(&t.stdin)) {
		return NewInputEventFromCtrlChar(&t.stdin)
	}

	return NewInputEventFromKey(b)
}

func NewInputEventFromKey(b byte) InputEvent {
	return InputEvent{
		EventType: KeyEvent,
		Key:       rune(b),
	}
}
func NewInputEventFromKeyboard(stream *bufio.Reader) InputEvent {
	return InputEvent{
		EventType: KeyEvent,
		Key:       nextRune(stream),
	}
}

func NewInputEventFromCtrlChar(stream *bufio.Reader) InputEvent {
	b := nextRune(stream)
	switch b {
	case '[':
		return NewInputEventFromCtrlSeq(stream)

	default:
		// Unsupported control char
		return InputEvent{}
	}

}

// Parse a CSI string into an InputEvent. Assumes CSI prefix has been removed.
// See: http://www.xfree86.org/4.7.0/ctlseqs.html
func NewInputEventFromCtrlSeq(stream *bufio.Reader) InputEvent {
	b := nextByte(stream)
	switch b {
	case 'M':
		return NewInputEventFromX10Mouse(stream)

	case '<':
		return NewInputEventFromSGRMouse(stream)

	default:
		return NewInputEventFromKey(b)
	}
}

func NewInputEventFromX10Mouse(stream *bufio.Reader) InputEvent {
	button := nextByte(stream) - 32
	x := int(nextByte(stream)) - 33
	y := int(nextByte(stream)) - 33

	// FIXME Support MouseUp
	action := MouseDown

	return InputEvent{
		EventType:   MouseEvent,
		MouseButton: MouseButton(button),
		MouseAction: action,
		MouseX:      x,
		MouseY:      y,
	}
}

func NewInputEventFromSGRMouse(stream *bufio.Reader) InputEvent {
	button := nextByte(stream) - 48
	action := MouseDown
	x := 0
	y := 0

	if button == 6 {
		// Mouse wheel
		button = nextByte(stream) - 48
	}

	scan := bufio.NewScanner(stream)
	scan.Split(bufio.ScanRunes)

	numStr := ""
	tokenIdx := 0
ScanLoop:
	for scan.Scan() {
		chr := scan.Text()[0]
		// Convert read ASCII number into a real number
		if chr == ';' || chr == 'm' || chr == 'M' {
			i, _ := strconv.Atoi(numStr)
			if tokenIdx == 1 {
				x = i - 1
			}
			if tokenIdx == 2 {
				y = i - 1
			}
			numStr = ""
			tokenIdx += 1
		}

		switch chr {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numStr += string(chr)

		// End of sequence
		case 'm':
			if button == 3 {
				action = MouseMove
			} else {
				action = MouseUp
			}
			break ScanLoop
		case 'M':
			if button == 3 {
				action = MouseMove
			} else {
				action = MouseDown
			}
			break ScanLoop
		}
	}

	return InputEvent{
		EventType:   MouseEvent,
		MouseButton: MouseButton(button), // FIXME translate button number
		MouseAction: action,
		MouseX:      x,
		MouseY:      y,
	}
}

func peekByte(s *bufio.Reader) byte {
	return peekBytes(s, 1)[0]
}

func peekBytes(s *bufio.Reader, num int) []byte {
	bytes, err := s.Peek(num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error peeking %v\n", err)
		return []byte{}
	}

	return bytes
}

func nextRune(s *bufio.Reader) rune {
	return rune(nextByte(s))
}

func nextByte(s *bufio.Reader) byte {
	b, err := s.ReadByte()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %v\n", err)
		return 0
	}

	return b
}

func isCtrlChar(b byte) bool {
	switch b {
	case 'D', 'E', 'H', 'M', 'N', 'O', 'P', 'V', 'W', 'X', 'Z', '[', '\\', ']', '^', '_':
		return true
	default:
		return false
	}
}
