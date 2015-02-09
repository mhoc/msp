%{
#include<stdio.h>
#include "my.h"
%}

%token 	SCRIPT_TAG_START SCRIPT_TAG_END 
	VARDEF IDENTIFIER 
	NEWLINE WHITESPACE SEMICOLON

%%
goal:
	file
;

file:
	SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END
	| SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
;

program:
	program line NEWLINE
	| /* Empty */
;

line:
	definition
	| definition SEMICOLON 
;

definition:
	VARDEF WHITESPACE IDENTIFIER
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
	if (my_debug) {
		printf("Advanced custom debugging output is enabled.\n");
	}

	//yydebug = 1;
	if (argc == 2) {
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


