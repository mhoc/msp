
package util

import (
  "fmt"
  "os"
)

func CompilerError(s string) {
  fmt.Fprintf(os.Stderr, s)
  os.Exit(0)
}

func TypeViolation(line int) {
  
}

func InternalError(s string) {
  panic(s)
}
