
// Symbol table
// All declarations, definitions, and retrievals are done through global
// functions. These functions interact with the global Scope variable
// which defines where the functions should put the variables which
// are being defined.

package ast

import (
	"fmt"
	"strconv"
	"mhoc.co/msp/log"
)

// This is Scope. If we are in a function, that functions symbol table
// will be copied here. If we are at the global level, this will be nil
// and we access the global symbol table
var Scope map[string]*Value = nil

// The global symbol table
var GlobalScope map[string]*Value = make(map[string]*Value)

// =============================================================================
// DECLARATION
// =============================================================================

// Declares a new variable in the current scope
func Declare(name string) {
  if Scope == nil {
    GlobalScope[name] = &Value{Type: VALUE_UNDEFINED, Written: false}
  } else {
    Scope[name] = &Value{Type:VALUE_UNDEFINED, Written: false}
  }
}

// =============================================================================
// ASSIGNMENT TO VARIABLE
// =============================================================================

func AssignToVariable(name string, value *Value, lineno int) {
  if Scope == nil {
    AssignToGlobalVariable(name, value, lineno)
  } else {
    AssignToScopedVariable(name, value, lineno)
  }
}

func AssignToGlobalVariable(name string, value *Value, lineno int) {
  if _, in := GlobalScope[name]; !in {
    log.UndeclaredVariable(lineno, name)
  }
  GlobalScope[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
}

func AssignToScopedVariable(name string, value *Value, lineno int) {
  if _, in1 := Scope[name]; !in1 {
    if _, in2 := GlobalScope[name]; !in2 {
      log.UndeclaredVariable(lineno, name)
      Scope[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
    } else {
      GlobalScope[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
    }
  } else {
    Scope[name] = &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
  }
}

// =============================================================================
// ASSIGNMENT TO OBJECT KEY
// =============================================================================

func AssignToObjectKey(objectName string, keyName string, value *Value, lineno int) {
  if Scope == nil {
    AssignToGlobalObjectKey(objectName, keyName, value, lineno)
  } else {
    AssignToScopedObjectKey(objectName, keyName, value, lineno)
  }
}

func AssignToGlobalObjectKey(objectName string, keyName string, value *Value, lineno int) {
  if _, in := GlobalScope[objectName]; !in {
    log.UndeclaredVariable(lineno, objectName)
  }

  // Get the object from the symbol table
  object := GetVariable(objectName, lineno)
  if object.Type != VALUE_OBJECT {
    log.TypeViolation(lineno)
    return
  }

  // If we are assigning an object or array, fail
  if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
    log.TypeViolation(lineno)
    return
  }

  object.Value.(map[string]*Value)[keyName] =
		&Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
}

func AssignToScopedObjectKey(objectName string, keyName string, value *Value, lineno int) {
  if _, in1 := Scope[objectName]; !in1 {
    if _, in2 := GlobalScope[objectName]; !in2 {
      // Not in either local or global
      log.UndeclaredVariable(lineno, objectName)
      return
    } else {
      // In global
      gObject := GetVariable(objectName, lineno)
      if gObject.Type != VALUE_OBJECT {
        log.TypeViolation(lineno)
        return
      }
      if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
        log.TypeViolation(lineno)
        return
      }
      gObject.Value.(map[string]*Value)[keyName] =
        &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
    }
  } else {
    lObject := GetVariable(objectName, lineno)
    if lObject.Type != VALUE_OBJECT {
      log.TypeViolation(lineno)
      return
    }
    if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
      log.TypeViolation(lineno)
      return
    }
    lObject.Value.(map[string]*Value)[keyName] =
      &Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
  }
}

// =============================================================================
// ASSIGNMENT TO ARRAY INDEX
// =============================================================================

func AssignToArrayIndex(arrayName string, indexNo int, value *Value, lineno int) {
  if Scope == nil {
    AssignToGlobalArrayIndex(arrayName, indexNo, value, lineno)
  } else {
    AssignToScopedArrayIndex(arrayName, indexNo, value, lineno)
  }
}

func AssignToGlobalArrayIndex(arrayName string, indexNo int, value *Value, lineno int) {

  if _, in := GlobalScope[arrayName]; !in {
    log.UndeclaredVariable(lineno, arrayName)
  }

  array := GetVariable(arrayName, lineno)
  if array.Type != VALUE_ARRAY {
    log.TypeViolation(lineno)
    return
  }

  if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
    log.TypeViolation(lineno)
    return
  }

  array.Value.(map[string]*Value)[strconv.Itoa(indexNo)] =
		&Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
}

