
// Primitive types in the language
// aka anything which the lexer directly returns with no semantic parsing knowledge
// INTEGER
// STRING
// VARIABLE

package ast

import (
  "fmt"
)

// ====================
// Integer. Any static integer inside the code.
// ====================
type Integer struct {
  Value int
  Line int
}

func (i Integer) Execute() interface{} {
  return i.Value
}

func (i Integer) LineNo() int {
  return i.Line
}

func (i Integer) Print(p string) {
  fmt.Printf("%s[integer] %d\n", p, i.Value)
}

// ====================
// String. Any static string inside the code.
// ====================
type String struct {
  Value string
  Line int
}

func (s String) Execute() interface{} {
  return s.Value
}

func (s String) LineNo() int {
  return s.Line
}

func (s String) Print(p string) {
  fmt.Println(p + "[string] " + s.Value)
}

// ====================
// Variable Usage
// ANY usage of a variable somehow derives this struct
// ====================
type Variable struct {
  VariableName string
  Line int
}

func (v Variable) Execute() interface{} {
  return v.VariableName
}

func (v Variable) LineNo() int {
  return v.Line
}

func (v Variable) Print(p string) {
  fmt.Println(p + v.VariableName)
}
