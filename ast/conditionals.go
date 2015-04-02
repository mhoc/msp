
package ast

import (
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
    rVal := branch.Execute()
    switch rVal.(type) {
      case bool:
        if rVal.(bool) {
          return nil
        }
      case Break, Continue:
        return rVal
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
  potentialJump := b.IfTrue.Execute()
  switch potentialJump.(type) {
    case Break, Continue:
      return potentialJump
  }
  return true

}

func (b Branch) LineNo() int {
  return b.Line
}
