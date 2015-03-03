
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
	VARDEF IDENTIFIER {
		declareSymbol($2.value.string);
	}
;

/** a = 1
		A corresponding update_value function is called.
		update_value will print an error if the value was not previously declared. */
assignment:
	IDENTIFIER EQUAL value {
		defineSymbol($1.value.string, &$3, 1);
	}
	| OBJKEY EQUAL value {
		defineSymbol($1.value.string, &$3, 0);
	}
;

/** var a = 1
		A combination declaration and assignment
		insert_symbol is called based upon the type of the value passed in.
		this function will redefine the type of the variable if it was already declared. */
definition:
	VARDEF IDENTIFIER EQUAL value {
		declareSymbol($2.value.string);
		defineSymbol($2.value.string, &$4, 1);
	}
;

/** A list of comma-separated values
		Right now this is hard-coded to only work with document.write() */
parameter_list:
	expression {
		printExpression(&$1);
	}
	| parameter_list COMMA expression {
		printExpression(&$3);
	}
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
	IDENTIFIER {
		$$ = *getSymbol($1.value.string);
	}
	| OBJKEY {
		$$ = *getSymbol($1.value.string);
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
		$$ = addTokens($1, $3);
	}
	| additive_expression MINUS multiplicative_expression {
		$$ = subtractTokens($1, $3);
	}
;

multiplicative_expression:
	primary_expression
	| multiplicative_expression MULT primary_expression {
		$$ = multiplyTokens($1, $3);
	}
	| multiplicative_expression DIVIDE primary_expression {
		$$ = divideTokens($1, $3);
	}
;

primary_expression:
	INTEGER
	| STRING
	| variable_reference {
		if ($1.type == TYPE_FIELDLIST) {
			printf("Line %d, type violation\n", yylineno);
		}
	}
	| LPAREN expression RPAREN {
		$$ = $2;
	}
;

/** { key:value ... } */
object_definition:
	LBRACE field_list RBRACE {
		$$ = $2;
	}
	| LBRACE NEWLINE field_list RBRACE {
		$$ = $3;
	}
;

field_list:
	interim_field_list final_field {
		addToFieldList($1.value.fieldList, &$2);
		$$ = $1;
	}
	| {
		$$ = *newFieldList();
	}
;

interim_field_list:
	interim_field_list interim_field {
		addToFieldList($1.value.fieldList, &$2);
		$$ = $1;
	}
	| {
		$$ = *newFieldList();
	}
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
	IDENTIFIER COLON expression {
		$$ = *createField($1.value.string, &$3);
	}
;

%%

void segvhandler(int sig, siginfo_t *si, void *unused) {
	printf("\n\n==SEGMENTATION FAULT==\n");
	fflush(stdout);
	fflush(stderr);
	exit(0);
}

yyerror(char *s) {
    fprintf(stderr, "%s\n", s);
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
