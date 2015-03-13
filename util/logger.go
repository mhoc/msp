
package util

import (
  "fmt"
)

// You can enable additional log levels by setting these to true
// Turning too many to true might cause a lot of crazy output
//  LOG_TOKENS : the raw token names in the exact form they appear in the file
//  LOG_EXPANDED_TOKENS : additional token metadata like string values or identifier names
//  LOG_TRACE_1 : highly detailed trace logging when anything at all happens.
//                DO NOT combine this with the token logging, as this contains all that data plus more
//  LOG_TRACE_2 : less detailed trace logging.
const (
  LOG_TOKENS bool =           false
  LOG_EXPANDED_TOKENS bool =  false
  LOG_TRACE bool =            false
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
