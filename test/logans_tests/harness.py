#!/usr/bin/python3
import os
import subprocess

TESTDIR = os.getcwd() + "/tests/"
PARSER_EXE = "./parser"

def colorize(color_code, text):
    return "\033[%dm%s\033[0m" % (color_code, text)

def red(text):
    return colorize(31, text)
def green(text):
    return colorize(32, text)
def yellow(text):
    return colorize(33, text)

def getCommand(filename):
    return [PARSER_EXE, filename]

def getCorrectOutput(filename):
    s = ""
    with open(filename + "_correct") as f:
        s = f.read()
    return s

def main():
    num_tests = 0
    incorrect = []

    # Ensure the program has been built
    out = subprocess.call("make")

    nl = True
    for filename in sorted(os.listdir(TESTDIR)):
        if "_correct" in filename or filename == "test":
            continue
        else:
            num_tests += 1
            command = getCommand(TESTDIR + filename)
            process = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            out,err = process.communicate()
            out += err
            out = out.decode("utf-8")

            # Correct output
            correct = getCorrectOutput(TESTDIR + filename)
            if out == correct:
                print(green("."), end="")
                nl = False
            else:
                if not nl:
                    print()
                print(red("%s: failed" % filename))
                incorrect.append(filename)
                nl = True

    if not nl:
        print()

    if len(incorrect) > 0:
        print("Failed {} tests.".format(len(incorrect)))
        incorrect.sort()
        print(", ".join(incorrect))
    print("Ran {} tests.".format(num_tests))

main()
