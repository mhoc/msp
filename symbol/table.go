
package symbol

import (
  "fmt"
  "mhoc.co/msp/util"
)

var symbolTable map[string]Symbol*

// Declares a new variable in the symbol table with an undefined value
func Declare(name string) {
  util.LogTrace("Declaring new symbol " + name)
  symbolTable[name] = &Symbol{}
}

func Define(name string, value interface{}) {
  util.LogTrace("Defining symbol value for " + name)

  // Set up our new symbol object
  s := &util.Symbol{Name: name}
  switch value.(type) {
    case int:
      s.Value = value.(int)
      break
    case string:
      s.Value = value.(string)
      break
    case util.Object:
      s.Value = value.(util.Object)
      break
    case util.Token:
      s.Value = value.(util.Token).Value
      break
    default:
      util.InternalError("Attempting to define a type in the symbol table which is not supported")
  }

  // Get the currently defined symbol
  defined := symbolTable[name]
  if defined == nil {

  }

}
