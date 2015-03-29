
GOPATH := ${PWD}
COLOR := "\033[0;32m"
WHITE := "\033[0;00m"

parser: main.go miniscript.nn.go y.go
	@echo $(COLOR) Overriding GOPATH to $(GOPATH) $(WHITE)
	@echo $(COLOR) Creating temp GOPATH fs structure to support multipackage compilation $(WHITE)
	mkdir -p src/mhoc.co/msp
	cp *.go src/mhoc.co/msp
	cp -r ast src/mhoc.co/msp
	cp -r log src/mhoc.co/msp
	cp -r symbol src/mhoc.co/msp
	@echo $(COLOR) Building parser binary $(WHITE)
	go build -o parser mhoc.co/msp
	@$(MAKE) uclean

y.go: yacc.y
	@echo $(COLOR) Compiling yacc grammar $(WHITE)
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo $(COLOR) Compiling lexical analyzer $(WHITE)
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo $(COLOR) Compiling nex lexical analyzer tool $(WHITE)
	cd nex && go build -o nexb nex.go
	mv nex/nexb .

uclean:
	@echo $(COLOR) Cleaning yacc intermediate files $(WHITE)
	rm -f y.output y.go
	@echo $(COLOR) Cleaning nex intermediate files $(WHITE)
	rm -f nexb miniscript.nn.go
	@echo $(COLOR) Cleaning GOPATH temp structure $(WHITE)
	rm -rf src bin cs352-integration-test

clean: uclean
	@echo $(COLOR) Deleting parser binary $(WHITE)
	rm -f parser cs352-integration-test

test: parser
	@echo $(COLOR) Downloading test cases $(WHITE)
	rm -rf cs352-integration-test
	git clone http://github.com/mhoc/cs352-integration-test.git
	@echo $(COLOR) Running test cases $(WHITE)
	@cd cs352-integration-test && go run main.go --exit-on-fail ../parser
	@$(MAKE) clean
