package ast

// ===================
// Array Type
// This is essentially just an ast.Object, but it is a separate type
// for the sake of execution
// ===================
type Array struct {
	Map  map[string]Node
	Line int
}

func (a Array) Execute() interface{} {
	// Build the new value
	v := Value{Type: VALUE_ARRAY, Value: make(map[string]Value)}
	for key, value := range a.Map {
		v.Value.(map[string]Value)[key] = value.Execute().(Value)
	}
	return v
}

func (a Array) LineNo() int {
	return a.Line
}
