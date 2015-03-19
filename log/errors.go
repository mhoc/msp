
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
  Line int
  Var string    // Some error types also require a variable to be reported
  Msg string    // Expanded error reporting message directly copied to stderr
}

const (
  LOG_EXPANDED_ERRORS bool = true

  INTERNAL ErrorType = iota
  TYPE_VIOLATION ErrorType = iota
  UNDECLARED_VAR ErrorType = iota
  VALUE ErrorType = iota
  CONDITION ErrorType = iota
)

var Stmt int = 0
var lastLogged int = -1

func (er Error) Report() {
  switch er.Type {
    case INTERNAL:
      panic(er.Msg)
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
      panic("Attempting to report an error type which log.ErrorType does not recognize")
  }
  lastLogged = Stmt
}

func (er Error) typeViolation() {
  if LOG_EXPANDED_ERRORS {
    fmt.Fprintf(os.Stderr, "[%d] Type Violation\n", er.Line)
    fmt.Fprintf(os.Stderr, "|-> %s\n", er.Msg)
  } else if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "line %d, type violation\n", er.Line)
  }
}

func (er Error) undeclaredVar() {
  if LOG_EXPANDED_ERRORS {
    fmt.Fprintf(os.Stderr, "[%d] Attempting to assign to undeclared variable %s\n", er.Line, er.Var)
  } else if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "line %d, %s undeclared\n", er.Line, er.Var)
  }
}

func (er Error) value() {
  if LOG_EXPANDED_ERRORS {
    fmt.Fprintf(os.Stderr, "[%d] Attempting to access the value of variable %s, which has no value\n", er.Line, er.Var)
  } else if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "line %d, %s has no value\n", er.Line, er.Var)
  }
}

func (er Error) condition() {
  if LOG_EXPANDED_ERRORS {
    fmt.Fprintf(os.Stderr, "[%d] Condition in branch could not be evaluated to true or false\n", er.Line)
  } else if lastLogged != Stmt {
    fmt.Fprintf(os.Stderr, "line %d, condition unknown\n", er.Line)
  }
}
