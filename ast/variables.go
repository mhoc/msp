
// Different usages of variables
//  DECLARATION
//  DEFINITION
//  ASSIGNMENT
//  REFERENCE

package ast

import (
  "fmt"
)

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  Name string
  Line int
}

func (d Declaration) Execute() interface{} {
  SymDeclare(d.Name)
  return nil
}

func (d Declaration) LineNo() int {
  return d.Line
}

func (d Declaration) Print(p string) {
  fmt.Println(p + "Declare")
  fmt.Printf(p + "| %s\n", d.Name)
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
  Name string
  Rhs Node
  Line int
}

func (a Assignment) Execute() interface{} {
  rhsResult := a.Rhs.Execute()

  // The type of the right side should always be a Value
  // This line is included just to throw an error if it ever isn't, which is
  // mainly for debugging
  rightValue := rhsResult.(Value)

  SymAssign(a.Name, rightValue)
  return nil
}

func (a Assignment) LineNo() int {
  return a.Line
}

func (a Assignment) Print(p string) {
  fmt.Println(p + "Assign")
  fmt.Printf(p + "| %s\n", a.Name)
  a.Rhs.Print(p + "| ")
}

// ====================
// Variable reference:: var something = [myvar];
// ====================
type Reference struct {
  Name string
  Value *Value
  Line int
}

func (vr Reference) Execute() interface{} {
  value := SymGet(vr.Name, vr.LineNo())
  vr.Value = value
  return vr
}

func (vr Reference) LineNo() int {
  return vr.Line
}

func (vr Reference) Print(p string) {
  fmt.Printf(p + "Reference\n")
  vr.Value.Print(p + "| ")
}
