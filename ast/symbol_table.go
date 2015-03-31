
// While not technically part of the AST, I have this
// in the AST package because it accesses type definitions in the AST,
// and Go does not like circular package references

package ast

import (
  "strings"
  "mhoc.co/msp/log"
)

var SymbolTable = make(map[string]*Value)

func SymDeclare(name string) {
  log.Trace("tbl", "Declaring variable " + name)

  // If we are declaring an object key, do nothing.
  // Object keys do not need to be declared, but for the time being I am
  // assuming that if they are declared they can be done so quietly
  if strings.Contains(name, ".") {
    return
  }

  // Put the variable in our symbol table
  SymbolTable[name] = &Value{Type: VALUE_UNDEFINED}

}

func SymAssign(name string, value *Value) {
  log.Tracef("tbl", "Assigning value %v to variable %s", value.Value, name)

  // If its an object key, we need to get the object and add it to the object map
  if strings.Contains(name, ".") {
    sp := strings.Split(name, ".")
    obj := SymGet(sp[0], value.LineNo())
    if (obj.Type != VALUE_OBJECT) {
      log.Error{Line: value.LineNo(), Type: log.TYPE_VIOLATION}.Report()
      return
    }
    obj.Value.(map[string]*Value)[sp[1]] = value
    return
  }

  // Check to ensure the variable is declared
  if _, in := SymbolTable[name]; !in {
    log.Error{Line: value.LineNo(), Type: log.UNDECLARED_VAR, Var: name}.Report()
  }

  SymbolTable[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line}

}

func SymGet(name string, lineno int) *Value {
  log.Tracef("tbl", "Getting the value for variable %s", name)

  var value *Value

  if (strings.Contains(name, ".")) {
    sp := strings.Split(name, ".")
    parent := sp[0]
    child := sp[1]

    value, in := SymbolTable[parent]
    if !in {
      log.Error{Line: lineno, Type: log.VALUE, Var: name}.Report()
      value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
      return value
    }

    value, in = value.Value.(map[string]*Value)[child]
    if !in {
      log.Error{Line: lineno, Type: log.VALUE, Var: name}.Report()
      value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
      return value
    }

    return value
  }

  value, in := SymbolTable[name]
  if !in {
    log.Error{Line: lineno, Type: log.VALUE, Var: name}.Report()
    value = &Value{Type: VALUE_UNDEFINED, Line: lineno}
    return value
  }

  return value

}
