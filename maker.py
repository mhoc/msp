
# Handles building, cleaning, and testing the project
# in one fell swoop

import os
import subprocess
import sys

build_graph = {
    "__start": "parser",

    "parser": {
        "__depends": ["y.tab.c", "lex.yy.c"],
        "__do": "gcc y.tab.c lex.yy.c -o parser -lfl -w"
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

clean_cmd = "rm -f lex.yy.c y.tab.c y.tab.h y.dot y.output parser"

test_parameters = {
    "directory": "test/",
    "command": "./parser {}",

    # Include each test as an object here.
    # The test runner will run each of these
    "tests": {

        "001_just-script-tags": {
            "expect_nothing": True
        },
        "002_script-tags-with-newlines": {
            "expect_nothing": True
        },
        "003_script-tags-not-at-start": {
            "expect_nothing": True,
            "optional": True
            # Handout says that script tags will be at start and end, but this shouldn't be hard to implement
        },
        "004_empty-file": {
            "expect_anything": True
        },
        "005_missing-start-script-tag": {
            "expect_anything": True
        },
        "006_missing-end-script-tag": {
            "expect_anything": True
        },
        "007_incorrect-script-tag-spelling": {
            "expect_anything": True
        },
        "008_single-declaration": {
            "expect_nothing": True
        },
        "009_single-declaration-semi": {
            "expect_nothing": True
        },
        "010_newline-inside-statement": {
            "expect_anything": True,
            "optional": True
            # Piazza @80. Li says that he doesnt want miniscript to be like js in this regard, but they wont test for it
        },
        "011_multiple-declarations-one-line": {
            "expect_nothing": True
        },
        "012_single-int-assignment": {
            "expect_nothing": True
        },
        "013_single-string-assignment": {
            "expect_nothing": True
        },
        "014_single-expression-assignment": {
            "expect_nothing": True
        },
        "015_multi-assignment": {
            "expect_nothing": True
        },
        "016_assign-with-nonexist-operator": {
            "expect_anything": True
        },
        "017_define-single-int": {
            "expect_nothing": True
        },
        "018_define-single-string": {
            "expect_nothing": True
        },
        "019_var-with-numbers": {
            "expect_nothing": True
        },
        "020_var-start-with-number": {
            "expect_anything": True
        },
        "021_newline-in-string": {
            "expect_anything": True
        },
        "022_assign-var-to-var": {
            "expect_nothing": True
        },
        "023_document-write-no-args": {
            "expect_nothing": True
        },
        "024_document-write-string": {
            "expect_nothing": True
        },
        "025_document-write-var": {
            "expect_nothing": True
        },
        "026_document-write-multi-vars": {
            "expect_nothing": True
        },
        "027_handout-example": {
            "expect_nothing": True
        },
        "028_advanced-parens": {
            "expect_nothing": True
        },
        "029_string-special-chars": {
            "expect_nothing": True
        },
        "030_multiple-statements-no-delim": {
            "expect_anything": True
        },
        "031_java-code": {
            "expect_anything": True
        },
        "032_complex-1": {
            "expect_nothing": True
        },
        "033_document-write-paren-expression": {
            "expect_nothing": True
        },
        "034_odd-whitespace": {
            "expect_nothing": True
        },
        "035_document-write-no-param-list": {
            "expect_anything": True
        }

    }
}

def build():
    build_level(build_graph["__start"])

def build_level(key):
    for dependency in build_graph[key]["__depends"]:
        build_level(dependency)
    print "====> Building " + key
    subprocess.call(build_graph[key]["__do"].split(" "))

def build_with_make():
    subprocess.call(["make", "clean"])
    subprocess.call(["make"])

def clean():
    subprocess.call(clean_cmd.split(" "))

def test():
    passed = 0
    total = 0

    for key in sorted(test_parameters["tests"].keys()):
        value = test_parameters["tests"][key]
        total += 1
        result = run_test(key, value)
        if result:
            passed += 1

    passed_perc = str(float(passed)/total)
    print "\nPassed " + str(passed) + " / " + str(total) + " tests (" + passed_perc + ")."
    if passed_perc == "1.0":
        print_flag()

def run_test(file, test_obj):
    test_file = test_parameters["directory"] + file
    command = test_parameters["command"].format(test_file).split(" ")
    output = subprocess.check_output(command, stderr=subprocess.STDOUT)
    test_result = validate_test(file, output, test_obj)
    return test_result

def validate_test(file, output, test_obj):
    result = False

    if "expect_nothing" in test_obj and test_obj["expect_nothing"]:
        result = output == ""

    if "expect_anything" in test_obj and test_obj["expect_anything"]:
        result = not output == ""

    # Check if the test is optional
    optional = "optional" in test_obj and test_obj["optional"]
    if result:
        report_pass(file)
    elif not result and optional:
        report_optional_fail(file, output)
    else:
        report_fail(file, output)

    return result

def report_pass(file):
    print_green(u"\u2713")
    print "\t" + file

def report_fail(file, output):
    print_red(u"\n\u2718")
    print "\t" + file
    print output

def report_optional_fail(file, output):
    print_yellow(u"\n\u2718")
    print "\t" + file + " (optional)"
    print output

def print_green(str):
    sys.stdout.write("\033[0;32m" + str + "\033[0;00m")

def print_yellow(str):
    sys.stdout.write("\033[1;33m" + str + "\033[0;00m")

def print_red(str):
    sys.stdout.write("\033[0;31m" + str + "\033[0;00m")

def print_flag():
    print ""
    print "\033[1;37mCongratulations soldier!"
    print "You've got the greatest compiler of all time. Your country thanks you for your courage!"
    print "                               \033[1;36m~~ Bill Clinton\033[0;00m"
    print ""
    print "\033[0;31mAMERICA \033[1;37mAMERICA \033[0;34mAMERICA \033[0;00m"
    print "\033[1;37m__________--___---____"
    print "|\033[1;34m * * * *\033[1;37m |\033[1;37m--\__\--\__\033[1;37m|"
    print "|\033[1;34m* * * * *\033[1;37m|\033[0;31m---\__\--\_\033[1;37m|"
    print "|\033[1;34m_*_*_*_*_\033[1;37m|\033[1;37m\---\__\---\033[1;37m|"
    print "|\033[0;31m___________\---\__\--\033[1;37m|"
    print "|\033[1;37m____________\---\__\-\033[1;37m|"
    print "|\033[0;31m_____________\---\___\033[1;37m|"
    print "||"
    print "||"
    print "||"
    print "||"
    print "||"
    print "||"
    print "||"
    print "||"
    print ""

def generate_makefile():
    fp = open("Makefile", "w")
    
    # Create a list which will contain a list of the things we will print out later
    list = []
    make_build_recurse(build_graph["__start"], list)

    # Print out to the makefile
    for line in list:
        fp.write(line)

    # Write out the clean command
    fp.write("clean:\n")
    fp.write("\t" + clean_cmd + "\n\n")

    fp.close()

def make_build_recurse(key, list):
    
    # Add this level to the list because it comes first
    title = key + ": "
    for dependency in build_graph[key]["__depends"]:
        title += dependency + " "
    list.append(title + "\n\t")
    list.append(build_graph[key]["__do"] + "\n\n")

    # Recurse into the dependencies of this level
    for dependency in build_graph[key]["__depends"]:
        make_build_recurse(dependency, list)

def print_usage():
    print "python maker.py [command]"
    print "no command : does all three! build -> test -> clean!\n"
    print "commands:"
    print "\tbuild    : builds the project as defined in the build_graph"
    print "\tclean    : cleans the folder as defined in clean_cmd"
    print "\ttest     : tests the project as defined in testing_parameters and test/ folder"
    print "\tmakefile : "
    

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

elif sys.argv[1] == "help" or sys.argv[1] == "--help":
    print_usage()

elif sys.argv[1] == "makefile":
    generate_makefile()
