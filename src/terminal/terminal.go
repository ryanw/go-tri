package terminal

import (
  "os"
  "fmt"
  "bufio"
)

type Terminal struct {
  width, height int
  buffer bufio.Writer
}

func NewTerminal(width, height int) Terminal {
  return Terminal {
    width: width,
    height: height,
    buffer: *bufio.NewWriterSize(os.Stdout, 4096),
  }
}


func (self *Terminal) Resize(width, height int) {
  self.width = width
  self.height = height
}

func (self *Terminal) Draw(callback func()) {
  callback()
  self.buffer.Flush()
}

func (self *Terminal) Write(format string, a ...interface{}) {
  fmt.Fprintf(&self.buffer, format, a...)
}

