# Finding Heartbleed with feedback-driven fuzzing

To do this, we'll use libFuzzer. Like go-fuzz, but for LLVM, so good for C, C++, Rust etc.

## Attack surface:

Something that always happens is the OpenSSL Handshake, so let's write a fuzzer that sets up a server, and then acts as the client trying to make a handshake to that server.

## Harness:

You can see an example of a harness for the OpenSSL handshake in `handshake-fuzzer.cc`

## Download and unpack a vulnerable version of OpenSSL:
curl -O https://ftp.openssl.org/source/old/1.0.1/openssl-1.0.1f.tar.gz
tar xf openssl-1.0.1f.tar.gz

## Build OpenSSL:

_NOTE: it helps to run these commands on a linux install. If you're on macOS you can start a Linux docker container to run all of these commands inside with the following command:_

`docker run -w /src -v $PWD:/src -it gcr.io/oss-fuzz-base/base-clang /bin/bash`

If you use a container, you'll also need to set up your build environment by running:

`apt update`

`apt install -y git make g++`

If you're not using a container, make sure you have `git`, `make`, and a recent build of `clang` installed.

Once your build environment is set up:

`cd openssl-1.0.1f/`

`./config`

`make CC="clang -g -fsanitize=fuzzer-no-link"` (this might take a while)

`cd ..`


## Build OpenSSL fuzz target:
`clang++ -g handshake-fuzzer.cc -fsanitize=fuzzer openssl-1.0.1f/libssl.a openssl-1.0.1f/libcrypto.a -std=c++17 -Iopenssl-1.0.1f/include/ -lstdc++fs -ldl -lstdc++ -o handshake-fuzzer`

Finally, run the fuzzer: `./handshake-fuzzer`

Why don't we find the bug? Heartbleed is a memory corruption. This means we read or write outside of the bounds of memory we're allowed to, but this doesn't always crash our code! We need to add something to detect memory corruptions.

## Sanitizers

Sanitizers are things that add extra checks to your code. This one, AddressSanitizer, maps out your program's memory, and crashes it if there's a memory corruption, like a Stack Overflow.

## $CC must be pointing to clang binary, see the "compiler section" link above.
`cd openssl-1.0.1f/`

`make clean`

`make CC="clang -g -fsanitize=address,fuzzer-no-link"`

`cd ..`

`clang++ -g handshake-fuzzer.cc -fsanitize=address,fuzzer openssl-1.0.1f/libssl.a openssl-1.0.1f/libcrypto.a -std=c++17 -Iopenssl-1.0.1f/include/ -lstdc++fs -ldl -lstdc++ -o handshake-fuzzer-asan`

Finally, run: `./handshake-fuzzer-asan` and it should crash almost instantly with the bug!