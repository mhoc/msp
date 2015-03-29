
// Contains entities which help form expressions
//  ADD
//  SUBTRACT
//  MULTIPLY
// DIVIDE

package ast

import (
  "fmt"
  "mhoc.co/msp/log"
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
  return nil
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
  return nil
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
  return nil
}

func (d Divide) LineNo() int {
  return d.Line
}

func (d Divide) Print(pre string) {
  fmt.Println(pre + "Divide")
  d.Lhs.Print(pre + "| ")
  d.Rhs.Print(pre + "| ")
}