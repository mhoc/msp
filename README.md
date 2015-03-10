# msp (a miniscript parser)

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
