
package main

import (
  "os"
  "mhoc.co/msp/util"
)

func main() {
  util.LogTrace("Opening input file", 2)

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

  util.LogTrace("Beginning lex", 2)
  yyParse(NewLexer(file))

}
