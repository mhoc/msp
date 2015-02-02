%{
#include<stdio.h>
%}

%token ASSIGN ID NUM PLUS
%start exp 

%%
exp     : exp PLUS NUM
        | ID
        ;
%%

FILE *yyin;
int yylineno;
yyerror(char *s)
{
    fprintf(stderr, "error: %s, line: %d\n", s, yylineno);
}

main(int argc, char *argv[])
{
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
    } else{
        fprintf(stderr, "format: ./yacc_example [filename]");
    }
}


