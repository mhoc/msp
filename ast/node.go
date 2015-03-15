
package ast

import (
  "fmt"
  "mhoc.co/msp/symbol"
)

// ====================
// Every Node in our AST decends from this Node interface
// ====================
type Node interface {

  // Execute is a function that "executes" the function of a node in the AST
  // This is the core of the compiler design. We build up an AST during lexing and semantic
  // analysis, then call Execute() on the root node of the ast, which calls its children's
  // execute function and so on
  // The leaf node types will have an empty or non-recursive execute function
  // Execute can provide an optional return value if the node being executed makes sense
  // to return something (say, a literal or variable reference)
  Execute() interface{}

  // We provide printing functionality for a visual representation of the AST contained
  // in ast/print.go. The process for this is very similar to Execute()
  Print()

}

// ====================
// A StatementList is any ordered collection of independent statements
// This is the type of the root Node, but also the type of things like if bodies
// ====================
type StatementList struct {
  List []Node
}

func (s *StatementList) Execute() interface{} {
  for _, child := range s.List {
    child.Execute()
  }
  return nil
}

func (s *StatementList) Print() {
  for _, child := range s.List {
    child.Print()
  }
}

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  VariableName string
}

func (d *Declaration) Execute() interface{} {
  symbol.Declare(d.VariableName)
  return d.VariableName
}

func (d *Declaration) Print() {
  fmt.Println(d.VariableName)
}
