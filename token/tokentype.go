
package token

type Type int
const (
  T_INTEGER TokenType = iota
  T_STRING TokenType = iota
  T_OBJECT TokenType = iota
  T_RUNE TokenType = iota
  T_FIELD TokenType = iota
  T_UNDEFINED TokenType = iota
  T_RESERVED TokenType = iota
)
