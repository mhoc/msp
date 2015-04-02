package ast

import (
	"mhoc.co/msp/log"
)

// ========================
// General Unary Expression
// ========================
type UnaryExpression struct {
	Op    string
	Value Node
	Line  int
}

func (ue UnaryExpression) Execute() interface{} {
	log.Tracef("ast", "Executing binary expression %s", ue.Op)

	// Execute the value
	value := ue.Value.Execute().(*Value)

	// Switch on the operator
	switch ue.Op {
	case "!":
		return handleNot(value, ue.Line)
	}

	panic("Supplied a unary operator not supported")

}

func (ue UnaryExpression) LineNo() int {
	return ue.Line
}

func handleNot(v *Value, line int) *Value {

	if v.Type == VALUE_BOOLEAN {
		v.Value = !v.Value.(bool)
		return v
	}

	if v.Type == VALUE_INT {
		if v.Value.(int) == 0 {
			v.Type = VALUE_BOOLEAN
			v.Value = true
		} else {
			v.Type = VALUE_BOOLEAN
			v.Value = false
		}
		return v
	}

	if v.Type == VALUE_STRING {
		if len(v.Value.(string)) == 0 {
			v.Type = VALUE_BOOLEAN
			v.Value = true
		} else {
			v.Type = VALUE_BOOLEAN
			v.Value = false
		}
		return v
	}

	log.Error{Line: line, Type: log.TYPE_VIOLATION}.Report()
	v.Type = VALUE_UNDEFINED
	return v

}
