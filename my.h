
#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define my_debug 0
#define SYMBOL_TABLE_SIZE 1000

/* If macros are used here to save on a comparison at run time
 * for a global debug variable */

#if my_debug == 0
#define debug(x) ;
#define debugf1(x,y) ;
#endif

#if my_debug != 0
#define debug(x) printf(x);
#define debugf1(x,y) printf(x,y);
#endif

#define YYSTYPE struct tok

/** Enum which stores the various types we recognize
 *	These enums are used for both the type of tokens AND the type of variables
 *  Some make sense in context while others don't. Its just a shortcut. */
typedef enum {
	TYPE_INTEGER,
	TYPE_RUNE,
	TYPE_STRING,
	TYPE_IDENTIFIER,
	TYPE_OBJECT,
	TYPE_UNDEFINED,
	TYPE_RESERVED
} E_TYPE;

/** Enum for each reserved keyword. We have 4 right now. */
typedef enum {
	RESERVE_START_TAG,
	RESERVE_END_TAG,
	RESERVE_F_DOC_WRITE,
	RESERVE_VAR
} E_RESERVED;

/** Definition for a token which is returned by the lexer
 		Each token has a value and type */
struct tok {

	union {
		int number;
		char rune;
		char* string;
		E_RESERVED reserved;
	} value;

	E_TYPE type;

};

/** Definition for a symbol which is stored in the symbol table
		Each symbol has a name, type, and value.

		If the type is undefined, value will be null and type is TYPE_UNDEFINED.

		If the type is an object, we do some bs.
		An entry for the object itself is stored under type=TYPE_OBJECT and no value.
		Each child of the object is stored as its own symbol with type of whatever it is and
		a name of "parentname.key".
		We don't maintain a reference to which children exist (yet) because there's no need.
*/
struct symbol {

	char* name;
	E_TYPE type;

	union {
		int number;
		char* string;
	} value;

};
