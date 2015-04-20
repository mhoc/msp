package ast

import (
	"mhoc.co/msp/log"
)

// I. Hate. This.
// Implementing Gotos with pure tree traversal is seriously difficult because
// lets say you are reading a statement list and you hit a break :=> you could
// just pass control up, but this is the incorrect behavior if you arent inside
// a loop. How can you know if you're in a loop without maintaining global state?
// That I don't know. You could keep passing control up until you hit a loop
// node, but that would be SO slow and if you didnt hit one you'd have to
// re-traverse back down without losing your place in the program
var LoopDepth = 0

// Do and While Loop
// PreCheck is set to true if it is a while loop, and false if its a dowhile loop
type Loop struct {
	Conditional Node
	Body        *StatementList
	PreCheck    bool
	Line        int
}

func (l Loop) Execute() interface{} {
	log.Trace("ast", "Executing loop")

	condition := &Value{Type: VALUE_BOOLEAN, Value: true}
	for {

		if l.PreCheck {
			condition = l.Conditional.Execute().(*Value).ToBoolean()
			if condition.Type != VALUE_BOOLEAN {
				log.ConditionError(l.Line)
				return nil
			}
		}

		var breakMet = false
		if condition.Value.(bool) {
			LoopDepth++
			jump := l.Body.Execute()
			LoopDepth--
			switch jump.(type) {
			case Break:
				breakMet = true
			}
		} else {
			break
		}

		if breakMet {
			break
		}

		if !l.PreCheck {
			condition = l.Conditional.Execute().(*Value).ToBoolean()
			if condition.Type != VALUE_BOOLEAN {
				log.ConditionError(l.Line)
				return nil
			}
		}

	}
	return nil

}

func (l Loop) LineNo() int {
	return l.Line
}

// One of the statements in a while loop could be a break statement
type Break struct {
	Line int
}

func (b Break) Execute() interface{} {
	return b
}

func (b Break) LineNo() int {
	return b.Line
}

// Another could be a continue
type Continue struct {
	Line int
}

func (c Continue) Execute() interface{} {
	return c
}

func (c Continue) LineNo() int {
	return c.Line
}
