
// #include "my.h" done in including file

// Symbol table
struct symbol* symbol_table[SYMBOL_TABLE_SIZE];

struct tok* tok_new() {
  struct tok* token = malloc(sizeof(struct tok));
  return token;
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

/** Inserts a symbol into the symbol table */
void insert_symbol(struct symbol* s) {

  int hash_index = hash(s->name);
  if (symbol_table[hash_index] != NULL) {
    if (strcmp(symbol_table[hash_index]->name, s->name)) {
      printf("There was just a hash collision in the symbol table.\n");
      printf("We're going to keep running, but I wanted you to know that its gonna break.\n");

    } else {
      //redefining an already defined variable

    }

  }

  symbol_table[hash_index] = s;

}

void declare_symbol(char* name) {
  struct symbol* s = malloc(sizeof(struct symbol));
  s->name = strdup(name);
  s->type = TYPE_UNDEFINED;
  s->value.string = "";

  insert_symbol(s);
}

/** Inserts raw symbol data into the symbol table */
void insert_symbol_i(char* name, int value) {

  struct symbol* s = malloc(sizeof(struct symbol));
  s->name = strdup(name);
  s->type = TYPE_INTEGER;
  s->value.number = value;

  insert_symbol(s);

}

void insert_symbol_s(char* name, char* value) {
  struct symbol* s = malloc(sizeof(struct symbol));
  s->name = strdup(name);
  s->type = TYPE_STRING;
  s->value.string = strdup(value);

  insert_symbol(s);
}

void insert_symbol_o(char* name) {
  struct symbol* s = malloc(sizeof(struct symbol));
  s->name = strdup(name);
  s->type = TYPE_OBJECT;
  s->value.string = "";

  insert_symbol(s);
}

struct symbol* get_symbol(char* name) {
  return symbol_table[hash(name)];
}

void update_value_i(char* name, int value) {
  struct symbol* sym = get_symbol(name);
  if (sym == NULL) {
    printf("Error line %d: Assigning value to an undeclared variable\n", yylineno);
    sym = malloc(sizeof(struct symbol));
    sym->name = strdup(name);
  }
  if (sym->type == TYPE_STRING) {
    free(sym->value.string);
  }
  sym->value.number = value;
  sym->type = TYPE_INTEGER;
}

void update_value_s(char* name, char* value) {
  struct symbol* sym = get_symbol(name);
  if (sym == NULL) {
    printf("Error line %d: Assigning value to an undeclared variable\n", yylineno);
    sym = malloc(sizeof(struct symbol));
    sym->name = strdup(name);
  }
  if (sym->type == TYPE_STRING) {
    free(sym->value.string);
  }
  sym->value.string = strdup(value);
  sym->type = TYPE_STRING;
}

void update_value_o(char* name) {
  struct symbol* sym = get_symbol(name);
  if (sym == NULL) {
    printf("Error line %d: Assigning value to an undeclared variable\n", yylineno);
    sym = malloc(sizeof(struct symbol));
    sym->name = strdup(name);
  }
  if (sym->type == TYPE_STRING) {
    free(sym->value.string);
  }
  sym->value.string = "";
  sym->type = TYPE_OBJECT;
}
