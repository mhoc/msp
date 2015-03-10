
package token

type Type int
const (
  T_INTEGER Type = iota
  T_STRING Type = iota
  T_OBJECT Type = iota
  T_RUNE Type = iota
  T_FIELD Type = iota
  T_UNDEFINED Type = iota
  T_RESERVED Type = iota
)
