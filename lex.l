%{
#include"y.tab.h"
#include <stdio.h>
#include "my.h"
%}

%option yylineno

%%
\n+ {
	/** This rule condenses any occurance of multiple newlines into a single newline */
	debug("NEWLINE+\n");
	return NEWLINE;
}
"<script type=\"text/JavaScript\">" {
	debug("SCRIPT_TAG_START ");
	return SCRIPT_TAG_START;
}
"</script>" {
	debug("SCRIPT_TAG_END ");
	return SCRIPT_TAG_END;
}
"document\.write" {
	debug("DOCUMENT_WRITE ");
	return DOCUMENT_WRITE;
}
\"[^\"]*\" {
	debug("STRING ");
	return STRING;
}
"var " {
	/* the var keyword requires a space at the end. this is taken care of here */
	debug("VARDEF ");
	return VARDEF;
}
[ \t]+ {
	/* We've already taken care of quoted strings so now remove all other whitespaces */
}
[0-9]+ {
	debug("INTEGER ");
	return INTEGER;
}
[A-Za-z]+ {
	debug("IDENTIFIER ");
	return IDENTIFIER;
}
"," {
	debug("COMMA, ");
	return COMMA;
}
"=" {
	debug("EQUAL ");
	return EQUAL;
}
"+" {
	debug("PLUS ");
	return PLUS;
}
"-" {
	debug("MINUS ");
	return MINUS;
}
"*" {
	debug("MULT ");
	return MULT;
}
"/" {
	debug("DIVIDE ");
	return DIVIDE;
}
";" {
	debug("SEMICOLON ");
	return SEMICOLON;
}
"(" {
	debug("LPAREN ");
	return LPAREN;
}
")" {
	debug("RPAREN ");
	return RPAREN;
}
%%

