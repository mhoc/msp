
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

  buffer := new(bytes.Buffer)

  for {
    c := x.peek()

    // If we read whitespace as the first thing, just ignore it
    isWs, delim := isWhiteSpace(c)
    if buffer.Len() == 0 && isWs {
      x.pop()
      continue
    }

    // If the first actual character we read is a "unique symbol" then
    // we just return that symbol as what was lexed
    isDelim, delim := isUniqueSymbol(c)
    if buffer.Len() == 0 && isDelim {
      x.pop()
      return delim
    }

    // If the character we popped is a delimiter of some kind (whitespace or unique)
    // and the buffer has something in it, then we return what was inside the buffer
    // without popping the next character
    if isWs || isDelim {
      found, value := matchBuffer(buffer, yylval)
      if found {
        return value
      } else {
        return -1
      }
    }

    // Now we can write the byte to the buffer
    buffer.WriteByte(x.pop())

  }
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

func matchBuffer(buffer bytes.Buffer, yylval *yySymType) (bool, int) {

  s := buffer.String()
  found, value := isReservedWord(s, yylval)
  if found {
    return value
  }

  found, value := isIdentifier(s, yylval)
  if found {
    return value
  }

}

func isIdentifier(s string, yylval *yySymType) (bool, int) {

}

func isReservedWord(s string, yllval *yySymType) (bool, int) {
  switch s {
    case "var":
      return true, VARDEF
    case "<script type=\"text/JavaScript\">":
      return true, SCRIPT_TAG_START
    case "</script>":
      return true, SCRIPT_TAG_END
    case "document.write":
      return true, DOCUMENT_WRITE
  }
  return false, -1
}

func isWhiteSpace(c byte) (bool, int) {
  switch c {
    case ' ':
    case '\t':
      return true, WHITESPACE
    case '\n':
      return true, NEWLINE
  }
  return false, -1
}

func isUniqueSymbol(c byte) (bool, int) {
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
