// While not technically part of the AST, I have this
// in the AST package because it accesses type definitions in the AST,
// and Go does not like circular package references

package ast

import (
	"fmt"
	"strconv"
	"mhoc.co/msp/log"
)

var SymbolTable = make(map[string]*Value)

func SymDeclare(name string) {
	log.Trace("tbl", "Declaring variable "+name)

	// Put the variable in our symbol table
	SymbolTable[name] = &Value{Type: VALUE_UNDEFINED, Written: false}
}

func SymAssignVar(name string, value Value, lineno int) {
	log.Tracef("tbl", "Assigning value %v to variable %s", value.ToString(), name)

	// Check to ensure the variable is declared
	if _, in := SymbolTable[name]; !in {
		log.UndeclaredVariable(lineno, name)
	}

	SymbolTable[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
}

func SymAssignObj(name string, child string, value Value, lineno int) {
	log.Tracef("tbl", "Assigning value %v to key %v on object %v", value.ToString(), child, name)

	if _, in := SymbolTable[name]; !in {
		log.UndeclaredVariable(lineno, name)
	}

	obj := SymGetVar(name, lineno)
	if obj.Type != VALUE_OBJECT {
		log.TypeViolation(lineno)
		return
	}

	if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
		log.TypeViolation(lineno)
		return
	}

	obj.Value.(map[string]*Value)[child] = &value
}

func SymAssignArr(name string, index int, value Value, lineno int) {
	log.Tracef("tbl", "Assigning value %v to %v[%v]", value.ToString(), name, index)

	// Check if the array is in the table
	if _, in := SymbolTable[name]; !in {
		log.UndeclaredVariable(lineno, name)
	}

	arr := SymGetVar(name, lineno)
	if arr.Type != VALUE_ARRAY {
		log.TypeViolation(lineno)
		return
	}

	if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
		log.TypeViolation(lineno)
		return
	}

	arr.Value.(map[string]*Value)[strconv.Itoa(index)] = &value
}

func SymGetVar(name string, lineno int) Value {
	log.Tracef("tbl", "Getting the value for variable %s", name)

	value, in := SymbolTable[name]
	if !in {
		log.ValueError(lineno, name)
		value = &Value{Written: false, Type: VALUE_UNDEFINED, Line: lineno}
		return *value
	}

	if value.Type == VALUE_UNDEFINED && !value.Written {
		log.ValueError(lineno, name)
	}

	log.Tracef("tbl", "Value was: %v", value.ToString())

	// Make a copy of the value. Trust me, this bug took YEARS to find.
	nVal := Value{Type: value.Type, Value: value.Value, Written: value.Written}
	return nVal
}

func SymGetObj(parent string, child string, lineno int) Value {
	log.Tracef("tbl", "Getting the value for child %s in object %s", child, parent)

	value, in := SymbolTable[parent]
	if !in {
		log.ValueError(lineno, parent)
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return *value
	}

	if value.Type == VALUE_UNDEFINED {
		log.ValueError(lineno, parent)
	}

	if value.Type != VALUE_OBJECT {
		log.TypeViolation(lineno)
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return *value
	}

	value, in = value.Value.(map[string]*Value)[child]
	if !in {
		log.ValueError(lineno, parent + "." + child)
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return value
	}

	log.Tracef("tbl", "Value was: %v", value.ToString())
	return *value

}

func SymGetArr(parent string, index int, lineno int) *Value {
	log.Tracef("tbl", "Getting the value for array member %s[%d]", parent, index)

	value, in := SymbolTable[parent]
	if !in {
		log.ValueError(lineno, parent)
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return value
	}

	if value.Type == VALUE_UNDEFINED {
		log.ValueError(lineno, parent)
	}

	if value.Type != VALUE_ARRAY {
		log.TypeViolation(lineno)
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return value
	}

	value, in = value.Value.(map[string]*Value)[strconv.Itoa(index)]
	if !in {
		log.ValueError(lineno, fmt.Sprintf("%s[%d]", parent, index))
		value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
		return value
	}

	log.Tracef("tbl", "Value was: %v", value.ToString())
	return *value

}
