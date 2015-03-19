
// Anything to do with single or groups of statements
// STATEMENT LIST

package ast

import "fmt"

// ====================
// A StatementList is any ordered collection of independent statements
// This is the type of the root Node, but also the type of things like if bodies
// ====================
type StatementList struct {
  List []Node
}

func (s StatementList) Execute() interface{} {
  for _, child := range s.List {
    child.Execute()
  }
  return nil
}

func (s StatementList) Print(p string) {
  for _, child := range s.List {
    fmt.Println("Statement")
    child.Print("| ")
  }
}
