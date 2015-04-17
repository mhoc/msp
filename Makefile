
GOPATH := ${PWD}

parser: main.go miniscript.nn.go y.go
	@echo `tput smul``tput setaf 2` Overriding GOPATH to $(GOPATH) `tput sgr0`
	@echo `tput smul``tput setaf 2` Creating temp GOPATH fs structure to support multipackage compilation `tput sgr0`
	mkdir -p src/mhoc.co/msp
	cp *.go src/mhoc.co/msp
	cp -r ast src/mhoc.co/msp
	cp -r log src/mhoc.co/msp
	@echo `tput smul``tput setaf 2` Building parser binary `tput sgr0`
	go build -o parser mhoc.co/msp
	@$(MAKE) uclean

y.go: yacc.y
	@echo `tput smul``tput setaf 2` Compiling yacc grammar `tput sgr0`
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo `tput smul``tput setaf 2` Compiling lexical analyzer `tput sgr0`
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo `tput smul``tput setaf 2` Compiling nex lexical analyzer tool `tput sgr0`
	cd nex && go build -o nexb nex.go
	mv nex/nexb .

uclean:
	@echo `tput smul``tput setaf 2` Cleaning yacc intermediate files `tput sgr0`
	rm -f y.output y.go
	@echo `tput smul``tput setaf 2` Cleaning nex intermediate files `tput sgr0`
	rm -f nexb miniscript.nn.go
	@echo `tput smul``tput setaf 2` Cleaning GOPATH temp structure `tput sgr0`
	rm -rf src bin

clean: uclean
	@echo `tput smul``tput setaf 2` Deleting parser binary `tput sgr0`
	rm -f parser

test: parser
	@echo `tput smul``tput setaf 2` Running test cases `tput sgr0`
	@echo
	@cd test && go run main.go ../parser
	@echo
	@$(MAKE) clean
