
parser: y.go
	@echo "--> Compiling parser"
	go build main.go lexer.go y.go
	@$(MAKE) pclean

y.go: yacc.y
	@echo "--> Compiling yacc grammar"
	go tool yacc yacc.y

lex.nn.go: lex.nex
	@echo "--> Runing lexical analyzer"
	cd nex
	go build nex.go
	nex -r -s ../miniscript.nex
	cp lex.nn.go

pclean:
	@echo "--> Cleaning yacc intermediate files"
	rm -f y.output y.go y.output

clean:
	@$(MAKE) pclean
	@echo "--> Cleaning binaries"
	rm -f main

testgh: parser
	wget bit.ly/1zMeiCA -O run.sh && chmod +x run.sh && ./run.sh

testlocal: parser
	@cp ~/src/cs352-test-cases/test.py .
	@cp ~/src/cs352-test-cases/cases.py .
	python test.py
	@$(MAKE) clean
