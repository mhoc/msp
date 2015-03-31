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

`./parser {input file}` to provide a file to parse

## Cleaning

`make clean`
