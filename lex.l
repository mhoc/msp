%{
#include"y.tab.h"
%}

%option yylineno

/* regular definitions */
delim   [ \t]
ws      {delim}+
letter  [A-Za-z]
digit   [0-9]
id      {letter}({letter}|{digit})*
num     {digit}+

%%
{id}                                    {return(ID);}
{num}                                   {return(NUM);}
"+"                                     {return(PLUS);}
{ws}                                    {}
%%

