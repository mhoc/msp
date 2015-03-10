
GOPATH := ${PWD}

parser: main.go miniscript.nn.go y.go
	@echo "\033[1;36mOverriding GOPATH to" $(GOPATH) " \033[0;00m"
	@echo "\033[1;36mCreating temp GOPATH fs structure to support multipackage compilation\033[0;00m"
	@printf "> "
	mkdir -p src/mhoc.co/msp
	@printf "> "
	cp *.go src/mhoc.co/msp
	@printf "> "
	cp -r token src/mhoc.co/msp
	cp -r util src/mhoc.co/msp
	@echo "\033[1;36mBuilding parser binary\033[0;00m"
	@printf "> "
	go build -o parser mhoc.co/msp
	@$(MAKE) uclean

y.go: yacc.y
	@echo "\033[1;36mCompiling yacc grammar\033[0;00m"
	@printf "> "
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo "\033[1;36mCompiling lexical analyzer\033[0;00m"
	@printf "> "
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo "\033[1;36mCompiling nex lexical analyzer tool\033[0;00m"
	@printf "> "
	cd nex && go build -o nexb nex.go
	@printf "> "
	mv nex/nexb .

uclean:
	@echo "\033[1;36mCleaning yacc intermediate files\033[0;00m"
	@printf "> "
	rm -f y.output y.go
	@echo "\033[1;36mCleaning nex intermediate files\033[0;00m"
	@printf "> "
	rm -f nexb miniscript.nn.go
	@echo "\033[1;36mCleaning goroot intermediate directories\033[0;00m"
	@printf "> "
	rm -rf src

clean: uclean
	@echo "\033[1;36mDeleting parser binary\033[0;00m"
	@printf "> "
	rm -f main
