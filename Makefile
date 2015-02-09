
parser.out: y.tab.c lex.yy.c
	gcc y.tab.c lex.yy.c -o parser.out -lfl

y.tab.c: yacc.y
	bison -y -d -g -t --verbose yacc.y

lex.yy.c: lex.l
	lex lex.l

test: parser.out
	./parser.out tests/1

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output

