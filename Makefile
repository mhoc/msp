yacc_example:y.tab.c lex.yy.c
	gcc y.tab.c lex.yy.c -o yacc_example -lfl
y.tab.c: yacc_example.y
	bison -y -d -g -t --debug yacc_example.y
lex.yy.c: yacc_example_lex.l
	lex yacc_example_lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output

