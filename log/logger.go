
// Handles all of the informational logging functionality of the interpreter

package log

import (
  "fmt"
  "mhoc.co/msp/ast"
)

// You can enable additional log levels by setting these to true
// Turning too many to true might cause a lot of crazy output
//  LOG_TOKENS : the raw token names in the exact form they appear in the file
//  LOG_TRACE : logs the trace execution of methods through the program. I try to be complete as possible here.
//  LOG_AST : logs the entire ast of any node which is passed into log.Ast
// These are overwritten by cmd line flags at runtime if they are set.

var LOG_TOKENS bool
var LOG_TRACE bool
var LOG_AST bool

func Token(s string) {
  if LOG_TOKENS {
    fmt.Printf(s)
  }
}

func Trace(s string) {
  if LOG_TRACE {
    fmt.Println(s)
  }
}

func Ast(n ast.Node) {
  if LOG_AST {
    n.Print("")
  }
}
