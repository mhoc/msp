
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
}

func (i Integer) Execute() interface{} {
  return i.Value
}

func (i Integer) Print(p string) {
  fmt.Printf("%s[integer] %d\n", p, i.Value)
}

// ====================
// String. Any static string inside the code.
// ====================
type String struct {
  Value string
}

func (s String) Execute() interface{} {
  return s.Value
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
}

func (v Variable) Execute() interface{} {
  return v.VariableName
}

func (v Variable) Print(p string) {
  fmt.Println(p + v.VariableName)
}
