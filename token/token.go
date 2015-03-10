
package token

import (
  "mhoc.co/msp/util"
)

/** Tokens are lexemes which are lexed by nex.
    Each token has a type, and dependng on its type a value stored inside
    the struct.
*/

type Token struct {
  Vtype Type
  Intv int
  Strv string
  Runv rune
  Fldv util.Field
  Objv util.Object
  Rswv util.ReservedWord
}
