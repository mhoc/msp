
yacc_example: y.tab.c lex.yy.c
	gcc y.tab.c lex.yy.c -o yacc_example -lfl

y.tab.c: yacc.y
	bison -y -d -g -t --debug yacc.y

lex.yy.c: lex.l
	lex lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output

