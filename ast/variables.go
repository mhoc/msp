
// Different usages of variables
//  DECLARATION
//  DEFINITION
//  ASSIGNMENT
//  REFERENCE

package ast

import (
  "fmt"
  //"reflect"
  "mhoc.co/msp/symbol"
)

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  Var *Variable
  Line int
}

func (d Declaration) Execute() interface{} {
  symbol.Declare(d.Var.VariableName)
  return nil
}

func (d Declaration) LineNo() int {
  return d.Line
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
  Line int
}

func (d Definition) Execute() interface{} {
  d.Decl.Execute()
  d.Assign.Execute()
  return nil
}

func (d Definition) LineNo() int {
  return d.Line
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
  Line int
}

func (a Assignment) Execute() interface{} {
  rhsResult := a.Rhs.Execute()
/*
  typestr := reflect.TypeOf(rhsResult).Name()
  if (typestr == "Reference") {
    if (rhsResult.(Reference).Undefined) {
      rhsResult = &symbol.Type{Undefined: true}
    } else {
      rhsResult = symbol.Get(rhsResult.(Reference).Var.VariableName, a.LineNo())
    }
  }
*/
  symbol.Assign(a.Lhs.VariableName, rhsResult, a.LineNo())
  return nil
}

func (a Assignment) LineNo() int {
  return a.Line
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
  Line int
}

func (vr Reference) Execute() interface{} {
  symbolType := symbol.Get(vr.Var.VariableName, vr.LineNo())
  vr.Undefined = symbolType.Undefined
  vr.Value = symbolType.Value
  return vr
}

func (vr Reference) LineNo() int {
  return vr.Line
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
