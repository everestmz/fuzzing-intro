# Math

## Addition

Let's assume we're building an addition function for a brand new programming language, so we can't just use a '+' operator. `add.go` contains an initial implementation of the addition function, and `add_fuzz.go` contains a fuzz test with checks for Commutativity and Associativity. 

- Commutativity. It shouldn’t matter what order I add numbers up in. I should get the same result

- Associativity: it doesn’t matter how we group the numbers for addition. The result should be the same.

## Fuzzing Add()

We can fuzz Add by building the fuzz test:

`$ go-fuzz-build`

and then running it:

`$ go-fuzz -workdir=workdir -bin=add-fuzz.zip`

But, we won't find any bugs!

## Adding the missing property

Can you see the problem in `add.go`? Why doesn't the fuzzer detect the issue when you run it?

Try adding a test for the third property of addition: the _identity_ property: x + 1 = x. This will cause the test to fail.

Why does this cause the fuzzer to break? Using the _input_ value to check the _output_ can be key to fuzzing in many cases. Note that associativity and commutativity didn't actually use the original numbers in the error checks.

