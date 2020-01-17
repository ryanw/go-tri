package main

import (
  "fmt"
  "time"
  m "math"
  . "./terminal"
)

func main() {
  fmt.Printf("\x1b[34;43m Hello, world \x1b[0m\n")
  start := Position { 10, 3 }
  end := Position { 40, 20 }
  Clear()


  t := 1.0
  for {
    // Clear prev line
    PlotLine(start, end, ' ')

    // Draw new line
    start.X = int32(m.Abs(m.Sin(t * 0.1)) * 50)
    end.Y = int32(m.Abs(m.Cos(t * 0.1)) * 50)
    PlotLine(start, end, '•')

    time.Sleep(100 * time.Millisecond)


    t += 1.0
  }
}

