# Mutational Fuzzing

This section is a quick introduction to the idea of Mutational fuzzing. It'll take you through how to use one of the most well-known mutators, Radamsa, to test software.

## Quick radamsa introduction

Radamsa is very simple - it takes whatever strings you give it, and modifies & squishes them together in new ways.

Install radamsa: https://gitlab.com/akihe/radamsa

To try it: `echo "Hello, world!" | radamsa`

Another way: `radamsa ./example.json`

## fuzz.sh: the world's smallest fuzzing tool

To use radamsa, we're going to build `fuzz.sh`, a tiny fuzzing tool that will run forever, generating new corrupted inputs to test our software with.

To fuzz a program that reads input from a file: `./fuzz.sh [file] [seed_directory]`