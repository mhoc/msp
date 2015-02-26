
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
	OPENBRACE
	CLOSEBRACE
	COLON

%%

file:
	SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END
	| SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
	| NEWLINE SCRIPT_TAG_START NEWLINE program SCRIPT_TAG_END NEWLINE
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
	VARDEF IDENTIFIER {
		declare_symbol($2.value.string);
	}
;

assignment:
	IDENTIFIER EQUAL expression {
		if ($3.type == TYPE_INTEGER) {
			update_value_i($1.value.string, $3.value.number);
		}
		else if ($3.type == TYPE_STRING) {
			update_value_s($1.value.string, $3.value.string);
		}
		else if ($3.type == TYPE_OBJECT) {

		}
	}
;

definition:
	VARDEF IDENTIFIER EQUAL expression {
		if ($4.type == TYPE_INTEGER) {
			insert_symbol_i($2.value.string, $4.value.number);
		}
		else if ($4.type == TYPE_STRING) {
			insert_symbol_s($2.value.string, $4.value.string);
		}
		else if ($4.type == TYPE_OBJECT) {

		}
	}
;

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
		}
	}
	| /* empty */
;

expression:
	constant
	| variable
	| parenthesized_expression
	| parenthesized_expression operator expression
	| constant operator expression {
		if ($1.type == TYPE_INTEGER && $3.type == TYPE_INTEGER) {
			if ($2.value.rune == '+') {
				$1.value.number += $3.value.number;
				$$ = $1;
			}
		}
	}
	| variable operator expression
	| OPENBRACE field_list CLOSEBRACE
	| OPENBRACE NEWLINE field_list CLOSEBRACE
;

parenthesized_expression:
	LPAREN expression RPAREN {
		$$ = $2;
	}
;

operator:
	PLUS {
		debugf1("(%c) ", $1.value.rune)
	}
	| MINUS {
		debugf1("(%c) ", $1.value.rune)
	}
	| MULT {
		debugf1("(%c) ", $1.value.rune)
	}
	| DIVIDE {
		debugf1("(%c) ", $1.value.rune)
	}
;

constant:
	INTEGER {
		debugf1("(%i) ", $1.value.number)
	}
	| STRING {
		debugf1("(%s) ", $1.value.string)
	}
;

variable:
	IDENTIFIER {
		debugf1("(%s) ", $1.value.string)

		struct symbol* s = get_symbol($1.value.string);

		if (s == NULL) {
			printf("Error line %d: Use of undeclared variable %s\n", yylineno, $1.value.string);
		}
		else if (s->type == TYPE_INTEGER) {
			$$.type = TYPE_INTEGER;
			$$.value.number = s->value.number;

		} else if (s->type == TYPE_STRING) {
			$$.type = TYPE_STRING;
			$$.value.string = s->value.string;
		}

	}
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

	if (argc == 2 || argc == 3) {
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
