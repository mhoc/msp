
// Objects are a bit strange
// They could be thought of as primitive because they are directly typed
// in the source code (not derived from semantic parsing), but simultaneously
// each entry in the object is itself an assignment, so it cant be primitive.
// So the Execute() function cannot return anything.

package ast

import "fmt"

// ===================
// Object literals
// {a: "hello", b: 5}
// ===================
type Object struct {
  Map map[string]Node
  Line int
}

func (o Object) Execute() interface{} {
  // Im not actually sure if anything needs to be done here
  // We definitely dont return anything (becuase objects do not have value and cannot be used in expressions)
  // We dont need to insert each kv pair into our symbol table because the
  // caller of this function will handle inserting the entire object node
  // into the table.
  return nil
}

func (o Object) LineNo() int {
  return o.Line
}

func (o Object) Print(p string) {
  fmt.Println(p + "Object")
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
