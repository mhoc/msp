
// Contains entities which help form expressions
//  ADD
//  SUBTRACT
//  MULTIPLY
// DIVIDE

package ast

import (
  "fmt"
  "mhoc.co/msp/log"
  "strings"
)

// =========================
// General Binary Expression
// This handles a lot of the error checking associated with undefined values
// in one location
// =========================
type BinaryExpression struct {
  Lhs Node
  Rhs Node
  Op string
  Line int
}

func (be BinaryExpression) Execute() interface{} {
  log.Tracef("ast", "Executing binary expression %s", be.Op)

  // Execute both sides
  left := be.Lhs.Execute().(*Value)
  right := be.Rhs.Execute().(*Value)

  // If one side is undefined and unwritten, we report a type violation and return undefined
  if (left.Type == VALUE_UNDEFINED && !left.Written) || (right.Type == VALUE_UNDEFINED && !right.Written) {
    log.Error{Line:be.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
    left.Type = VALUE_UNDEFINED
    return left
  }

  // If one side is undefined and written, we just return undefined
  if left.Type == VALUE_UNDEFINED || right.Type == VALUE_UNDEFINED {
    left.Type = VALUE_UNDEFINED
    return left
  }

  // If the types are simply not the same and none are undefined, we report a type violation and return undefined
  if left.Type != right.Type {
    log.Error{Line:be.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
    left.Type = VALUE_UNDEFINED
    return left
  }

  // Handle each operation separately
  switch (be.Op) {
    case "+":
      return handlePlus(left, right, be.Line)
    case "-":
      return handleMinus(left, right, be.Line)
    case "*":
      return handleMult(left, right, be.Line)
    case "/":
      return handleDivide(left, right, be.Line)
  }

  // Just return left to save an allocation
  return left

}

func (be BinaryExpression) LineNo() int {
  return be.Line
}

func (be BinaryExpression) Print(p string) {
  fmt.Printf(p + "%s\n", be.Op)
  be.Lhs.Print(p + "| ")
  be.Rhs.Print(p + "| ")
}

// Some functions to clean up the binary expression code
// in handling execution of different operators

func handlePlus(left *Value, right *Value, line int) *Value {

  // Integers
  if left.Type == VALUE_INT && right.Type == VALUE_INT {
    left.Value = left.Value.(int) + right.Value.(int)
    return left
  }

  // Strings
  if left.Type == VALUE_STRING && right.Type == VALUE_STRING {
    lStr := left.Value.(string)
    rStr := right.Value.(string)
    if strings.Contains(lStr, "<br />") || strings.Contains(rStr, "<br />") {
      log.Error{Line:line, Type: log.TYPE_VIOLATION, Msg: "Attepting to concat a string which contains a linebreak"}.Report()
      left.Type = VALUE_UNDEFINED
      return left
    }
    left.Value = left.Value.(string) + right.Value.(string)
    return left
  }

  // Otherwise, undefined. This should never be reached, though
  log.Error{Line:line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  left.Type = VALUE_UNDEFINED
  return left

}

func handleMinus(left *Value, right *Value, line int) *Value {

  // Integers
  if left.Type == VALUE_INT && right.Type == VALUE_INT {
    left.Value = left.Value.(int) - right.Value.(int)
    return left
  }

  // Otherwise, undefined. This should never be reached, though
  log.Error{Line:line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  left.Type = VALUE_UNDEFINED
  return left

}

func handleMult(left *Value, right *Value, line int) *Value {

  // Integers
  if left.Type == VALUE_INT && right.Type == VALUE_INT {
    left.Value = left.Value.(int) * right.Value.(int)
    return left
  }

  // Otherwise, undefined. This should never be reached, though
  log.Error{Line:line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  left.Type = VALUE_UNDEFINED
  return left

}

func handleDivide(left *Value, right *Value, line int) *Value {

  // Integers
  if left.Type == VALUE_INT && right.Type == VALUE_INT {
    left.Value = left.Value.(int) / right.Value.(int)
    return left
  }

  // Otherwise, undefined. This should never be reached, though
  log.Error{Line:line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  left.Type = VALUE_UNDEFINED
  return left

}
