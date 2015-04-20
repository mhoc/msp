// Contains anything related to functions
//  FUNCTIONCALL

package ast

import (
	"fmt"
	"mhoc.co/msp/log"
)

// ====================
// Any function call
// Right now we only recognize one functionc all: document.write
// ====================
type FunctionCall struct {
	Name string
	Args []Statement
	Line int
}

func (f FunctionCall) Execute() interface{} {
	// For now, assume all function calls are document.write
	// this will be improved later with a function lookup table
	if f.Name != "document.write" {
		panic("Error: Attempting to call function that is not document.write")
	}
	for _, arg := range f.Args {
		argv := arg.Execute().(*Value)
		fmt.Print(argv.ToString())
		switch argv.Type {
			case VALUE_OBJECT:
				if !log.EXTENSIONS {
					log.TypeViolation(f.LineNo())
				}
			case VALUE_ARRAY:
				if !log.EXTENSIONS {
					log.TypeViolation(f.LineNo())
				}
		}
	}
	return nil
}

func (f FunctionCall) LineNo() int {
	return f.Line
}
