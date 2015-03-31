
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
  Args []Node
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
    switch argv.Type {
      case VALUE_UNDEFINED:
        fmt.Printf("undefined")
        break
      case VALUE_INT:
        fmt.Printf("%d", argv.Value)
        break
      case VALUE_STRING:
        if argv.Value == "<br />" {
          argv.Value = "\n"
        }
        fmt.Printf("%s", argv.Value)
        break
      default:
        log.Error{Line:f.Line, Type: log.TYPE_VIOLATION, Msg: "Attempting to print a non-int or string type"}.Report()
        fmt.Print("undefined")
    }
  }
  return nil
}

func (f FunctionCall) LineNo() int {
  return f.Line
}

func (f FunctionCall) Print(p string) {
  fmt.Printf(p + "Call\n")
  fmt.Printf(p + "| %s\n", f.Name)
  for _, arg := range f.Args {
    arg.Print(p + "| ")
  }
}
