
// Objects are a bit strange
// They could be thought of as primitive because they are directly typed
// in the source code (not derived from semantic parsing), but simultaneously
// each entry in the object is itself an assignment, so it cant be primitive.
// So the Execute() function cannot return anything.

package ast

import "fmt"

// ====================
// Object literals
// When objects are statically typed into the source code, these are generated.
// When an object is executed, it returns a Value of type VALUE_OBJECT
// This is strange but actually makes sense. By definition, Values are
// primitive and can only exist as leaves on an AST tree. Objects are not
// generally leaves. They are nodes themselves with children which are values
// ====================
type Object struct {
  IsArray bool
  Map map[string]Node
  Line int
}

func (o Object) Execute() interface{} {
  // Here, we traverse each value in the object and evaluate it
  // then build a new map (stored in a Value) containing the evaluated values

  // Build the new value
  var v *Value
  if (o.IsArray) {
    v = &Value{Type:VALUE_ARRAY, Value: make(map[string]*Value)}
  } else {
    v = &Value{Type: VALUE_OBJECT, Value: make(map[string]*Value)}
  }

  for key, value := range o.Map {
    v.Value.(map[string]*Value)[key] = value.Execute().(*Value)
  }

  return v

}

func (o Object) LineNo() int {
  return o.Line
}

func (o Object) Print(p string) {
  fmt.Println(p + "Object Node")
  for key, value := range o.Map {
    fmt.Printf(p + "| %s\n", key)
    value.Print(p + "| | ")
  }
}

// ====================
// Field literals
// Ironically these are never actually stored in the AST
// They are just used to pass back field information to another method
// which puts their information into the ast.Object.Map
// ====================
type Field struct {
  FieldName string
  FieldValue Node
  Line int
}

func (f Field) Execute() interface{} {
  // Doesnt have to do anything, but it does have to be an ast.Node
  return nil
}

func (f Field) LineNo() int {
  return f.Line
}

func (f Field) Print(p string) {
  // Again, this will never be called because we don't put fields in the ast
  // That being said, I'll include it just to be safe
  fmt.Println(p + "Field")
  fmt.Printf(p + "| %s\n", f.FieldName)
  f.FieldValue.Print(p + "| ")
}
