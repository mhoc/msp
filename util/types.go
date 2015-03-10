
package util

/** Miscellaneous type definitions which don't deserve their own package */

// Objects are anything such as { a: 5, b: "hello" }
// Our implementation supports nested objects
// We store them as a generic interface{} (java's Object type),
// then type-coerse them later through something like
//    switch var.(type) {
//      case int: vari := var.(int) ...
//    }
// Isn't Go cool?
type Object map[string]interface{}

// Fields are singular keyvalue pairs inside objects
// We dont store these inside objects, but instead simply use them for type
// safety during the semantic parsing process
type Field struct {
  Name string
  Value interface{}
}

// Reserved words should be pretty obvious
// We use them during the lexing process
// These aren't used in the AST because we use custom structures for each
// type of language construct. And we don't store things like start-tag, end-tag,
// and var.
type ReservedWord int
const (
  RW_START_TAG ReservedWord = iota
  RW_END_TAG ReservedWord = iota
  RW_F_DOCWRITE ReservedWord = iota
  RW_VAR ReservedWord = iota
)
