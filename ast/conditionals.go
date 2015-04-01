
package ast

import (
  "fmt"
  "mhoc.co/msp/log"
)

// ======================
// An entire if statement
// ======================
type If struct {
  Branches []*Branch
  HasElse bool
  Else *StatementList
  Line int
}

func (i If) Execute() interface{} {
  log.Tracef("ast", "Executing if statement with %d branches", len(i.Branches))

  for _, branch := range i.Branches {
    if branch.Execute().(bool) == true {
      return nil
    }
  }

  // At this point, we can execute the else branch if it exists
  if i.HasElse {
    i.Else.Execute()
  }
  return nil

}

func (i If) LineNo() int {
  return i.Line
}

func (i If) Print(p string) {
  fmt.Println("")
}

// ==================================
// A single branch in an if statement
// ==================================
type Branch struct {
  Conditional Node
  IfTrue *StatementList
  Line int
}

func (b Branch) Execute() interface{} {
  log.Trace("ast", "Traversing branch")

  // Execute the conditional node
  cond := b.Conditional.Execute().(*Value)
  // Convert it to a boolean
  cond = cond.ToBoolean()

  // If it is undefined, throw a condition error and nope out of here
  if cond.Type == VALUE_UNDEFINED {
    log.Error{Line: b.Line, Type: log.CONDITION}.Report()
    return true
  }

  // If it is false, just return false
  if !cond.Value.(bool) {
    return false
  }

  // If true, execute the statement list
  b.IfTrue.Execute()
  return true

}

func (b Branch) LineNo() int {
  return b.Line
}

func (b Branch) Print(p string) {
  fmt.Println("")
}
