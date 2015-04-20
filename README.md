## msp (a miniscript parser)

[![Build Status](https://magnum.travis-ci.com/mhoc/msp.svg?token=noEZqHCNie4E6CC2GGKT&branch=master)](https://magnum.travis-ci.com/mhoc/msp)

Michael Hockerman (mhockerm)
CS 352 Project 3
3 April 2015

## Prerequisites:

* GoLang (http://golang.org) `go version`

## Setup

In the process of running `make`, the compile scripts will set the $GOPATH
envvar. If you use Go on a regular basis, this means it will temporarily
override your's. Assuming you have an `export GOPATH=` in your .bashrc/.zshrc,
this override will only last for as long as you keep this shell open.

This won't cause any problems with compilation, its just something to keep
in mind.

## Compiling

`make`

## Running

`./parser` will parse from stdin

`./parser [input file]` to provide a file to parse

## Cleaning

`make clean`

## Extensions

The parser can be provided an additional flag `./parser -extensions {input file}` to enable
the optional parser extensions. With these enabled, additional functionality is enabled
which could cause it to fail the official test cases, but which is overall very handy
for debugging and general usage. An outline of these additional features is provided
below:

* Enables printing of arrays and objects inside document.write() instead of throwing
an error and printing undefined.

Example: `document.write([1,2,3,4])` Prints `[1,2,3,4]`

* Enables use of function `len(v) -> int`

This function provides the length of both strings and arrays. Attempting to use any
other types on it will throw a type violation and return undefined. Attempting to
redefine this function will succeed.
