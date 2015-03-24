
package main

import (
  "flag"
  "os"
  "mhoc.co/msp/log"
)

func main() {
  log.Trace("Opening input file")

  // Parse debug flags
  flag.BoolVar(&log.LOG_TOKENS, "log-tokens", false, "Enable list of parsed tokens")
  flag.BoolVar(&log.LOG_TRACE, "log-trace", false, "Enable trace logging of debug output")
  flag.BoolVar(&log.LOG_AST, "log-ast", false, "Enable output of the ast after it is parsed")
  flag.Parse()

  // Parse command line arguments
  var file *os.File;
  var err error;
  if flag.NArg() == 0 {
    log.Trace("Reading from stdin")
    file = os.Stdin
  } else if flag.NArg() == 1 {
    log.Trace("Reading from file " + flag.Arg(0))
    file, err = os.Open(flag.Arg(0))
  } else {
    panic("Must provide filename to read from or no filename at all")
  }

  if err != nil {
    panic("File name provided does not exist")
  }

  log.Trace("Beginning lex")
  yyParse(NewLexer(file))

}
