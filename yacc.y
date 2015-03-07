
%{

package main

type ValueType int
type ReservedWord int

const (
  TYPE_INTEGER ValueType = iota
  TYPE_STRING ValueType = iota
  TYPE_OBJECT ValueType = iota
	TYPE_UNDEFINED ValueType = iota
	TYPE_RESERVED ValueType = iota
)

const (
	RESERVED_START_TAG ReservedWord = iota
	RESERVED_END_TAG ReservedWord = iota
	RESERVED_DOCUMENT_WRITE ReservedWord = iota
	RESERVED_VARDEF ReservedWord = iota
)

%}

%union {
	vtype ValueType
	n int
	s string
	c byte

}

%token
	SCRIPT_TAG_START
	SCRIPT_TAG_END
	VARDEF
	IDENTIFIER
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
	OBJKEY

%%

/** A file is the entire thing we are parsing.
		It takes care of the script tags at the start and end. */
file:
	SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END
	| SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
	| NEWLINE SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
;

/** A program is a series of lines, each followed by a newline. */
program:
	program line NEWLINE
	| /* Empty */
;

/** Each line is a statement terminated by an optional semicolon.
  	This rule takes care of the possibility that a line could contain multiple statements.
		If it does, each one in the center must have a semicolon. The last one is optional. */
line:
	statement
	| statement SEMICOLON
	| statement SEMICOLON line
;

/** A statement is any single operation. This includes declarations (var a), assignments (a = 1),
  	definitions (var a = 1), and function calls (document.write(a)). */
statement:
	declaration
	| assignment
	| definition
	| DOCUMENT_WRITE LPAREN parameter_list RPAREN
;

/** var a
		In a declaration, all we do is add the symbol as undefined in our symbol table. */
declaration:
	VARDEF IDENTIFIER
;

/** a = 1
		A corresponding update_value function is called.
		update_value will print an error if the value was not previously declared. */
assignment:
	IDENTIFIER EQUAL value
	| OBJKEY EQUAL value
;

/** var a = 1
		A combination declaration and assignment
		insert_symbol is called based upon the type of the value passed in.
		this function will redefine the type of the variable if it was already declared. */
definition:
	VARDEF IDENTIFIER EQUAL value
;

/** A list of comma-separated values
		Right now this is hard-coded to only work with document.write() */
parameter_list:
	expression
	| parameter_list COMMA expression
	| /* empty */
;

/** A value is any entity which can be assigned to a variable.
		expressions, in particular, are the only ones that can be recursed into
		sub-expressions AND are the only ones that can be printed. */
value:
	expression
	| object_definition
;

/** Any reference to a variable which has been previously declared and defined.
    If it has not been declared or defined, a type error is printed */
variable_reference:
	IDENTIFIER
	| OBJKEY
;

/** An expression is a combination of multiple subexpressions or values to
		produce a single expression_value. */
expression:
	additive_expression
;

/** These sub-expression rules define an order of operations such that
			1) "a", 1, a, (exp)
			2) exp * exp, exp / exp
			3) exp + exp, exp - exp */
additive_expression:
	multiplicative_expression
	| additive_expression PLUS multiplicative_expression
	| additive_expression MINUS multiplicative_expression
;

multiplicative_expression:
	primary_expression
	| multiplicative_expression MULT primary_expression
	| multiplicative_expression DIVIDE primary_expression
;

primary_expression:
	INTEGER
	| STRING
	| variable_reference
	| LPAREN expression RPAREN
;

/** { key:value ... } */
object_definition:
	LBRACE field_list RBRACE
	| LBRACE NEWLINE field_list RBRACE
;

/** A list of key:value pairs without the braces around them */
field_list:
	interim_field_list final_field
	| /* Empty */
;

/** A list of key:value pairs except the last item in the list
 		The key difference being that the last item has no comma after it */
interim_field_list:
	interim_field_list interim_field
	| /* Empty */
;

interim_field:
	field COMMA
	| field COMMA NEWLINE
;

final_field:
	field
	| field NEWLINE
;

field:
	IDENTIFIER COLON expression
;

%%
