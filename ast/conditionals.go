
package ast

import (
  "mhoc.co/msp/log"
)

// An if statement
type Branch struct {
  Conditional Node
  IfTrue *StatementList
  IfFalse *StatementList
  Line int
}

func (b Branch) Execute() interface{} {
  log.Trace("ast", "Traversing branch")

  // Execute the conditional node
  //cond := b.Conditional.Execute().(*Value)

  return nil
}

func (b Branch) LineNo() int {
  return b.Line
}

func (b Branch) Print(p string) {

}
