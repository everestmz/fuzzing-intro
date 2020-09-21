# Finding Heartbleed with feedback-driven fuzzing

To do this, we'll use libFuzzer. Like go-fuzz, but for LLVM, so good for C, C++, Rust etc.

## Download and unpack a vulnerable version of OpenSSL:
curl -O https://ftp.openssl.org/source/old/1.0.1/openssl-1.0.1f.tar.gz
tar xf openssl-1.0.1f.tar.gz

## Download the fuzz target and its data dependencies:

Something that always happens is the OpenSSL Handshake - seems like a good place to start!

`curl -O https://raw.githubusercontent.com/google/clusterfuzz/master/docs/setting-up-fuzzing/heartbleed/handshake-fuzzer.cc`
`curl -O https://raw.githubusercontent.com/google/clusterfuzz/master/docs/setting-up-fuzzing/heartbleed/server.key`
`curl -O https://raw.githubusercontent.com/google/clusterfuzz/master/docs/setting-up-fuzzing/heartbleed/server.pem`


## Build OpenSSL:
`cd openssl-1.0.1f/`
`./config` ON MACOS: `./Configure darwin64-x86_64-cc`
`make CC="clang -g -fsanitize=fuzzer-no-link"`
`cd ..`


## Build OpenSSL fuzz target:
`clang++ -g handshake-fuzzer.cc -fsanitize=address,fuzzer openssl-1.0.1f/libssl.a openssl-1.0.1f/libcrypto.a -std=c++17 -Iopenssl-1.0.1f/include/ -lstdc++fs -ldl -lstdc++ -o handshake-fuzzer`

Finally, run the fuzzer: `./handshake-fuzzer`

Why don't we find the bug? Heartbleed is a memory corruption. This means we read or write outside of the bounds of memory we're allowed to, but this doesn't always crash our code! We need to add something to detect memory corruptions.

## Sanitizers!

Sanitizers are things that add extra checks to your code. This one maps out your program's memory, and crashes it if there's a memory corruption, like a Stack Overflow.

## $CC must be pointing to clang binary, see the "compiler section" link above.
make CC="clang -g -fsanitize=address,fuzzer-no-link"
cd ..