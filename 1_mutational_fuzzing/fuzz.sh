#!/bin/bash
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