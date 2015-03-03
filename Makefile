parser: y.tab.c lex.yy.c helper.c my.h
	gcc y.tab.c lex.yy.c -o parser -lfl -w

y.tab.c: yacc.y
	bison -y -d -g -t --verbose yacc.y

lex.yy.c: lex.l
	lex lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output parser y.vcg test.py cases.py cases.pyc
	rm -rf tests

testgh: parser
	wget bit.ly/1zMeiCA -O run.sh && chmod +x run.sh && ./run.sh

testlocal: parser
	@cp ~/src/cs352-test-cases/test.py .
	@cp ~/src/cs352-test-cases/cases.py .
	python test.py
	@$(MAKE) clean

