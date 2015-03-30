# msp (a miniscript parser)

[![Build Status](https://magnum.travis-ci.com/mhoc/msp.svg?token=noEZqHCNie4E6CC2GGKT&branch=master)](https://magnum.travis-ci.com/mhoc/msp)

Michael Hockerman (mhockerm)
CS 352 Project 3
() April 2015

## Prerequisites:

* GoLang (http://golang.org) `go version`

## Setup

In the process of running `make`, the compile scripts will set the $GOPATH
envvar. If you use Go on a regular basis, this means it will temporarily
override your's. Assuming you have an `export GOPATH=` in your .bashrc/.zshrc,
this override will only last for as long as you keep this shell open.

Just something to keep in mind.

## Compiling

`make`

## Running

`./parser` will parse from stdin

`./parser {input file}` has desired behavior

## Cleaning

`make clean`

# General Compiler/AST Design

yyLex is called from main which runs lex on an input file then uses go yacc to
run semantic analysis. This process builds an AST for the input file. After
the AST is built, we DFS-traverse the AST to execute each step.

Each node in the ast is an ast.Node struct interface. This provides Execute()
functionality among other things. Every struct which implements ast.Node implements
an Execute() function which then "executes" that node, sometimes by re-calling
Execute() on its children, sometimes by accessing the symbol table, etc etc.
