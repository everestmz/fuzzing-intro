# Fuzzing a small REST service

This is a small web application that consumes a JSON payload `{ "to_encode": "ABCDDDDDDDD" }`, and returns a run length encoded version: `{ "encoded": "1A1B1C8D" }`

Try it! `go run ./main.go`, and then `curl -i -H "Content-Type: application/json" -X POST -d '{"to_encode":"ABCDDDDDDDDDD"}' localhost:1323/rle`

## Attack Surface

What's the attack surface here? In this case, it's the JSON body we parse. In another web application it might be worth trying a header, like the Authorization header, or potentially Query Parameters. We can see here, however, that the part this webapp actually reads from is the JSON body, so that's what we'll focus on.

## Harness

So what's the ideal place to harness? We could hook a fuzzer up to a command line utility like cURL, and pipe something like Radamsa in for mutation. That would work, but we all know that the network is pretty slow, and we wouldn't get very quick fuzzing. Moreover, the Go network code has been fuzzed quite extensively, so we'll really just be wasting CPU clock cycles if we end up testing that any more.

How about just putting data directly into the JSON payload of the http method then? that could work too, but since the JSON decoding code of Go is also well-tested, we'd just be wasting lots of cycles as the fuzzer generates a ton of garbage tests that don't actually conform to the input format that we see in the structs `EncodeRequest` and `EncodeResponse`.

Since we're confident about those two parts of our stack, we could just tackle the contents of the `EncodeHandler` directly, by providing it a valid `EncodeRequest`. In this case, we could even fuzz the `RunLengthEncode` method directly, but we'd like this fuzzer to be future-proof, so if the handler itself changes at some point, the fuzzer still applies to our code, agnostic of what's actually going on inside the handler

## Tooling

Let's build our fuzzer: `go-fuzz-build -o fuzz_1.zip ./handlers`

And then run it: `go-fuzz -bin=fuzz_1.zip -workdir=workdir`

## Defining properties
So, we've got a fuzzer that sends garbled data into `RunLengthEncode`. Why aren't we finding any bugs? Well, similar to our heartbleed example, there's a chance that if there are any bugs in our code, they're a little more subtle than a bug that just crashes it outright.

For heartbleed, we used Address Sanitizer. The _invariant_ we were working with is "This code has no memory corruption bugs". Since this web service is run-length-encoding strings, why not assert that it actually encodes them properly, by trying to decode them after?

Add this snippet to the bottom of the Fuzz method:

```
	response := &EncodeResponse{}
	err = json.Unmarshal(rec.Body.Bytes(), response)
	if err != nil {
		panic(err)
	}
	if RunLengthDecode(response.Encoded) != fuzzPayload.ToEncode {
		panic("decoded != encoded")
	}
```

This will check to ensure that if we decode the value we receive back, we'll get the original value again.

Let's build our fuzzer again: `go-fuzz-build -o fuzz_2.zip ./handlers`

And then run it: `go-fuzz -bin=fuzz_2.zip -workdir=workdir`

You should see a crash pop out almost immediately! If we inspect it, we can see that we incorrectly wrote our Encode method, since on line 50 we never reset the counter when we receive a new variable. Since our unit test only decided to test Run Length Encoding a string with repeated characters at the end, we never caught this state issue before.

