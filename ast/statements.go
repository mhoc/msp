// Anything to do with single or groups of statements
// STATEMENT LIST

package ast

import (
	"mhoc.co/msp/log"
)

// ====================
// A StatementList is any ordered collection of independent statements
// This is the type of the root Node, but also the type of things like if bodies
// ====================
type StatementList struct {
	List []*Statement
	Line int
}

func (s StatementList) Execute() interface{} {
	for _, child := range s.List {
		potentialJump := child.Execute()
		switch potentialJump.(type) {
		case Break, Continue:
			return potentialJump
		case Return:
			return potentialJump.(Return).Value.Execute().(Value)
		}
	}
	return nil
}

func (s StatementList) LineNo() int {
	return s.Line
}

// ====================
// A single statement in a statement list
// This is broken out into its own
// ====================
type Statement struct {
	N Node
	ErrorHasBeenReported bool
	Line int
}

func (s *Statement) Execute() interface{} {
	v := s.N.Execute()
	if log.ErrorToReport && !s.ErrorHasBeenReported {
		s.ErrorHasBeenReported = true
		log.ErrorReport.Report()
	}
	log.ErrorToReport = false
	return v
}

func (s Statement) LineNo() int {
	return s.Line
}
