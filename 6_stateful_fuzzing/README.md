# Stateful Fuzzing

In `heap.go` we have an implementation of a min-heap, but it's got a bug! We can run heap_fuzz.go by building the fuzzer:

`$ go-fuzz-build`

and then running it:

`$ go-fuzz -workdir=workdir -bin=heap-fuzz.zip`

But, we won't find the bug!

## Careful harnessing

It turns out our fuzz test isn't robust enough. We only remove one item from the heap. But a correct implementation of a min-heap requires that every call to `Remove()` return the current smallest item in the heap, not just the first.

Extend the fuzz test to ensure we check this property carefully. (there's some code you can uncomment)