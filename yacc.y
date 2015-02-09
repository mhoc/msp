%{
#include <stdio.h>
#include <stdlib.h>
#include "my.h"
%}

%token 	SCRIPT_TAG_START SCRIPT_TAG_END 
	VARDEF IDENTIFIER DOCUMENT_WRITE
	NEWLINE WS OWS SEMICOLON EQUAL INTEGER PLUS MINUS MULT DIVIDE 
	STRING LPAREN RPAREN COMMA

%%

file:
	SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END
	| SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
;

program:
	program line NEWLINE
	| /* Empty */
;

line:
	unterminated_line
	| unterminated_line SEMICOLON
	| unterminated_line SEMICOLON line
;

unterminated_line:
	declaration
	| assignment
	| definition
	| DOCUMENT_WRITE LPAREN parameter_list RPAREN
;

declaration:
	VARDEF IDENTIFIER
;

assignment:
	IDENTIFIER EQUAL expression
;

definition:
	VARDEF IDENTIFIER EQUAL expression
;

parameter_list:
	expression
	| expression COMMA parameter_list 
;

expression:
	value
	| parenthesized_expression
	| parenthesized_expression operator expression
	| value operator expression
;

parenthesized_expression:
	LPAREN expression RPAREN
;

operator:
	PLUS
	| MINUS
	| MULT
	| DIVIDE
;

value:
	INTEGER
	| IDENTIFIER
	| STRING
;

%%

FILE *yyin;
int yylineno;

int my_debug = 1;

yyerror(char *s)
{
    fprintf(stderr, "error: %s, line: %d\n", s, yylineno);
}

int main(int argc, char *argv[])
{
	//yydebug = 1;
	if (argc == 2 || argc == 3) {
		if (argc == 3 && strcmp(argv[2], "--debug") == 0) {
			my_debug = 1;
		}
		FILE *file;
		file = fopen(argv[1], "r");
		if (!file) {
            		fprintf(stderr, "could not open %s\n", argv[1]);
        	} else{
            		yyin = file;
            		//yyparse() will call yylex()
            		yyparse();
        	}
    	
	} else {
        	fprintf(stderr, "format: ./yacc_example [filename]");
    	}

    	return 0;
}


