
package main

import "fmt"

type LogLevel int
const (
  LOG_SYNTAX LogLevel = iota
  LOG_TRACE LogLevel = iota
)

var DisplayLevel []LogLevel = []LogLevel{ LOG_SYNTAX, LOG_TRACE }

func contains(s []LogLevel, e LogLevel) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func LogSyntax(s string) {
  if contains(DisplayLevel, LOG_SYNTAX) {
    fmt.Printf(s)
  }
}

func LogTrace(s string) {
  if contains(DisplayLevel, LOG_TRACE) {
    fmt.Printf(s)
  }
}
