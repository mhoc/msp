
parser.out: y.tab.c lex.yy.c
	gcc y.tab.c lex.yy.c -o parser.out -lfl

y.tab.c: yacc.y
	bison -y -d -g -t --verbose yacc.y

lex.yy.c: lex.l
	lex lex.l

clean:
	rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output

test: test1 test2 test3 test31 test32 test4 test5 test6 testgh
	@echo ""
	@echo ""

test1: parser.out
	@echo ""
	@echo "==> Testing empty file (empty1)"
	./parser.out tests/empty1

test2: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing empty file with multiple newlines (empty2)"
	./parser.out tests/empty2

test3: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing simple assignment (simple_assignment)"
	./parser.out tests/simple_assignment

test31: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing assignment with odd whitespaces (whitespace_assignment)"
	./parser.out tests/whitespace_assignment

test32: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing assignment with odd newlines (newline_assignment)"
	./parser.out tests/newline_assignment

test4: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing complex assignment (complex_assignment)"
	./parser.out tests/complex_assignment

test5: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing simple declaration (simple_declaration)"
	./parser.out tests/simple_declaration

test6: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing parenthesis with assignment (parens)"
	./parser.out tests/parens

testgh: parser.out
	@echo ""
	@echo ""
	@echo "==> Testing handout code (handout)"
	./parser.out tests/handout

