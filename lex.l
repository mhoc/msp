%{
#include"y.tab.h"
#include <stdio.h>
#include "my.h"
%}

%option yylineno

%%
"<script type=\"text/JavaScript\">" {
	if (my_debug) { printf("SCRIPT_TAG_START "); }
	return SCRIPT_TAG_START;
}
"</script>" {
	if (my_debug) { printf("SCRIPT_TAG_END "); }
	return SCRIPT_TAG_END;
}
"var" {
	if (my_debug) { printf("VARDEF "); }
	return VARDEF;
}
[A-Za-z]+ {
	if (my_debug) { printf("IDENTIFIER "); }
	return IDENTIFIER;
}
[ \t]+ {
	if (my_debug) { printf("WHITESPACE "); }
	return WHITESPACE;
}
";" {
	if (my_debug) { printf("SEMICOLON "); }
	return SEMICOLON;
}
"\n" {
	if (my_debug) { printf("NEWLINE\n"); }
	return NEWLINE;
}
%%

