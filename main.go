
package main

import (
  "bufio"
  "io"
  "log"
  "os"
)

func main() {

  in := bufio.NewReader(os.Stdin)
  for {
    if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}
		yyParse(&yyLex{line: line})
  }

}
