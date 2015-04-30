// Primitive types in the language
// aka anything which the lexer directly returns with no semantic parsing knowledge
// VALUE

package ast

import (
	"fmt"
	"sort"
	"strconv"
	"mhoc.co/msp/log"
)

// ===================
// Value
// Represents a binding between a go primitive type and a ms
// primitive type. Also represents undefined types
// Values have types which are defined by an enum
// We also store functions in values so they can be put in the symbol table,
// as per the definition that function names and var names will never
// overlap
// ===================
type ValueType int

const (
	VALUE_UNDEFINED ValueType = iota // Disregard value of Value
	VALUE_INT       ValueType = iota // type(Value) == int
	VALUE_STRING    ValueType = iota // type(Value) == string
	VALUE_OBJECT    ValueType = iota // type(Value) == map[string]Value
	VALUE_BOOLEAN   ValueType = iota // type(Value) == bool
	VALUE_ARRAY     ValueType = iota // type(Value) == map[string]Value
	VALUE_FUNCTION 	ValueType = iota // type(Value) == FunctionDef
)

type Value struct {
	Type    ValueType
	Value   interface{}
	Line    int
	Written bool
}

func (v Value) Execute() interface{} {
	// We just return the value itself, not the containing interface
	// because we need information about its type in parent ast nodes
	return v
}

func (v Value) LineNo() int {
	return v.Line
}

// Commits type coercion on a value to convert it to an msp boolean value
func (v Value) ToBoolean() *Value {
	nv := &Value{Type: VALUE_BOOLEAN, Line: v.Line}
	switch v.Type {
	case VALUE_BOOLEAN:
		nv.Value = v.Value.(bool)
	case VALUE_INT:
		nv.Value = v.Value.(int) != 0
	case VALUE_STRING:
		nv.Value = len(v.Value.(string)) > 0
	default:
		nv.Type = VALUE_UNDEFINED
	}
	return nv
}

// Converts a value to a string
// The result of this function can be passed directly into
// fmt.Print() for document.write().
// It does not handle errors, however.
func (v Value) ToString() string {
	switch v.Type {
		case VALUE_UNDEFINED:
			return "undefined"
		case VALUE_INT:
			return fmt.Sprintf("%v", v.Value.(int))
		case VALUE_STRING:
			if v.Value.(string) == "<br />" {
				return "\n"
			}
			return fmt.Sprintf("%v", v.Value.(string))
		case VALUE_BOOLEAN:
			if v.Value.(bool) {
				return "true"
			} else {
				return "false"
			}
		case VALUE_OBJECT:
			if log.EXTENSIONS {
				return v.ObjToString()
			} else {
				return "undefined"
			}
		case VALUE_ARRAY:
			if log.EXTENSIONS {
				return v.ArrayToString()
			} else {
				return "undefined"
			}
		case VALUE_FUNCTION:
			if log.EXTENSIONS {
				vf := v.Value.(FunctionDef)
				sp := vf.Name + "("
				for i, argn := range vf.ArgNames {
					sp += argn
					if i != len(vf.ArgNames)-1 {
						sp += ","
					}
				}
				sp += "){}"
				return sp
			} else {
				return "undefined"
			}
		default:
			return "undefined"
	}
}

func (v Value) ObjToString() string {
	return "obj fill in here"
}

func (v Value) ArrayToString() string {
	keys := []int{}
	ar := v.Value.(map[string]*Value)
	for key, _ := range ar {
		keyindex, _ := strconv.Atoi(key)
		keys = append(keys, keyindex)
	}
	sort.Ints(keys)
	res := "["
	for _, key := range keys {
		res += ar[strconv.Itoa(key)].ToString()
		res += ", "
	}
	res += "\b\b]"
	return res
}
