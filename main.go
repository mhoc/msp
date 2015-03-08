
package main

import (
  "os"
)

func main() {
  LogTrace("Opening input file")

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

  LogTrace("Beginning lex")
  yyParse(NewLexer(file))

}
