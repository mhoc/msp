
package symbol

// The struct our table stores (symbol.Type)
// There are essentially four types of variables our table can store:
//  int : stored directly
//  string : stored directly
//  ast.Object : the ast.Object struct is stored directly
//  Undefined : Below
// If the variable is undefined, then the Undefined value is set to true
// and the content of Value is undefined behavior
type Type struct {
  Undefined bool
  Value interface{}
}
