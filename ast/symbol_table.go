
// While not technically part of the AST, I have this
// in the AST package because it accesses type definitions in the AST,
// and Go does not like circular package references

package ast

import (
  "fmt"
  "mhoc.co/msp/log"
)

var SymbolTable = make(map[string]*Value)

func SymDeclare(name string) {
  log.Trace("tbl", "Declaring variable " + name)

  // Put the variable in our symbol table
  SymbolTable[name] = &Value{Type: VALUE_UNDEFINED, Written: false}
}

func SymAssignVar(name string, value *Value) {
  log.Tracef("tbl", "Assigning value %v to variable %s", value.Value, name)

  // Check to ensure the variable is declared
  if _, in := SymbolTable[name]; !in {
    log.Error{Line: value.LineNo(), Type: log.UNDECLARED_VAR, Var: name}.Report()
  }

  SymbolTable[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
}

func SymAssignObj(name string, child string, value *Value) {
  log.Tracef("tbl", "Assigning value %v to key %v on object %v", value.Value, child, name)

  obj := SymGetVar(name, value.LineNo())
  if (obj.Type != VALUE_OBJECT && obj.Type != VALUE_ARRAY) {
    log.Error{Line: value.LineNo(), Type: log.TYPE_VIOLATION}.Report()
    return
  }
  obj.Value.(map[string]*Value)[child] = value
}

func SymAssignArr(name string, index int, value *Value) {
  log.Tracef("tbl", "Assigning value %v to %v[%v]", value.Value, name, index)
  SymAssignObj(name, string(index), value)
}

func SymGetVar(name string, lineno int) *Value {
  log.Tracef("tbl", "Getting the value for variable %s", name)

  value, in := SymbolTable[name]
  if !in {
    log.Error{Line: lineno, Type: log.TYPE_VIOLATION}.Report()
    value = &Value{Written: false, Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  if value.Type == VALUE_UNDEFINED && !value.Written {
    log.Error{Line: lineno, Type: log.VALUE, Var: name}.Report()
  }

  log.Tracef("tbl", "Value was: %v", value.Value)

  // Make a copy of the value. Trust me, this bug took YEARS to find.
  nVal := &Value{Type: value.Type, Value: value.Value, Written: value.Written}
  return nVal
}

func SymGetObj(parent string, child string, lineno int) *Value {
  log.Tracef("tbl", "Getting the value for child %s in object %s", child, parent)

  value, in := SymbolTable[parent]
  if !in {
    log.Error{Line: lineno, Type: log.TYPE_VIOLATION}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  if value.Type != VALUE_OBJECT {
    log.Error{Line: lineno, Type: log.TYPE_VIOLATION}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  value, in = value.Value.(map[string]*Value)[child]
  if !in {
    log.Error{Line: lineno, Type: log.VALUE, Var: parent + "." + child}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  return value

}

func SymGetArr(parent string, index int, lineno int) *Value {
  log.Tracef("tbl", "Getting the value for array member %s[%d]", parent, index)

  value, in := SymbolTable[parent]
  if !in {
    log.Error{Line: lineno, Type: log.TYPE_VIOLATION}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  if value.Type != VALUE_ARRAY {
    log.Error{Line: lineno, Type: log.TYPE_VIOLATION}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  value, in = value.Value.(map[string]*Value)[string(index)]
  if !in {
    log.Error{Line: lineno, Type: log.VALUE, Var: fmt.Sprintf("%s[%d]", parent, index)}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  return value

}
