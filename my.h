
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

#define YYSTYPE Token

/**	Enum for the types of entity values this language supports */
typedef enum {
	TYPE_INTEGER,
	TYPE_RUNE,
	TYPE_STRING,
	TYPE_FIELD,
	TYPE_FIELDLIST,
	TYPE_UNDEFINED,
	TYPE_RESERVED
} Type;

/** Enum for each reserved keyword. We have 4 right now. */
typedef enum {
	RESERVE_START_TAG,
	RESERVE_END_TAG,
	RESERVE_F_DOC_WRITE,
	RESERVE_VAR
} E_RESERVED;

typedef struct s Symbol;

typedef struct {
	int size;
	int maxSize;
	Symbol** list;
} FieldList;

typedef struct {

	Type type;
	union {
		int number;
		char rune;
		char* string;
		Symbol* field;
		FieldList* fieldList;
		E_RESERVED reserved;
	} value;

} Token;

struct s {

	char* name;
	Type type;
	union {
		int number;
		char* string;
		FieldList* fieldList;
	} value;

};
