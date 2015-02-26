
%{

#include "my.h"
int yylineno;
#include "helper.c"
FILE *yyin;

%}

%token
	SCRIPT_TAG_START
	SCRIPT_TAG_END
	VARDEF
	IDENTIFIER
	DOCUMENT_WRITE
	NEWLINE
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
	VARDEF IDENTIFIER {
		declare_symbol($2.value.string);
	}
;

/** a = 1
		A corresponding update_value function is called.
		update_value will print an error if the value was not previously declared. */
assignment:
	IDENTIFIER EQUAL value {
		if ($3.type == TYPE_INTEGER) {
			update_symbol_i($1.value.string, $3.value.number);
		}
		else if ($3.type == TYPE_STRING) {
			update_symbol_s($1.value.string, $3.value.string);
		}
		else if ($3.type == TYPE_OBJECT) {

		}
	}
;

/** var a = 1
		A combination declaration and assignment
		insert_symbol is called based upon the type of the value passed in.
		this function will redefine the type of the variable if it was already declared. */
definition:
	VARDEF IDENTIFIER EQUAL value {
		declare_symbol($2.value.string);
		if ($4.type == TYPE_INTEGER) {
			update_symbol_i($2.value.string, $4.value.number);
		}
		else if ($4.type == TYPE_STRING) {
			update_symbol_s($2.value.string, $4.value.string);
		}
		else if ($4.type == TYPE_OBJECT) {

		}
	}
;

/** A list of comma-separated values
		Right now this is hard-coded to only work with document.write() */
parameter_list:
	expression {
		if ($1.type == TYPE_STRING) {
			if (strcmp($1.value.string, "<br />")) {
				printf("%s", $1.value.string);
			} else {
				printf("\n");
			}
		}
		else if ($1.type == TYPE_INTEGER) {
			printf("%i", $1.value.number);
		} else {
			printf("Error (line %d): Attempting to print a non-integer or string value\n");
		}
	}
	| parameter_list COMMA expression {
		if ($3.type == TYPE_STRING) {
			if (strcmp($3.value.string, "<br />")) {
				printf("%s", $3.value.string);
			} else {
				printf("\n");
			}
		}
		else if ($3.type == TYPE_INTEGER) {
			printf("%i", $3.value.number);
		} else {
			printf("Error (line %d): Attempting to print a non-integer or string value\n");
		}
	}
	| /* empty */
;

/** A value is any entity which can be assigned to a variable.
		expressions, in particular, are the only ones that can be recursed into
		sub-expressions AND are the only ones that can be printed. */
value:
	expression
	| object_definition

/** Any reference to a variable which has been previously declared and defined.
    If it has not been declared or defined, a type error is printed */
variable_reference:
	IDENTIFIER {
		debugf1("(%s) ", $1.value.string)

		struct entity* e = get_symbol($1.value.string);

		if (e == NULL) {
			printf("Variable Access Error (line %d): Use of undeclared variable %s\n", yylineno, $1.value.string);
			$$.type = TYPE_UNDEFINED;
		}
		else if (e->type == TYPE_INTEGER) {
			$$.name = e->name;
			$$.type = e->type;
			$$.value.number = e->value.number;

		} else if (e->type == TYPE_STRING) {
			$$.name = e->name;
			$$.type = e->type;
			$$.value.string = e->value.string;
		}

	}
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
	| additive_expression PLUS multiplicative_expression {
		if ($1.type == TYPE_INTEGER && $3.type == TYPE_INTEGER) {
			$1.value.number += $3.value.number;
			$$ = $1;
		}
		else if ($1.type == TYPE_STRING && $3.type == TYPE_STRING) {
			char* newstr = malloc(strlen($1.value.string) + strlen($3.value.string));
			strcpy(newstr, $1.value.string);
			strcat(newstr, $3.value.string);
			$1.value.string = newstr;
			$$ = $1;
		} else {
			printf("Type Violation Error (line %d): Attempting to apply addition to unsupported types\n", yylineno);
		}
	}
	| additive_expression MINUS multiplicative_expression {
		if ($1.type == TYPE_INTEGER && $3.type == TYPE_INTEGER) {
			$1.value.number -= $3.value.number;
			$$ = $1;
		}
		else {
			printf("Type Violation Error (line %d): Attempting to apply subtraction to unsupported types\n", yylineno);
		}
	}
;

multiplicative_expression:
	primary_expression
	| multiplicative_expression MULT primary_expression {
		if ($1.type == TYPE_INTEGER && $3.type == TYPE_INTEGER) {
			$1.value.number *= $3.value.number;
			$$ = $1;
		}
		else {
			printf("Type Violation (line %d): Attempting to apply multiplication to unsupported types\n", yylineno);
		}
	}
	| multiplicative_expression DIVIDE primary_expression {
		if ($1.type == TYPE_INTEGER && $3.type == TYPE_INTEGER) {
			$1.value.number /= $3.value.number;
			$$ = $1;
		}
		else {
			printf("Type Violation (line %d): Attempting to apply division to unsupported types\n", yylineno);
		}
	}
;

primary_expression:
	INTEGER
	| STRING
	| variable_reference
	| LPAREN expression RPAREN {
		$$ = $2;
	}
;

/** { key:value ... } */
object_definition:
	LBRACE field_list RBRACE
	| LBRACE NEWLINE field_list RBRACE
;

field_list:
	interim_field_list final_field
	| /* empty */
;

interim_field_list:
	interim_field_list interim_field
	| /* empty */
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

void segvhandler(int sig, siginfo_t *si, void *unused) {
	printf("\n\n==SEGMENTATION FAULT==\n");
	fflush(stdout);
	fflush(stderr);
	exit(0);
}

yyerror(char *s) {
    fprintf(stderr, "error: %s, line: %d\n", s, yylineno);
    return 1;
}

int main(int argc, char *argv[]) {

	signal(SIGSEGV, segvhandler);

	if (argc == 1) {
		yyin = stdin;
		yyparse();
	}
	else if (argc == 2) {
		FILE *file;
		file = fopen(argv[1], "r");
		if (!file) {
      		fprintf(stderr, "could not open %s\n", argv[1]);
  	} else {
      		yyin = file;
      		//yyparse() will call yylex()
      		yyparse();
  	}

	} else {
        	fprintf(stderr, "format: ./yacc_example [filename]");
	}

	return 0;
}
