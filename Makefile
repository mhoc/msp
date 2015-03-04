
parser: y.go
	@echo "-> Compiling parser"

y.go: yacc.y
	@echo "-> Compiling yacc grammar"
	go tool yacc yacc.y

clean:
	@echo "-> Cleaning yacc output files"
	rm -f y.output y.go

testgh: parser
	wget bit.ly/1zMeiCA -O run.sh && chmod +x run.sh && ./run.sh

testlocal: parser
	@cp ~/src/cs352-test-cases/test.py .
	@cp ~/src/cs352-test-cases/cases.py .
	python test.py
	@$(MAKE) clean
