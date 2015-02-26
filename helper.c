
// #include "my.h" done in including file

/** Symbol tables for storing variables */
struct entity* symbol_table[SYMBOL_TABLE_SIZE];

/** Alocates a new entity which can be freed later */
struct entity* new_entity() {
  struct entity* e = malloc(sizeof(struct entity));
  return e;
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

/** Inserts a symbol entity into the symbol table */
void insert_symbol_entity(struct entity* e) {

  if (e->name == NULL) {
    printf("Attempting to store a non-variable entity.\n");
    return;
  }

  int hash_index = hash(e->name);
  if (symbol_table[hash_index] != NULL) {
    if (strcmp(symbol_table[hash_index]->name, e->name)) {
      printf("There was just a hash collision in the symbol table.\n");
      printf("We're going to keep running, but I wanted you to know that its gonna break.\n");
    } else {

      // Redefining an already defined variable
      // I'm not going to error handle this quite yet

    }

  }

  symbol_table[hash_index] = e;

}

/** Returns the symbol of a given name
    Returns null if the symbol is undeclared */
struct entity* get_symbol(char* name) {
  return symbol_table[hash(name)];
}

/** Declares a new symbol under the given name
    This symbol will be of type undefined until it is later defined */
void declare_symbol(char* name) {
  struct entity* e = malloc(sizeof(struct entity));
  e->name = strdup(name);
  e->type = TYPE_UNDEFINED;
  e->value.string = "";
  insert_symbol_entity(e);
}

/** Updates the value of a given symbol to become an integer */
void update_symbol_i(char* name, int value) {
  struct entity* e = get_symbol(name);
  if (e == NULL) {
    printf("Variable Access Error (line %d): Assigning value %d to an undeclared variable %s\n", yylineno, value, name);
    e = malloc(sizeof(struct entity));
    e->name = strdup(name);
  }
  if (e->type == TYPE_STRING && e->value.string != NULL) {
    free(e->value.string);
  }
  e->value.number = value;
  e->type = TYPE_INTEGER;
}

/** Updates the value of a given symbol to become a string */
void update_symbol_s(char* name, char* value) {
  struct entity* e = get_symbol(name);
  if (e == NULL) {
    printf("Variable Access Error (line %d): Assigning new string to an undeclared variable %s\n", yylineno, name);
    e = malloc(sizeof(struct entity));
    e->name = strdup(name);
  }
  if (e->type == TYPE_STRING && e->value.string != NULL) {
    free(e->value.string);
  }
  e->value.string = strdup(value);
  e->type = TYPE_STRING;
}

/** Updates the value of a given symbol to become an object */
void update_value_o(char* name) {
  struct entity* e = get_symbol(name);
  if (e == NULL) {
    printf("Variable Access Error (line %d): Assigning new object to an undeclared variable %s\n", yylineno, name);
    e = malloc(sizeof(struct entity));
    e->name = strdup(name);
  }
  if (e->type == TYPE_STRING && e->value.string != NULL) {
    free(e->value.string);
  }
  e->value.string = "";
  e->type = TYPE_OBJECT;
}
