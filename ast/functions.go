// Contains anything related to functions
//  FUNCTIONCALL

package ast

import (
	"mhoc.co/msp/log"
)

// ====================
// Function Definitions
// ====================
type FunctionDef struct {
	Name string
	ArbitraryArgs bool
	ArgNames []string
	ExecMiniscript bool
	MSBody StatementList
	GoBody func(FunctionCall) interface{}
	Line int
}

func (fd FunctionDef) Execute() interface{} {
	log.Trace("ast", "Defining function " + fd.Name)

	// Store this function in the symbol table
	value := Value{Type: VALUE_FUNCTION, Value: fd, Line: fd.Line, Written: true}
	Declare(fd.Name)
	AssignToVariable(fd.Name, value, fd.Line)
	return nil
}

func (fd FunctionDef) LineNo() int {
	return fd.Line
}

// ====================
// Any function call
// Right now we only recognize one functionc all: document.write
// ====================
type FunctionCall struct {
	Name string
	Args []Statement
	LocalScope map[string]Value
	ReturnVal Value
	Line int
}

func (f FunctionCall) Execute() interface{} {
	log.Trace("ast", "Executing function " + f.Name)

	// Create the local stack and return val
	f.LocalScope = make(map[string]Value)
	f.ReturnVal = Value{Type: VALUE_UNDEFINED, Line: f.Line}
	// Retrieve the function definition
	funVal := GetVariable(f.Name, f.Line)
	if funVal.Type != VALUE_FUNCTION {
		log.TypeViolation(f.Line)
		return f.ReturnVal
	}
	funDef := funVal.Value.(FunctionDef)
	// Check number of arguments
	if !funDef.ArbitraryArgs && len(f.Args) != len(funDef.ArgNames) {
		log.TypeViolation(f.Line)
		return f.ReturnVal
	}
	if funDef.ExecMiniscript {
		// Load the arguments into the local stack
		for i, name := range funDef.ArgNames {
			f.LocalScope[name] = f.Args[i].Execute().(Value)
		}
		// Save the old scope and set this local as its own scope
		oldScope := Scope
		Scope = f.LocalScope
		// Execute the function
		retVal := funDef.MSBody.Execute()
		var nRetVal Value
		switch retVal.(type) {
			case Return:
				nRetVal = retVal.(Return).Value.Execute().(Value)
			case Value:
				nRetVal = retVal.(Value)
		}
		// Restore the old scope
		Scope = oldScope
		return nRetVal
	} else {
		funDef.GoBody(f)
		return nil
	}
}

func (f FunctionCall) LineNo() int {
	return f.Line
}

// ========================
// Return statement
// ========================
type Return struct {
	Line int
	Value Node
}

func (r Return) Execute() interface{} {
	// Just directly return the return statement
	return r
}

func (r Return) LineNo() int {
	return r.Line
}
