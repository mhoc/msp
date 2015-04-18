// Different usages of variables
//  DECLARATION
//  DEFINITION
//  ASSIGNMENT
//  REFERENCE

package ast

import (
	"mhoc.co/msp/log"
)

type VariableType int

const (
	VAR_NORM   VariableType = iota
	VAR_OBJECT VariableType = iota
	VAR_ARRAY  VariableType = iota
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

// ====================
// Variable Definition:: [var a = 1]
// Definitions are essentially just typedefed assignments in this language,
// But the Execute() function is different
// ====================
type Definition struct {
	Decl   *Declaration
	Assign *Assignment
	Line   int
}

func (d Definition) Execute() interface{} {

	// This is more complicated than it needs to be because of the fucking
	// var x = x corner case. Normally I'd just execute the decl and
	// execute the assign in that order, but NNOOOO apparently the order has
	// to be GET the right, make the decl, THEN do the assign
	d.Assign.Rhs.Execute()
	d.Decl.Execute()
	d.Assign.Execute()
	return nil
}

func (d Definition) LineNo() int {
	return d.Line
}

// ====================
// Equals, Assignment:: var [a  =  1]
//                          LHS   RHS
// ====================
type Assignment struct {
	Type     VariableType
	Name     string
	ObjChild string
	Index    Node
	Rhs      Node
	Line     int
}

func (a Assignment) Execute() interface{} {
	rhsResult := a.Rhs.Execute()

	// The type of the right side should always be a Value
	// This line is included just to throw an error if it ever isn't, which is
	// mainly for debugging
	rightValue := rhsResult.(*Value)

	switch a.Type {
	case VAR_NORM:
		SymAssignVar(a.Name, rightValue, a.Line)
	case VAR_OBJECT:
		SymAssignObj(a.Name, a.ObjChild, rightValue, a.Line)
	case VAR_ARRAY:
		// Check to ensure the index is an int: otherwise type error
		index := a.Index.Execute().(*Value)
		if index.Type != VALUE_INT {
			log.TypeViolation(a.Line)
			return nil
		}
		SymAssignArr(a.Name, index.Value.(int), rightValue, a.Line)
	}

	return nil
}

func (a Assignment) LineNo() int {
	return a.Line
}

// ====================
// Variable reference:: var something = [myvar];
// ====================
type Reference struct {
	Type     VariableType
	Name     string
	ObjChild string
	Index    Node
	Line     int
}

func (vr Reference) Execute() interface{} {
	switch vr.Type {
	case VAR_NORM:
		return SymGetVar(vr.Name, vr.LineNo())
	case VAR_OBJECT:
		return SymGetObj(vr.Name, vr.ObjChild, vr.LineNo())
	case VAR_ARRAY:
		// Check to ensure the index is an int: otherwise type error
		index := vr.Index.Execute().(*Value)
		if index.Type != VALUE_INT {
			log.TypeViolation(vr.Line)
			return &Value{Type: VALUE_UNDEFINED, Line: vr.Line}
		}
		return SymGetArr(vr.Name, index.Value.(int), vr.LineNo())
	default:
		panic("Bad variable reference type")
	}
	return nil
}

func (vr Reference) LineNo() int {
	return vr.Line
}
