
package main

import (
  //"bufio"
  //"io"
  "os"
)

func main() {

  yyParse(NewLexer(os.Stdin))

}
