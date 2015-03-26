
// Different usages of variables
//  DECLARATION
//  DEFINITION
//  ASSIGNMENT
//  REFERENCE

package ast

import (
  "fmt"
  "mhoc.co/msp/symbol"
)

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  Var *Variable
}

func (d Declaration) Execute() interface{} {
  symbol.Declare(d.Var.VariableName)
  return nil
}

func (d Declaration) Print(p string) {
  fmt.Println(p + "Declare")
  d.Var.Print(p + "| ")
}

// ====================
// Variable Definition:: [var a = 1]
// Definitions are essentially just typedefed assignments in this language,
// But the Execute() function is different
// ====================
type Definition struct {
  Decl *Declaration
  Assign *Assignment
}

func (d Definition) Execute() interface{} {
  d.Decl.Execute()
  d.Assign.Execute()
  return nil
}

func (d Definition) Print(p string) {
  fmt.Println(p + "Define")
  d.Decl.Print(p + "| ")
  d.Assign.Print(p + "| ")
}

// ====================
// Equals, Assignment:: var [a  =  1]
//                          LHS   RHS
// ====================
type Assignment struct {
  Lhs *Variable
  Rhs Node
}

func (a Assignment) Execute() interface{} {
  symbol.Assign(a.Lhs.VariableName, a.Rhs.Execute())
  return nil
}

func (a Assignment) Print(p string) {
  fmt.Println(p + "Assign")
  a.Lhs.Print(p + "| ")
  a.Rhs.Print(p + "| ")
}

// ====================
// Variable reference:: var something = [myvar];
// ====================
type Reference struct {
  Var *Variable
  Value interface{}
  Undefined bool
}

func (vr Reference) Execute() interface{} {
  symbolType := symbol.Get(vr.Var.VariableName)
  vr.Undefined = symbolType.Undefined
  vr.Value = symbolType.Value
  return vr
}

func (vr Reference) Print(p string) {
  fmt.Printf(p + "Reference\n")
  vr.Var.Print(p + "| ")
  switch vr.Value.(type) {
    case int:
      fmt.Println(p + "| " + string(vr.Value.(int)))
      break
    case string:
      fmt.Println(p + "| " + vr.Value.(string))
      break
  }

}
