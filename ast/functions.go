
package ast

import "fmt"

// ====================
// Any function call
// Right now we only recognize one functionc all: document.write
// ====================
type FunctionCall struct {
  Name string
  Args []Node
}

func (f FunctionCall) Execute() interface{} {
  // For now, assume all function calls are document.write
  // this will be improved later with a function lookup table
  if f.Name != "document.write" {
    panic("Error: Attempting to call function that is not document.write")
  }
  for _, arg := range f.Args {
    argv := arg.Execute()
    switch argv.(type) {
      case int:
        fmt.Printf("%d", argv)
        break
      case string:
        if argv == "<br />" {
          argv = "\n"
        }
        fmt.Printf("%s", argv)
        break
    }
  }
  return nil
}

func (f FunctionCall) Print(p string) {
  fmt.Printf(p + "Call\n")
  fmt.Printf(p + "| %s\n", f.Name)
  for _, arg := range f.Args {
    arg.Print(p + "| ")
  }
}
