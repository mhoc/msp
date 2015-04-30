// Objects are a bit strange
// They could be thought of as primitive because they are directly typed
// in the source code (not derived from semantic parsing), but simultaneously
// each entry in the object is itself an assignment, so it cant be primitive.
// So the Execute() function cannot return anything.

package ast

// ====================
// Object literals
// When objects are statically typed into the source code, these are generated.
// When an object is executed, it returns a Value of type VALUE_OBJECT
// This is strange but actually makes sense. By definition, Values are
// primitive and can only exist as leaves on an AST tree. Objects are not
// generally leaves. They are nodes themselves with children which are values
// ====================
type Object struct {
	Map  map[string]Node
	Line int
}

func (o Object) Execute() interface{} {
	// Here, we traverse each value in the object and evaluate it
	// then build a new map (stored in a Value) containing the evaluated values

	// Build the new value
	v := &Value{Type: VALUE_OBJECT, Value: make(map[string]Value)}
	for key, value := range o.Map {
		v.Value.(map[string]Value)[key] = *value.Execute().(*Value)
	}

	return v

}

func (o Object) LineNo() int {
	return o.Line
}

// ====================
// Field literals
// Ironically these are never actually stored in the AST
// They are just used to pass back field information to another method
// which puts their information into the ast.Object.Map
// ====================
type Field struct {
	FieldName  string
	FieldValue Node
	Line       int
}

func (f Field) Execute() interface{} {
	// Doesnt have to do anything, but it does have to be an ast.Node
	return nil
}

func (f Field) LineNo() int {
	return f.Line
}
