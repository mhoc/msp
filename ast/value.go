
// Primitive types in the language
// aka anything which the lexer directly returns with no semantic parsing knowledge
// VALUE

package ast

import (
  "fmt"
)

// ===================
// Value
// Represents a binding between a go primitive type and a ms
// primitive type. Also represents undefined types
// Values have types which are defined by an enum
// ===================
type ValueType int
const (
  VALUE_UNDEFINED ValueType = iota  // Disregard value of Value
  VALUE_INT ValueType = iota        // type(Value) == int
  VALUE_STRING ValueType = iota     // type(Value) == string
  VALUE_OBJECT ValueType = iota     // type(Value) == map[string]*Value
)
type Value struct {
  Type ValueType
  Value interface{}
  Line int
  Written bool
}

func (v Value) Execute() interface{} {
  // We just return the value itself, not the containing interface
  // because we need information about its type in parent ast nodes
  return &v
}

func (v Value) LineNo() int {
  return v.Line
}

func (v Value) Print(p string) {
  switch v.Type {
    case VALUE_UNDEFINED:
      fmt.Println(p + "undefined")
      break
    case VALUE_INT:
      fmt.Printf(p + "[int] %v\n", v.Value)
      break
    case VALUE_STRING:
      fmt.Printf(p + "[str] %v\n", v.Value)
      break
    case VALUE_OBJECT:
      fmt.Println(p + "Object Value")
      for key, value := range v.Value.(map[string]*Value) {
        fmt.Printf(p + "| %s\n", key)
        value.Print(p + "| | ")
      }
      break
  }
}
