// Contains entities which help form expressions

package ast

import (
	"mhoc.co/msp/log"
	"strings"
)

// =========================
// General Binary Expression
// This handles a lot of the error checking associated with undefined values
// in one location
// =========================
type BinaryExpression struct {
	Lhs  Node
	Rhs  Node
	Op   string
	Line int
}

func (be *BinaryExpression) Execute() interface{} {
	log.Tracef("ast", "Executing binary expression %s", be.Op)

	// Execute both sides
	left := be.Lhs.Execute().(*Value)
	right := be.Rhs.Execute().(*Value)

	// If one side is undefined and unwritten, we report a type violation and return undefined
	if (left.Type == VALUE_UNDEFINED && !left.Written) || (right.Type == VALUE_UNDEFINED && !right.Written) {
		log.TypeViolation(be.Line)
		left.Type = VALUE_UNDEFINED
		return left
	}

	// If one side is undefined and written, we just return undefined
	if left.Type == VALUE_UNDEFINED || right.Type == VALUE_UNDEFINED {
		left.Type = VALUE_UNDEFINED
		return left
	}

	// If the types are simply not the same and none are undefined, we report a type violation and return undefined
	// We also dont do this for && and || (truthy) operations b
	if (left.Type != right.Type) && (!opIsCoersive(be.Op)) {
		log.TypeViolation(be.Line)
		left.Type = VALUE_UNDEFINED
		return left
	}

	// Handle each operation separately
	switch be.Op {
	case "+":
		return handlePlus(left, right, be.Line)
	case "-":
		return handleMinus(left, right, be.Line)
	case "*":
		return handleMult(left, right, be.Line)
	case "/":
		return handleDivide(left, right, be.Line)
	case "==":
		return handleEquiv(left, right, be.Line)
	case "!=":
		return handleNequiv(left, right, be.Line)
	case "||":
		return handleOr(left, right, be.Line)
	case "&&":
		return handleAnd(left, right, be.Line)
	case ">":
		return handleGt(left, right, be.Line)
	case "<":
		return handleLt(left, right, be.Line)
	case ">=":
		return handleGte(left, right, be.Line)
	case "<=":
		return handleLte(left, right, be.Line)
	}

	// This should never be reached
	panic("You just supplied a binary operator we dont support")

}

func (be BinaryExpression) LineNo() int {
	return be.Line
}

// =====================================================
// Some functions to clean up the binary expression code
// in handling execution of different operators
// =====================================================

func handlePlus(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Value = left.Value.(int) + right.Value.(int)
		return left
	}

	// Strings
	if left.Type == VALUE_STRING && right.Type == VALUE_STRING {
		lStr := left.Value.(string)
		rStr := right.Value.(string)
		if strings.Contains(lStr, "<br />") || strings.Contains(rStr, "<br />") {
			log.TypeViolation(line)
			left.Type = VALUE_UNDEFINED
			return left
		}
		left.Value = left.Value.(string) + right.Value.(string)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleMinus(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Value = left.Value.(int) - right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleMult(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Value = left.Value.(int) * right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleDivide(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Value = left.Value.(int) / right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleEquiv(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) == right.Value.(int)
		return left
	}

	// Booleans
	if left.Type == VALUE_BOOLEAN && right.Type == VALUE_BOOLEAN {
		left.Value = left.Value.(bool) == right.Value.(bool)
		return left
	}

	// Strings
	if left.Type == VALUE_STRING && right.Type == VALUE_STRING {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(string) == right.Value.(string)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleNequiv(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) != right.Value.(int)
		return left
	}

	// Booleans
	if left.Type == VALUE_BOOLEAN && right.Type == VALUE_BOOLEAN {
		left.Value = left.Value.(bool) != right.Value.(bool)
		return left
	}

	// Strings
	if left.Type == VALUE_STRING && right.Type == VALUE_STRING {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(string) != right.Value.(string)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleAnd(left *Value, right *Value, line int) *Value {

	// Convert the type of the left and right side if they arent boolean
	if left.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) > 0
	}
	if right.Type == VALUE_INT {
		right.Type = VALUE_BOOLEAN
		right.Value = right.Value.(int) > 0
	}
	if left.Type == VALUE_STRING {
		left.Type = VALUE_BOOLEAN
		left.Value = len(left.Value.(string)) > 0
	}
	if right.Type == VALUE_STRING {
		right.Type = VALUE_BOOLEAN
		right.Value = len(right.Value.(string)) > 0
	}

	// Booleans
	if left.Type == VALUE_BOOLEAN && right.Type == VALUE_BOOLEAN {
		left.Value = left.Value.(bool) && right.Value.(bool)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleOr(left *Value, right *Value, line int) *Value {

	// Convert the type of the left and right side if they arent boolean
	if left.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) > 0
	}
	if right.Type == VALUE_INT {
		right.Type = VALUE_BOOLEAN
		right.Value = right.Value.(int) > 0
	}
	if left.Type == VALUE_STRING {
		left.Type = VALUE_BOOLEAN
		left.Value = len(left.Value.(string)) > 0
	}
	if right.Type == VALUE_STRING {
		right.Type = VALUE_BOOLEAN
		right.Value = len(right.Value.(string)) > 0
	}

	// Booleans
	if left.Type == VALUE_BOOLEAN && right.Type == VALUE_BOOLEAN {
		left.Value = left.Value.(bool) || right.Value.(bool)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleGt(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) > right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleLt(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) < right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleGte(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) >= right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func handleLte(left *Value, right *Value, line int) *Value {

	// Integers
	if left.Type == VALUE_INT && right.Type == VALUE_INT {
		left.Type = VALUE_BOOLEAN
		left.Value = left.Value.(int) <= right.Value.(int)
		return left
	}

	log.TypeViolation(line)
	left.Type = VALUE_UNDEFINED
	return left

}

func opIsCoersive(op string) bool {
	switch op {
	case "&&", "||":
		return true
	}
	return false
}
