
package ast

import (
  "os"
  "mhoc.co/msp/log"
)

type Assert struct {
  Line int
  Value Node
}

func (a Assert) Execute() interface{} {
  // Execute the value
  result := a.Value.Execute()
  // If its not a value, type error
  switch result.(type) {
    case Value:
    default:
      log.TypeViolation(a.Line)
  }
  // Evalute the value to a boolean
  bVal := result.(Value).ToBoolean()
  // Exit the program if it evaluates to false
  if !bVal.Value.(bool) {
    os.Exit(1)
  }
  return nil
}

func (a Assert) LineNo() int {
  return a.Line
}
