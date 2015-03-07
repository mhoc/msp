
package main

import (
  "log"
  "os"
)

func main() {

  // Parse command line arguments
  var file *os.File;
  var err error;
  if len(os.Args) == 1 {
    file = os.Stdin
  } else {
    file, err = os.Open(os.Args[1])
  }

  if err != nil {
    log.Panic("File name provided does not exist")
  }

  yyParse(NewLexer(file))

}
