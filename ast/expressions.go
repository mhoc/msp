
// Contains entities which help form expressions
//  ADD
//  SUBTRACT
//  MULTIPLY
// DIVIDE

package ast

import (
  "fmt"
)

// ====================
// Additive Expression
// ====================
type Add struct {
  Lhs Node
  Rhs Node
}

func (a Add) Execute() interface{} {
  return nil
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
}

func (s Subtract) Execute() interface{} {
  return nil
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
}

func (m Multiply) Execute() interface{} {
  return nil
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
}

func (d Divide) Execute() interface{} {
  return nil
}

func (d Divide) Print(pre string) {
  fmt.Println(pre + "Divide")
  d.Lhs.Print(pre + "| ")
  d.Rhs.Print(pre + "| ")
}
