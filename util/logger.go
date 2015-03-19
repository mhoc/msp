
package util

import (
  "fmt"
)

// You can enable additional log levels by setting these to true
// Turning too many to true might cause a lot of crazy output
//  LOG_TOKENS : the raw token names in the exact form they appear in the file
//  LOG_EXPANDED_TOKENS : additional token metadata like string values or identifier names
//  LOG_TRACE : logs the trace execution of methods through the program. I try to be complete as possible here.
//  LOG_AST : logs the AST of the program after it is entierly parsed
const (
  LOG_TOKENS bool =           false
  LOG_EXPANDED_TOKENS bool =  false
  LOG_TRACE bool =            false
  LOG_AST bool =              true
)

func LogToken(s string) {
  if LOG_TOKENS {
    fmt.Printf(s)
  }
}

func LogTokenData(s string) {
  if LOG_EXPANDED_TOKENS {
    fmt.Printf(s)
  }
}

func LogTrace(s string) {
  if LOG_TRACE {
    fmt.Println(s)
  }
}
