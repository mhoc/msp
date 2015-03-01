
// #include "my.h" done in including file

/** Symbol tables for storing variables */
Symbol* symbolTable[SYMBOL_TABLE_SIZE];

/********************
  * Utility methods *
  ********************/

char* strcatc(char* c1, char* c2) {
  char* ns = malloc(strlen(c1) + strlen(c2));
  strcpy(ns, c1);
  strcat(ns, c2);
  return ns;
}

/** FNV-1a hashing algorithm */
int hash(char* string) {
  int fnv_prime = 16777619;
  int fnv_offset_basis = 2166136261;

  int hash = fnv_offset_basis;
  for (int i = 0; i < strlen(string); i++) {
    char c = string[i];
    hash = hash ^ c;
    hash = hash * fnv_prime;
  }

  hash = hash % SYMBOL_TABLE_SIZE;
  if (hash >= 0) {
    return hash;
  } else {
    return hash - (2 * hash);
  }

}

/****************************
  * Symbol table management *
  ****************************/

void declareSymbol(char* name) {

  // Hash the name and get the current variable stored
  int hashIndex = hash(name);
  Symbol* declared = symbolTable[hashIndex];

  // If we aren't storing one, create it here
  if (declared == NULL) {
    declared = malloc(sizeof(Symbol));
    declared->name = strdup(name);
  }
  else if (strcmp(name, declared->name)) {
    printf("~~ Hash Collision ~~");
  }
  else {
    // Redeclaration
    // TODO FREE OLD RESOURCES
  }

  // Set the symbol's value as undefined
  declared->type = TYPE_UNDEFINED;

  // Declare it in the table
  symbolTable[hashIndex] = declared;

}

/** Generic symbol definition dispatch */
void defineSymbol(char* name, Token* token, int displayError) {

  // Hash and get the name
  int hashIndex = hash(name);
  Symbol* declared = symbolTable[hashIndex];

  // Throw an error if the symbol isnt defined
  if (declared == NULL) {
    if (token->type != TYPE_FIELD && token->type != TYPE_FIELDLIST && displayError) {
      printf("Error (line %d): Attempting to define the value of a variable which has not been declared\n", yylineno);
    }
    declared = malloc(sizeof(Symbol));
    declared->name = strdup(name);
  }

  // Free old resources
  // TODO
  // fuck it we'll use c++ for the next part lol
  // (said no one in the real world ever)

  // Assign a new type and value
  declared->type = token->type;
  switch (declared->type) {
    case TYPE_INTEGER:
      declared->value.number = token->value.number; break;
    case TYPE_STRING:
      declared->value.string = token->value.string; break;
    case TYPE_FIELDLIST:
      declared->value.fieldList = token->value.fieldList; break;
    case TYPE_FIELD:
      declared->type = token->value.field->type;
      if (token->value.field->type == TYPE_INTEGER) {
        declared->value.number = token->value.field->value.number;
      }
      else if (token->value.field->type == TYPE_STRING) {
        declared->value.string = token->value.field->value.string;
      }
    //case TYPE_UNDEFINED:
      // TODO
  }

  // If this is a field list, recall defineSymbol on every item in the field list
  if (declared->type == TYPE_FIELDLIST) {
    for (int i = 0; i < declared->value.fieldList->size; i++) {
      // holy jesus christ the spaghetti
      char* tname = strcatc(name, strcatc(".", declared->value.fieldList->list[i]->name));

      // its coming out of my pockets
      int hashindex = hash(tname);
      // who wrote this stuff?
      Symbol* field = malloc(sizeof(Symbol));
      field->name = tname;
      field->type = declared->value.fieldList->list[i]->type;


      switch (field->type) {
        case TYPE_INTEGER:
          field->value.number = declared->value.fieldList->list[i]->value.number; break;
        case TYPE_STRING:
          field->value.number = declared->value.fieldList->list[i]->value.string; break;
      }

      symbolTable[hashindex] = field;

    }
  }

  symbolTable[hashIndex] = declared;

}

