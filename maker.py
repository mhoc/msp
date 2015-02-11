
# A replacement for stupid make files lol

import math
import os
import subprocess
import sys

build_cmds = [
    "lex lex.l",
    "bison -y -d -g -t --verbose yacc.y",
    "gcc y.tab.c lex.yy.c -o parser -lfl"
]

clean_cmd = "rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output"

test_cmd = "./parser {}"

def build():
    for cmd in build_cmds:
        subprocess.call(cmd.split(" "))

def clean():
    subprocess.call(clean_cmd.split(" "))

def test():
    passed = 0
    total = 0
    for file in os.listdir("test"):
        if ".expected" in file:
            continue
        cmd = test_cmd.format("test/" + file)
        result = subprocess.check_output(cmd.split(" "), stderr=subprocess.STDOUT)
        passing = open(file + ".expected", "r").read() == result

        if passing:
            print_test_pass(file)
        else:
            print_test_fail(file, result)

    print "\nPassed " + str(passed) + " / " + str(total) + " tests (" + str(math.ceil(100*float(passed)/total)) + ")."

def print_test_pass(file):
    print u"\033[0;32m\u2713" + "\033[0;00m\t" + file

def print_test_fail(file, output):
    print ""
    print u"\033[0;31m\u2718" + "\t" + file
    print output
    print "\033[0;00m"

def print_usage():
    print "python maker.py [command]"
    print "commands:"
    print "\tbuild : runs the build() process as defined in the build graph"
    print "\tclean : runs the command supplied in the clean command"
    print "\ttest  : runs the testing framework as provided"
    print "no command : does all three! build -> test -> clean!"

# Check arguments

if len(sys.argv) > 2:
    print_usage()
    exit()

if len(sys.argv) == 1:
    build()
    print "\nRunning Tests..."
    test()
    print "\nCleaning..."
    clean()

elif sys.argv[1] == "build":
    build()

elif sys.argv[1] == "test":
    test()

elif sys.argv[1] == "clean":
    clean()


