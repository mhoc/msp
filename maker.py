
# A replacement for stupid make files lol

import os
import subprocess
import sys

build_graph = {
    "__start": "parser",
    
    "parser": {
        "__depends": ["y.tab.c", "lex.yy.c"],
        "__do": "gcc y.tab.c lex.yy.c -o parser -lfl"
    },
    
    "y.tab.c": {
        "__depends": [],
        "__do": "bison -y -d -g -t --verbose yacc.y"
    },

    "lex.yy.c": {
        "__depends": [],
        "__do": "lex lex.l"
    }
}

clean_cmd = "rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output"

test_parameters = {
    "__directory": "tests/",
    "__command": "./parser {}",
    
    "__pass_behavior": "no output",
    "__exceptions": {
        "provided_1": {
            "__pass_behavior": "any output"
        },
        "provided_2": {
            "__pass_behavior": "any output"
        }
    }
}

def build():
    build_level(build_graph["__start"], 1)

def build_level(key, l):
    for dependency in build_graph[key]["__depends"]:
        build_level(dependency, l + 1)
    arrow = ""
    while l >= 0:
        arrow += "="
        l -= 1
    print arrow + "> Building " + key
    subprocess.call(build_graph[key]["__do"].split(" "))

def build_with_make():
    subprocess.call(["make", "clean"])
    subprocess.call(["make"])

def clean():
    subprocess.call(clean_cmd.split(" "))

def test():
    passed = 0
    total = 0
    for file in os.listdir(test_parameters["__directory"]):
        if os.path.isdir(test_parameters["__directory"] + file):
            continue
        total += 1
        if file in test_parameters["__exceptions"]:
            result = run_test(file, test_parameters["__exceptions"][file]["__pass_behavior"])
        else:
            result = run_test(file, test_parameters["__pass_behavior"])
        if result:
            passed += 1
    print "\nPassed " + str(passed) + " / " + str(total) + " tests (" + str(float(passed)/total) + ")."

def run_test(file, expect):
    actual_name = test_parameters["__directory"] + file
    output = subprocess.check_output(test_parameters["__command"].format(actual_name).split(" "), stderr=subprocess.STDOUT)
    if expect == "no output":
        if output == "":
            print_test_pass(file)
            return True
        else:
            print_test_fail(file, output)
            return False
    elif expect == "any output":
        if output == "":
            print_test_fail(file, output)
            return False
        else:
            print_test_pass(file)
            return True

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


