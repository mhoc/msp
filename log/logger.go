
package log

import (
  "fmt"
)

const (
  LOG_SYNTAX bool = true
  LOG_TRACE bool = false
)

func Syntax(s string) {
  if LOG_SYNTAX {
    fmt.Printf(s)
  }
}

func Trace(s string) {
  if LOG_TRACE {
    fmt.Printf(s)
  }
}

func Traceln(s string) {
  Trace(s + "\n")
}
