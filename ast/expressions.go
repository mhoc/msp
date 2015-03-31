
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

// ====================
// Additive Expression
// ====================
type Add struct {
  Lhs Node
  Rhs Node
  Line int
}

func (a Add) Execute() interface{} {
  left := a.Lhs.Execute()
  right := a.Rhs.Execute()
  leftVal := left.(*Value)
  rightVal := right.(*Value)

  if leftVal.Type == VALUE_INT && rightVal.Type == VALUE_INT {
    leftVal.Value = leftVal.Value.(int) + rightVal.Value.(int)
    return leftVal
  }

  if leftVal.Type == VALUE_STRING && rightVal.Type == VALUE_STRING {
    lStr := leftVal.Value.(string)
    rStr := rightVal.Value.(string)
    if strings.Contains(lStr, "<br />") || strings.Contains(rStr, "<br />") {
      log.Error{Line:a.Line, Type: log.TYPE_VIOLATION, Msg: "Attepting to concat a string which contains a linebreak"}.Report()
      leftVal.Type = VALUE_UNDEFINED
      return leftVal
    }
    leftVal.Value = leftVal.Value.(string) + rightVal.Value.(string)
    return leftVal
  }

  log.Error{Line:a.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  leftVal.Type = VALUE_UNDEFINED
  return leftVal

}

func (a Add) LineNo() int {
  return a.Line
}

func (a Add) Print(pre string) {
  fmt.Println(pre + "Add")
  a.Lhs.Print(pre + "| ")
  a.Rhs.Print(pre + "| ")
}

// ====================
// Subtractive Expression
// ====================
type Subtract struct {
  Lhs Node
  Rhs Node
  Line int
}

func (s Subtract) Execute() interface{} {
  left := s.Lhs.Execute()
  right := s.Rhs.Execute()
  leftVal := left.(*Value)
  rightVal := right.(*Value)

  if leftVal.Type == VALUE_INT && rightVal.Type == VALUE_INT {
    leftVal.Value = leftVal.Value.(int) - rightVal.Value.(int)
    return leftVal
  }

  log.Error{Line:s.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  leftVal.Type = VALUE_UNDEFINED
  return leftVal

}

func (s Subtract) LineNo() int {
  return s.Line
}

func (s Subtract) Print(pre string) {
  fmt.Println(pre + "Subtract")
  s.Lhs.Print(pre + "| ")
  s.Rhs.Print(pre + "| ")
}

// ====================
// Multiplicative Expression
// ====================
type Multiply struct {
  Lhs Node
  Rhs Node
  Line int
}

func (m Multiply) Execute() interface{} {
  left := m.Lhs.Execute()
  right := m.Rhs.Execute()
  leftVal := left.(*Value)
  rightVal := right.(*Value)

  if leftVal.Type == VALUE_INT && rightVal.Type == VALUE_INT {
    leftVal.Value = leftVal.Value.(int) * rightVal.Value.(int)
    return leftVal
  }

  log.Error{Line:m.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  leftVal.Type = VALUE_UNDEFINED
  return leftVal
}

func (m Multiply) LineNo() int {
  return m.Line
}

func (m Multiply) Print(pre string) {
  fmt.Println(pre + "Multiply")
  m.Lhs.Print(pre + "| ")
  m.Rhs.Print(pre + "| ")
}

// ====================
// Divide Expression
// ====================
type Divide struct {
  Lhs Node
  Rhs Node
  Line int
}

func (d Divide) Execute() interface{} {
  left := d.Lhs.Execute()
  right := d.Rhs.Execute()
  leftVal := left.(*Value)
  rightVal := right.(*Value)

  if leftVal.Type == VALUE_INT && rightVal.Type == VALUE_INT {
    leftVal.Value = leftVal.Value.(int) / rightVal.Value.(int)
    return leftVal
  }

  log.Error{Line:d.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to add types which are not supported"}.Report()
  leftVal.Type = VALUE_UNDEFINED
  return leftVal
}

func (d Divide) LineNo() int {
  return d.Line
}

func (d Divide) Print(pre string) {
  fmt.Println(pre + "Divide")
  d.Lhs.Print(pre + "| ")
  d.Rhs.Print(pre + "| ")
}
