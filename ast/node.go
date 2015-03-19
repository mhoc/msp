
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
  Print(prefix string)

}

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

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  Var *Variable
}

func (d Declaration) Execute() interface{} {
  symbol.Declare(d.Var.VariableName)
  return d.Var.VariableName
}

func (d Declaration) Print(p string) {
  fmt.Println(p + "Declaration")
  d.Var.Print(p + "| ")
}

// ====================
// Variable Definition:: var a = 1
// Definitions are essentially just typedefed assignments in this language,
// But the Execute() function is different
// ====================
type Definition struct {
  Decl *Declaration
  AssignNode *Assignment
}

func (d Definition) Execute() interface{} {
  // TODO Declare Variable
  // TODO Assign Variable
  return nil
}

func (d Definition) Print(p string) {
  fmt.Println(p + "Definition")
  d.Decl.Print(p + "| ")
  d.AssignNode.Print(p + "| ")
}

// ====================
// Equals, Assignment:: a    =  1
//                      LHS     RHS
// ====================
type Assignment struct {
  Lhs *Variable
  Rhs Node
}

func (a Assignment) Execute() interface{} {
  // Reset the value at lhs to be rhs
  // Return nothing
  return nil
}

func (a Assignment) Print(p string) {
  fmt.Println(p + "Assign")
  a.Lhs.Print(p + "| ")
  a.Rhs.Print(p + "| ")
}

// ====================
// Variable reference:: var something = myvar;
// ====================
type VarReference struct {
  Var *Variable
  Value interface{}
}

func (vr VarReference) Execute() interface{} {
  // TODO: GET Value of variable
  return vr.Value
}

func (vr VarReference) Print(p string) {
  switch vr.Value.(type) {
    case int:
      fmt.Println(p + vr.Var.VariableName + "=" + string(vr.Value.(int)))
      break
    case string:
      fmt.Println(p + vr.Var.VariableName + "=" + vr.Value.(string))
      break
  }
}
