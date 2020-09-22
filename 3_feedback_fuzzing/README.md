# Feedback Driven Fuzzing

## go-fuzz

Uses _code coverage feedback_ to decide which input is interesting. If an input covers new code than previous ones, it's saved for further mutation.

Install: `go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build`

## Limitations of mutational fuzzing alone:

Run `go build -o bin/broken ./main.go`

Execute `../1_mutational_fuzzing/fuzz.sh ./broken ./seeds`

No matter how long we run this for, very unlikely to actually find the bug! That's because of the nature of the bug: it's only triggered when the _exact_ input FUZ is executed. No more, no less.

## Enter coverage feedback-driven fuzzing: 

`go-fuzz-build -o bin/fuzz.zip ./fuzzme`

`go-fuzz -bin=bin/fuzz.zip -workdir=seeds`

You'll see that we find the bug in a matter of seconds! That's because the genetic algorithm of go-fuzz is much more effective at discovering bugs than blind mutation, since it saves it's progress. If you look at the corpus in `seeds/corpus`, you can see the fuzzer saving off interesting inputs as it makes its way through the code.