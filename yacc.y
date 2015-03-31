
%{

package main

import (
  "fmt"
  "mhoc.co/msp/ast"
  "mhoc.co/msp/log"
)

// Each of the node types is stored in our %union
// The return types of each grammar rule is written in a comment above the rule
// If we know the type, we use a specific element in this struct so as to
// cut down on type inferencing
// If we dont know the type (aka: statement), then we just use the "No" element

%}

%union {
  N ast.Node
  Str string
  Int int
}

%token
	SCRIPT_TAG_START
	SCRIPT_TAG_END
	VARDEF
	IDENTIFIER
  OBJKEY
	DOCUMENT_WRITE
	NEWLINE
	WHITESPACE
	SEMICOLON
	EQUAL
	INTEGER
	PLUS
	MINUS
	MULT
	DIVIDE
	STRING
	LPAREN
	RPAREN
	COMMA
	LBRACE
	RBRACE
	COLON

%%

target:
  file {
    log.Trace("grm", "Target")
    fmt.Print("")
    if log.LOG_AST {
      $1.N.Print("")
    }
    $1.N.Execute()
  }
;

// File -> StatementList
// Beginning and end script tags with a program in-between them
file:
	SCRIPT_TAG_START newlines program SCRIPT_TAG_END {
    log.Trace("grm", "File: start,newlines,program,end")
    $$.N = $3.N
  }
	| SCRIPT_TAG_START newlines program SCRIPT_TAG_END newlines {
    log.Trace("grm", "File: start,newlines,program,end,newlines")
    $$.N = $3.N
  }
	| newlines SCRIPT_TAG_START newlines program SCRIPT_TAG_END newlines {
    log.Trace("grm", "File: newlines,start,newlines,program,end,newlines")
    $$.N = $4.N
  }
;

// Program -> StatementList
// A list of statement lines each separated by a newlines
program:
	program line {
    log.Trace("grm", "Program: Appending program line")
    line_statements := $2.N.(*ast.StatementList).List
    for _, item := range line_statements {
      $1.N.(*ast.StatementList).List = append($1.N.(*ast.StatementList).List, item)
    }
    $$.N = $1.N
  } newlines
  | {
    log.Trace("grm", "Program: Creating new statement list")
    $$.N = &ast.StatementList{Line: log.LineNo, List: make([]ast.Node, 0, 0)}
  }
;

// Line -> StatementList
// A single line in the program. A line can contain multiple statements through
// the use of a semicolon
line:
	statement {
    log.Trace("grm", "Line: Creating a new single statement statementlist")
    $$.N = &ast.StatementList{Line: log.LineNo, List: []ast.Node{$1.N}}
  }
	| statement SEMICOLON {
    log.Trace("grm", "Line Creating a new single statement statementlist;")
    $$.N = &ast.StatementList{Line: log.LineNo, List: []ast.Node{$1.N}}
  }
	| statement SEMICOLON line {
    // Prepend this statement to the list already created above
    // Because of the weird way the recursion is set up here, we have to prepend instead of append
    // This is the idiomatic way to prepend in go. Looks weird. It works.
    $3.N.(*ast.StatementList).List = append([]ast.Node{$1.N}, $3.N.(*ast.StatementList).List...)
    $$.N = $3.N
  }
;

// Statement -> Node
// Any single statement in the program. Statements have no value in this language.
statement:
	declaration
	| assignment
	| definition
	| DOCUMENT_WRITE LPAREN parameter_list RPAREN {
    $3.N.(*ast.FunctionCall).Name = "document.write"
    $$.N = $3.N
  }
;

// Declaration -> Declaration
// The declaration or redeclaration of a variable
declaration:
	VARDEF IDENTIFIER {
    $$.N = &ast.Declaration{Line: log.LineNo, Name: $2.Str}
  }
;

// Assignment -> Assignment
// The assignment of a value to a variable which has already been declared
assignment:
	IDENTIFIER EQUAL value {
    log.Trace("grm", "Assignment Identifier")
    $$.N = &ast.Assignment{Line: log.LineNo, Name: $1.Str, Rhs: $3.N}
  }
  | OBJKEY EQUAL value {
    log.Trace("grm", "Assignment Obj Key")
    $$.N = &ast.Assignment{Line: log.LineNo, Name: $1.Str, Rhs: $3.N}
  }
;

// Definition -> Definition
// A combination declaration and assignment
definition:
	VARDEF IDENTIFIER EQUAL value {
    // Create the declaration
    decl := &ast.Declaration{Name: $2.Str}
    // Create the assignment
    assign := &ast.Assignment{Name: $2.Str, Rhs: $4.N}
    // Combine them into a definition
    def := &ast.Definition{Decl: decl, Assign: assign}
    $$.N = def
  }
