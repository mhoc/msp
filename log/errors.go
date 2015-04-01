
// Handles any error which the compiler might throw during parsing.

// There are various "levels" of errors which we recognize
//
//  INTERNAL
//  Anytime during execution we reach a state which should not happen
//  We always panic when one of these is thrown
//
//  REPORTED
//  Any error which the lab parameters require, in the format
//  which is required.
//  There is an additional requirement that reported errors only be reported
//  once per statement. Because of this, this file maintains the field
//  log.StatementNumber which should be incremented after every statement
//  is finished processing.
//  All reported output goes to stderr
//
//  EXPANDED
//  These are reported errors with expanded debugging output
//  This is more of just a "i can do this" thing instead of anything useful,
//  as it will be turned off upon turnin, but its still interesting.

package log

import (
  "fmt"
  "os"
)

type ErrorType int

type Error struct {
  Type ErrorType
  Var string    // Some error types also require a variable to be reported
  Line int      // The line number the error occured on
}

const (
  GENERIC ErrorType = iota
  INTERNAL ErrorType = iota
  TYPE_VIOLATION ErrorType = iota
  UNDECLARED_VAR ErrorType = iota
  VALUE ErrorType = iota
  CONDITION ErrorType = iota
)

// The line we are currently lexing
// Nex offers this functionality in a struct but it isn't exported
// from the package, and we need it here in errors
// Note that if this member is accessed during the process of node.Execute(),
// it will always read the line number of the last line of the file (for obvious
// reasons).
var LineNo int = 1

// The current statement we are executing
// This is incremented by the StatementLine's Execute() function call
// It is used in error reporting, as we only report one error per line
var Stmt int = 0
var lastLogged int = -1

func (er Error) Report() {
  switch er.Type {
    case GENERIC:
      fmt.Fprintf(os.Stderr, "[%d] Generic Error\n", er.Line)
      break
    case INTERNAL:
      panic("Internal compiler error reported")
    case TYPE_VIOLATION:
      er.typeViolation()
      break
    case UNDECLARED_VAR:
      er.undeclaredVar()
      break
    case VALUE:
      er.value()
      break
    case CONDITION:
      er.condition()
    default:
      fmt.Fprintf(os.Stderr, "[%d] Error\n", er.Line)
  }
  lastLogged = Stmt
}

func (er Error) typeViolation() {
  if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "Line %d, type violation\n", er.Line)
  }
}

func (er Error) undeclaredVar() {
  if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "Line %d, %s undeclared\n", er.Line, er.Var)
  }
}

func (er Error) value() {
  if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "Line %d, %s has no value\n", er.Line, er.Var)
  }
}

func (er Error) condition() {
  if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "Line %d, condition unknown\n", er.Line)
  }
}
