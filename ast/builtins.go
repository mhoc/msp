
package ast

import (
  "fmt"
  "mhoc.co/msp/log"
)

// Initializes builtin functions and definitions
func InitBuiltins() {

  // Document.write
  Declare("document.write")
  fd := FunctionDef{
    Name: "document.write", ArbitraryArgs: true, ArgNames: []string{},
    ExecMiniscript: false, GoBody: DocumentWrite, Line: -1,
  }
  v := Value{Type: VALUE_FUNCTION, Value: fd, Line: -1, Written: true}
  AssignToVariable("document.write", v, -1)

  // Assert

}

func DocumentWrite(fc FunctionCall) interface{} {
  for _, arg := range fc.Args {
    argv := arg.Execute().(*Value)
    fmt.Print(argv.ToString())
    switch argv.Type {
      case VALUE_OBJECT:
        if !log.EXTENSIONS {
          log.TypeViolation(fc.LineNo())
        }
      case VALUE_ARRAY:
        if !log.EXTENSIONS {
          log.TypeViolation(fc.LineNo())
        }
    }
  }
  return nil
}
