package main

import (
	"flag"
	"os"
	"mhoc.co/msp/ast"
	"mhoc.co/msp/log"
)

func main() {
	log.Trace("-m-", "Opening input file")

	// Parse debug flags
	flag.BoolVar(&log.LOG_TOKENS, "log-tokens", false, "Enable list of parsed tokens")
	flag.BoolVar(&log.LOG_TRACE, "log-trace", false, "Enable trace logging of debug output")
	flag.BoolVar(&log.EXTENSIONS, "extensions", false, "Enables parser extensions for additional features. See README for a list of these.")
	flag.Parse()

	// Init builtin functions
	ast.InitBuiltins()

	// Parse command line arguments
	var file *os.File
	var err error
	if flag.NArg() == 0 {
		log.Trace("-m-", "Reading from stdin")
		file = os.Stdin
	} else if flag.NArg() == 1 {
		log.Trace("-m-", "Reading from file "+flag.Arg(0))
		file, err = os.Open(flag.Arg(0))
	} else {
		panic("Must provide filename to read from or no filename at all")
	}

	if err != nil {
		panic("File name provided does not exist")
	}

	log.Trace("-m-", "Beginning lex")
	yyParse(NewLexer(file))

}