func AssignToScopedArrayIndex(arrayName string, indexNo int, value *Value, lineno int) {

  if _, in1 := Scope[arrayName]; !in1 {
    if _, in2 := GlobalScope[arrayName]; !in2 {
      log.UndeclaredVariable(lineno, arrayName)
      return
    } else {
      gArray := GetVariable(arrayName, lineno)
      if gArray.Type != VALUE_ARRAY {
        log.TypeViolation(lineno)
        return
      }
      if value.Type == VALUE_OBJECT || value.Type == VALUE_ARRAY {
        log.TypeViolation(lineno)
        return
      }
      gArray.Value.(map[string]*Value)[strconv.Itoa(indexNo)] =
    		&Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
    }
  } else {
    lArray := GetVariable(arrayName, lineno)
    if lArray.Type != VALUE_ARRAY {
      log.TypeViolation(lineno)
      return
    }
    if value.Type == VALUE_ARRAY || value.Type == VALUE_OBJECT {
      log.TypeViolation(lineno)
      return
    }
    lArray.Value.(map[string]*Value)[strconv.Itoa(indexNo)] =
  		&Value{Type: value.Type, Value: value.Value, Line: value.Line, Written: true}
  }
}

// =============================================================================
// RETRIEVAL FROM VARIABLE
// =============================================================================

func GetVariable(varName string, lineno int) *Value {
  return GetVariableGeneric(varName, varName, lineno)
}

func GetVariableGeneric(errorPrint string, varName string, lineno int) *Value {

  // Check local table first
  if Scope != nil {
    if value1, in1 := Scope[varName]; !in1 {
      // Not in scope
      var value2 *Value
      if value2, in2 := GlobalScope[varName]; !in2 {
        // Not in global
        log.ValueError(lineno, errorPrint)
        return &Value{Written: false, Type: VALUE_UNDEFINED, Line: lineno}
        // In global
      } else if value2.Type == VALUE_UNDEFINED && !value2.Written {
        log.ValueError(lineno, errorPrint)
      }
      return &Value{Type: value2.Type, Value: value2.Value, Written: value2.Written}
    } else {
      // In scope
      if value1.Type == VALUE_UNDEFINED && !value1.Written {
        log.ValueError(lineno, errorPrint)
      }
      return &Value{Type: value1.Type, Value: value1.Value, Written: value1.Written}
    }
  } else {
    // No scope exists
    if value, in := GlobalScope[varName]; !in {
      // Not in global
      log.ValueError(lineno, errorPrint)
      return &Value{Written: false, Type: VALUE_UNDEFINED, Line: lineno}
    } else {
      // In global
      if value.Type == VALUE_UNDEFINED && !value.Written {
        log.ValueError(lineno, errorPrint)
      }
      return &Value{Type: value.Type, Value: value.Value, Written: value.Written}
    }
  }

}

// =============================================================================
// RETRIEVAL FROM OBJECT KEY
// =============================================================================

func GetObjectMember(objectName string, keyName string, lineno int) *Value {

  // Find the object itself
  object := GetVariable(objectName, lineno)
  if object.Type == VALUE_UNDEFINED {
    log.ValueError(lineno, objectName)
  }
  if object.Type != VALUE_OBJECT {
    log.TypeViolation(lineno)
    return &Value{Type: VALUE_UNDEFINED, Line: lineno}
  }

  // Set this object's fields as the new scope then restore it later
  // This is, quite possibly, the smartest thing I have ever devised. Lol. Lol. Lol.
  // Full of himself meter = 10/10. Chance of anyone ever seeing this: 1/10
  oldScope := Scope
  Scope = object.Value.(map[string]*Value)
  keyValue := GetVariableGeneric(objectName + "." + keyName, keyName, lineno)
  Scope = oldScope

  // All other error handling is done in gvg()
  return &Value{Type: keyValue.Type, Value: keyValue.Value, Written: keyValue.Written, Line: keyValue.Line}

}

// =============================================================================
// RETRIEVAL FROM OBJECT KEY
// =============================================================================

func GetArrayMember(arrayName string, index int, lineno int) *Value {

  // Find the array itself
  array := GetVariable(arrayName, lineno)
  if array.Type == VALUE_UNDEFINED {
    log.ValueError(lineno, arrayName)
  }
  if array.Type != VALUE_ARRAY {
    log.TypeViolation(lineno)
    return &Value{Type:VALUE_UNDEFINED, Line: lineno}
  }

  oldScope := Scope
  Scope = array.Value.(map[string]*Value)
  keyValue := GetVariableGeneric(fmt.Sprintf("%s[%d]", arrayName, index), strconv.Itoa(index), lineno)
  Scope = oldScope

  return &Value{Type: keyValue.Type, Value: keyValue.Value, Written: keyValue.Written, Line: keyValue.Line}

}
