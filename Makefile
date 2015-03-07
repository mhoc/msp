
parser: main.go miniscript.nn.go y.go logger.go
	@echo "--> Compiling parser"
	go build main.go y.go miniscript.nn.go logger.go
	@$(MAKE) uclean

y.go: yacc.y
	@echo "--> Compiling yacc grammar"
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo "--> Creating lexical analyzer"
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo "--> Compiling nex lexical analyzer tool"
	cd nex && go build -o nexb nex.go
	@mv nex/nexb .

uclean:
	@echo "--> Cleaning yacc intermediate files"
	rm -f y.output y.go
	@echo "--> Cleaning nex intermediate files"
	rm -f nexb miniscript.nn.go

clean:
	@$(MAKE) uclean
	@echo "--> Cleaning binaries"
	rm -f main

testgh: parser
	wget bit.ly/1zMeiCA -O run.sh && chmod +x run.sh && ./run.sh

testlocal: parser
	@cp ~/src/cs352-test-cases/test.py .
	@cp ~/src/cs352-test-cases/cases.py .
	python test.py
	@$(MAKE) clean
