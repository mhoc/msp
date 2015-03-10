
GOPATH := ${PWD}

parser: main.go miniscript.nn.go y.go
	@echo "Overriding GOPATH to" $(GOPATH)
	@echo "Creating temp GOPATH fs structure to support multipackage compilation"
	@printf "> "
	mkdir -p src/mhoc.co/msp
	@printf "> "
	cp *.go src/mhoc.co/msp
	@printf "> "
	cp -r log src/mhoc.co/msp
	@echo "Building parser binary"
	@printf "> "
	go build -o parser mhoc.co/msp
	@$(MAKE) uclean

y.go: yacc.y
	@echo "Compiling yacc grammar"
	@printf "> "
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo "Compiling lexical analyzer"
	@printf "> "
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo "Compiling nex lexical analyzer tool"
	@printf "> "
	cd nex && go build -o nexb nex.go
	@printf "> "
	mv nex/nexb .

uclean:
	@echo "Cleaning yacc intermediate files"
	@printf "> "
	rm -f y.output y.go
	@echo "Cleaning nex intermediate files"
	@printf "> "
	rm -f nexb miniscript.nn.go
	@echo "Cleaning goroot intermediate directories"
	@printf "> "
	rm -rf src

clean: uclean
	@echo "Deleting parser binary"
	@printf "> "
	rm -f main
