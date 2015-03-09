
package log

import (
  "fmt"
)

type LogLevel int
const (
  LOG_SYNTAX LogLevel = iota
  LOG_TRACE LogLevel = iota
)

var DisplayLevel []LogLevel = []LogLevel{ LOG_TRACE }

func contains(s []LogLevel, e LogLevel) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func Syntax(s string) {
  if contains(DisplayLevel, LOG_SYNTAX) {
    fmt.Printf(s)
  }
}

func Trace(s string) {
  if contains(DisplayLevel, LOG_TRACE) {
    fmt.Println(s)
  }
}
