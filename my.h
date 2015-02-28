
#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define my_debug 0
#define SYMBOL_TABLE_SIZE 1000
#define MAX_OBJECT_KEYS 20

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

#define YYSTYPE struct entity

/**	Enum for the types of entity values this language supports */
typedef enum {
	TYPE_INTEGER,
	TYPE_RUNE,
	TYPE_STRING,
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

/** The master definition for the structs our interpreter passes and stores
		The lexer converts tokens into entities during the lexing process before
		they undergo semantic analysis.

		During lexing, name will be null because no entities should have names.
		If the name is NOT null then that means the entity is a variable symbol.
		In either case, the type of the entity is stored in type.
		If its a variable, then that type is also the type of the variable. Obviously.

		Objects are strange.
		The root object is stored as type=TYPE_OBJECT, the name of the object, and no value.
		Each key of the object is stored as its own type, with name = "parentobject.keyname".
		This is possible becasue (A) we never need a list of all the keys of an object, and (B)
		because periods are otherwise NOT allowed in variable names.
*/
struct entity {

	char* name;

	E_TYPE type;
	union {
		int number;
		char rune;
		char* string;
		E_RESERVED reserved;
	} value;

};
