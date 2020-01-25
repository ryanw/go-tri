package logger

import (
	"fmt"
	"os"
)

func Log(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