Token* getSymbol(char* name) {
  Symbol* s = symbolTable[hash(name)];

  if (s == NULL) {
    printf("Error (line %d): Use of undeclared variable %s\n", yylineno, name);
  }

  Token* t = malloc(sizeof(Token));
  t->type = s->type;

  switch (t->type) {
    case TYPE_INTEGER:
      t->value.number = s->value.number; break;
    case TYPE_STRING:
      t->value.string = s->value.string; break;
    case TYPE_FIELDLIST:
      t->value.fieldList = s->value.fieldList; break;
    //case TYPE_UNDEFINED:
      // Do nothing
  }

  return t;
}

/****************
  * Printing    *
  ****************/

void printExpression(Token* token) {
  if (token->type == TYPE_STRING) {
    if (strcmp(token->value.string, "<br />")) {
      printf("%s", token->value.string);
    } else {
      printf("\n");
    }
  }
  else if (token->type == TYPE_INTEGER) {
    printf("%i", token->value.number);
  }
  else if (token->type == TYPE_UNDEFINED) {
    printf("undefined");
  }
  else {
    printf("Error (line %d): Attempting to print a non-integer or string value\n");
  }
}

/****************
  * Expressions *
  ****************/

Token addTokens(Token t1, Token t2) {

  if (t1.type == TYPE_INTEGER && t2.type == TYPE_INTEGER) {
    t1.value.number += t2.value.number;
  }

  else if (t1.type == TYPE_STRING && t2.type == TYPE_STRING) {
    char* newstr = malloc(strlen(t1.value.string) + strlen(t2.value.string));
    strcpy(newstr, t1.value.string);
    strcat(newstr, t2.value.string);
    t1.value.string = newstr;
  }
  else if (t1.type == TYPE_UNDEFINED) {
    // Basically just return T1.
  }
  else {
    printf("Error (line %d): Attempting to apply addition to unsupported types\n", yylineno);
  }

  return t1;

}

Token subtractTokens(Token t1, Token t2) {

  if (t1.type == TYPE_INTEGER && t2.type == TYPE_INTEGER) {
    t1.value.number -= t2.value.number;
  }

  else {
    printf("Error (line %d): Attempting to apply subtraction to unsupported types\n", yylineno);
  }

  return t1;

}

Token multiplyTokens(Token t1, Token t2) {

  if (t1.type == TYPE_INTEGER && t2.type == TYPE_INTEGER) {
    t1.value.number *= t2.value.number;
  }

  else {
    printf("Error (line %d): Attempting to apply multiplication to unsupported types\n", yylineno);
  }

  return t1;

}

Token divideTokens(Token t1, Token t2) {

  if (t1.type == TYPE_INTEGER && t2.type == TYPE_INTEGER) {
    t1.value.number /= t2.value.number;
  }

  else {
    printf("Error (line %d): Attempting to apply division to unsupported types\n", yylineno);
  }

  return t1;

}

/************
  * OBJECTS *
  ************/

Token* createField(char* name, Token* value) {
  Token* t = malloc(sizeof(Token));
  t->type = TYPE_FIELD;

  Symbol* s = malloc(sizeof(Symbol));
  s->name = strdup(name);
  s->type = value->type;

  switch (s->type) {
    case TYPE_INTEGER:
      s->value.number = value->value.number; break;
    case TYPE_STRING:
      s->value.string = value->value.string; break;
  }

  t->value.field = s;
  return t;
}

Token* newFieldList() {
  Token* t = malloc(sizeof(Token));
  t->type = TYPE_FIELDLIST;
  t->value.fieldList = malloc(sizeof(FieldList));
  t->value.fieldList->size = 0;
  t->value.fieldList->maxSize = 4;
  t->value.fieldList->list = malloc(sizeof(Symbol*) * 4);
  return t;
}

void addToFieldList(FieldList* list, Token* t) {

  if (list->size >= list->maxSize) {
    printf("Too many!");
    return;
  }

  list->list[list->size++] = t->value.field;

}
