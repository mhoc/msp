%{

#include "my.h"
#include "y.tab.h"

#define val_rune				yylval.value.rune = yytext[0]; yylval.type = TYPE_RUNE;
#define val_int					yylval.value.number = atoi(yytext); yylval.type = TYPE_INTEGER;
#define val_id					yylval.value.string = strdup(yytext); yylval.type = TYPE_STRING;
#define val_res(x)			yylval.value.reserved = x; yylval.type = TYPE_RESERVED;

%}

%option yylineno

%%

\n+ {
	/** This rule condenses any occurrence of multiple newlines into a single newline */
	debug("NEWLINE+\n")
	val_rune
	return NEWLINE;
}

"<script type=\"text/JavaScript\">" {
	debug("SCRIPT_TAG_START ")
	val_res(RESERVE_START_TAG)
	return SCRIPT_TAG_START;
}

"</script>" {
	debug("SCRIPT_TAG_END ")
	val_res(RESERVE_END_TAG)
	return SCRIPT_TAG_END;
}

"document\.write" {
	debug("DOCUMENT_WRITE ")
	val_res(RESERVE_F_DOC_WRITE)
	return DOCUMENT_WRITE;
}

\"[^\"\n]*\" {
	debug("STRING ")
	char* nstr = malloc(strlen(yytext));
	memcpy(nstr, yytext + sizeof(char), strlen(yytext)-2);
	nstr[strlen(yytext)-2] = '\0';
	yylval.value.string = nstr;
	yylval.type = TYPE_STRING;
	return STRING;
}

"var"[ \t] {
	/* the var keyword requires a space at the end. this is taken care of here */
	debug("VARDEF ")
	val_res(RESERVE_VAR)
	return VARDEF;
}

[ \t]+ {
	/* We've already taken care of quoted strings so now remove all other whitespaces */
}

[0-9]+ {
	debug("INTEGER ")
	val_int
	return INTEGER;
}

[A-Za-z][A-Za-z0-9_]* {
	debug("IDENTIFIER ")
	val_id
	return IDENTIFIER;
}

[A-Za-z][A-Za-z0-9]*\.[A-Za-z][A-Za-z0-9]* {
	debug("OBJKEY ");
	val_id
	return OBJKEY;
}

"{" {
	debug("LBRACE ")
	val_rune
	return LBRACE;
}

"}" {
	debug("RBRACE ")
	val_rune
	return RBRACE;
}

"," {
	debug("COMMA ")
	val_rune
	return COMMA;
}

"=" {
	debug("EQUAL ")
	val_rune
	return EQUAL;
}

"+" {
	debug("PLUS ")
	val_rune
	return PLUS;
}

"-" {
	debug("MINUS ")
	val_rune
	return MINUS;
}

"*" {
	debug("MULT ")
	val_rune
	return MULT;
}

"/" {
	debug("DIVIDE ")
	val_rune
	return DIVIDE;
}

";" {
	debug("SEMICOLON ")
	val_rune
	return SEMICOLON;
}

":" {
	debug("COLON ")
	val_rune
	return COLON;
}

"(" {
	debug("LPAREN ")
	val_rune
	return LPAREN;
}

")" {
	debug("RPAREN ")
	val_rune
	return RPAREN;
}

. {
	printf("Uncaught token %s. This is probably a syntax error.", yytext);
}

%%
