# Mutational Fuzzing

This section is a quick introduction to the idea of Mutational fuzzing. It'll take you through how to use one of the most well-known mutators, Radamsa, to test software.

## Quick radamsa introduction

Radamsa is very simple - it takes whatever strings you give it, and modifies & squishes them together in new ways.

Install radamsa: https://gitlab.com/akihe/radamsa

To try it: `echo "Hello, world!" | radamsa`

Another way: `radamsa ./example.json`

## fuzz.sh: the world's smallest fuzzing tool

To use radamsa, we're going to build `fuzz.sh`, a tiny fuzzing tool that will run forever, generating new corrupted inputs to test our software with.

To fuzz a program that reads input from a file: `./fuzz.sh [program] [seed_directory]`, where `program` is the program you'd like to test, and `seed_directory` is a directory of example inputs for radamsa to start from.

write this program into `fuzz.sh`

```#!/bin/bash
while true; do
	# Generate a test with radamsa
  radamsa $2/* > input.txt
	# Display the test
	cat input.txt
	# Execute the program we're testing with the test
  $1 input.txt > /dev/null 2>&1
	# If it crashed, save the input and exit
  if [ $? -gt 1 ]; then
    cp input.txt crash.txt
    echo "Crash found!"
		exit 1
  fi
done
```

Don't forget to make sure it's executable: `chmod 775 ./fuzz.sh`.

## Finding our first bug:

run `go build -o bin/buggy ./buggy_program.go` to compile our program.

It's a program that checks if the words "foo", "bar", and "omg" are in the input file we give it, and then returns how many of those words were present. It seems pretty simple, but there's a bug in it that we'll find with fuzzing.

Run: `./fuzz.sh ./bin/buggy ./seeds`.