;

// Parameter List -> FunctionCall
// The list of parameters which a function is called.
// This is where we build the actual function call that gets added to the ast
parameter_list:
	expression {
    // Create a new function call with a single argument
    fc := &ast.FunctionCall{Line: log.LineNo, Args: []ast.Node{$1.N}}
    $$.N = fc
  }
	| parameter_list COMMA expression {
    // Append this expression to our list of arguments from $1
    $1.N.(*ast.FunctionCall).Args = append($1.N.(*ast.FunctionCall).Args, $3.N)
    $$.N = $1.N
  }
	| {
    // Create an empty argument function call
    $$.N = &ast.FunctionCall{Line: log.LineNo, Args: []ast.Node{}}
  }
;

// Value -> Node
// Anything in the language which can be assigned to a variable
value:
	expression
	| object_definition
;

// Expression -> Node
// Any combination of multiple sub-expressions to produce a single value
expression:
	additive_expression
;

// Additive Expression -> Node
// Order of operations level 3
additive_expression:
	multiplicative_expression
	| additive_expression PLUS multiplicative_expression {
    $$.N = &ast.Add{Line: log.LineNo, Lhs: $1.N, Rhs: $3.N}
  }
	| additive_expression MINUS multiplicative_expression {
    $$.N = &ast.Subtract{Line: log.LineNo, Lhs: $1.N, Rhs: $3.N}
  }
;

// Multiplicative Expression -> Node
// Order of operations level 2
multiplicative_expression:
	primary_expression
	| multiplicative_expression MULT primary_expression {
    $$.N = &ast.Multiply{Line: log.LineNo, Lhs: $1.N, Rhs: $3.N}
  }
	| multiplicative_expression DIVIDE primary_expression {
    $$.N = &ast.Divide{Line: log.LineNo, Lhs: $1.N, Rhs: $3.N}
  }
;

// Primary Expression -> Node
// Order of operations level 1
primary_expression:
	INTEGER
	| STRING
	| variable_reference
	| LPAREN expression RPAREN {
    $$.N = $2.N
  }
;

// Variable Reference -> Reference
// Any usage of a variable inside a value
variable_reference:
	IDENTIFIER {
    // Create the reference object
    // We dont actually look up and store the value of the variable until execution
    $$.N = &ast.Reference{Line: log.LineNo, Name: $1.Str}
  }
  | OBJKEY {
    $$.N = &ast.Reference{Line: log.LineNo, Name: $1.Str}
  }
;

// Object definition -> Object
// The typed definition of an object inside the source code
object_definition:
	LBRACE field_list RBRACE {
    $$.N = $2.N
  }
	| LBRACE newlines field_list RBRACE {
    $$.N = $3.N
  }
;

// Field List -> Object
// The list of fields without the braces around them
field_list:
	interim_field_list final_field {
    // Add the final field to the list of all the fields
    $1.N.(*ast.Object).Map[$2.N.(*ast.Field).FieldName] = $2.N.(*ast.Field).FieldValue
    $$.N = $1.N
  }
	| {
    // Return an empty object
    $$.N = &ast.Object{Line: log.LineNo, Map: make(map[string]ast.Node)}
  }
;

// Interim Field List -> Object
// This is every field in the object definition except for the last one
// due to the fact that the last one is the only one without a comma after it
interim_field_list:
	interim_field_list interim_field {
    // Add the interim field to the list of all interim fields
    $1.N.(*ast.Object).Map[$2.N.(*ast.Field).FieldName] = $2.N.(*ast.Field).FieldValue
    $$.N = $1.N
  }
	| {
    // Return an empty list of interim fields
    $$.N = &ast.Object{Line: log.LineNo, Map: make(map[string]ast.Node)}
  }
;

// Interim field -> Field
// A single field followed by a required comma
interim_field:
	field COMMA
	| field COMMA newlines
;

// Final field -> Field
// A single field followed by no comma
final_field:
	field
	| field newlines
;

// Field -> Field
// A single key:value pair
field:
	IDENTIFIER COLON expression {
    $$.N = &ast.Field{Line: log.LineNo, FieldName: $1.Str, FieldValue: $3.N}
  }
;

// New Lines -> Nothing
// This is so weird and I hate it but it works
// Previously I had '\n+' as NEWLINE in my lexer, but I wanted to be able to maintain
// my own linenumber count so I changed it to '\n'. Then everything stopped
// working if the user had more than 1 newline. So this is an emulation of
// the \n+ behavior. Oh yes.
newlines:
  NEWLINE {
    log.LineNo++
  }
  | newlines NEWLINE {
    log.LineNo++
  }
;

%%
