
// Maintains the symbol table

package symbol

import (
  "strings"
  "mhoc.co/msp/log"
)

var SymbolTable = make(map[string]*Type)

// Declares a new variable in the symbol table and sets it value to undefined
// If the variable has already been declared, its old value is deleted
func Declare(varn string) {
  log.Trace("tbl", "Declaring variable " + varn)

  // Make sure we aren't declaring an object key
  if strings.Contains(varn, ".") {
    log.Error{Msg: "Attempting to declare an object key"}.Report()
    return
  }

  // Put the variable in our symbol table
  SymbolTable[varn] = &Type{Undefined: true}

}

// Assigns a given value to a variable in the symbol table
// If the variable is undeclared, this throws an error and returns
// If the value passed in is not of a supported type, this throws an
// internal error and panics
func Assign(varn string, value interface{}, lineno int) {
  log.Tracef("tbl", "Assigning value %v to variable %s", value, varn)

  // Check to ensure the variable is declared
  if _, in := SymbolTable[varn]; !in {
    log.Error{Line: lineno, Type: log.UNDECLARED_VAR, Var: varn}.Report()
  }

  SymbolTable[varn] = &Type{Undefined: false, Value: value}

}

// Retrieves and returns the value stored in varn
// If the value does not exist, prints an error and returns an undefined type
func Get(varn string, lineno int) *Type {
  log.Tracef("tbl", "Getting the value for variable %s", varn)

  value, in := SymbolTable[varn]
  if !in {
    log.Error{Line: lineno, Type: log.VALUE, Var: varn}.Report()
    value = &Type{Undefined: true}
  }

  return value

}
