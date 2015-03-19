
// Contains the node definition
// NODE

package ast

// ====================
// Every Node in our AST decends from this Node interface
// ====================
type Node interface {

  // Execute is a function that "executes" the function of a node in the AST
  // This is the core of the compiler design. We build up an AST during lexing and semantic
  // analysis, then call Execute() on the root node of the ast, which calls its children's
  // execute function and so on
  // The leaf node types will have an empty or non-recursive execute function
  // Execute can provide an optional return value if the node being executed makes sense
  // to return something (say, a literal or variable reference)
  Execute() interface{}

  // We provide printing functionality for a visual representation of the AST contained
  // in ast/print.go. The process for this is very similar to Execute()
  Print(prefix string)

}
