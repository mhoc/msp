
package main

import (
  "os"
  "mhoc.co/msp/log"
)

func main() {
  log.Traceln("Opening input file")

  // Parse command line arguments
  var file *os.File;
  var err error;
  if len(os.Args) == 1 {
    file = os.Stdin
  } else {
    file, err = os.Open(os.Args[1])
  }

  if err != nil {
    panic("File name provided does not exist")
  }

  log.Traceln("Beginning lex")
  yyParse(NewLexer(file))

}
