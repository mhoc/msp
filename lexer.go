
/** Custom lexer implementation for the miniscript language
    The core of this lexer is built around golang's yacc tool, but I use
    a very custom way of parsing tokens because regex is slow and
    this gives very find-grained control over how the lexing happens
*/

package main

import (
  "bytes"
  "fmt"
  "log"
)

const EOF = 0

type yyLex struct {
  line []byte
}

func (x *yyLex) Lex(yylval *yySymType) int {
  c := x.pop()
  var found bool
  var val int

  // Lex all of the unique single characters
  found, val = lexUniqueChar(c)
  if found {
    return val
  }

  // Build a buffer of
  found, val = lexReservedWord(x.line)

}

func (x *yyLex) Error(msg string) {
  fmt.Println("syntax error")
}

func (x *yyLex) pop() byte {
	if len(x.line) == 0 {
		return EOF
	}
	c := x.line[0]
	x.line = x.line[1:]
	return c
}

func (x *yyLex) peek() byte {
  if len(x.line) == 0 {
    return EOF
  }
  return x.line[0]
}

func lexUniqueChar(c byte) (bool, int) {
  switch c {
    case EOF:
      return true, EOF
    case '{':
      return true, LBRACE
    case '}':
      return true, RBRACE
    case ',':
      return true, COMMA
    case '=':
      return true, EQUAL
    case '+':
      return true, PLUS
    case '-':
      return true, MINUS
    case '*':
      return true, MULT
    case '/':
      return true, DIVIDE
    case ':':
      return true, COLON
    case ';':
      return true, SEMICOLON
    case '(':
      return true, LPAREN
    case ')':
      return true, RPAREN
  }
  return false, -1
}

func lexReservedWords(c byte) (bool, int) {

  // We maintain a byte buffer of what we've seen going forward
  add := func(b *bytes.Buffer, b byte) {
		if _, err := b.WriteByte(b); err != nil {
			log.Fatalf("WriteByte: %s", err)
		}
	}


	var b bytes.Buffer
	add(&b, c)

}
