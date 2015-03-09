
parser: main.go miniscript.nn.go y.go
	@echo "=== Compiling parser ===="
	go build -o parser main.go y.go miniscript.nn.go
	@$(MAKE) uclean

y.go: yacc.y
	@echo "==== Compiling yacc grammar ===="
	go tool yacc yacc.y

miniscript.nn.go: miniscript.nex nexb
	@echo "==== Creating lexical analyzer ===="
	./nexb miniscript.nex

nexb: nex/nex.go
	@echo "==== Compiling nex lexical analyzer tool ===="
	cd nex && go build -o nexb nex.go
	@mv nex/nexb .

uclean:
	@echo "==== Cleaning yacc intermediate files ===="
	rm -f y.output y.go
	@echo "==== Cleaning nex intermediate files ===="
	rm -f nexb miniscript.nn.go

clean: uclean
	@echo "==== Deleting parser binary ===="
	rm -f main

testsm:
	@$(MAKE) parser
	@echo "==== Runing small test, output below ===="
	@./parser < testsmall
	@echo "=========="
	@$(MAKE) clean
