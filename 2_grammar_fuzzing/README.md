# Grammar Fuzzing

Useful for fuzzing programs that consume a very well-defined input language.

## Dharma

To install: `pip install dharma` (Python required!)

Think back to when we mutated that one JSON file with Radamsa - we got some interesting inputs, but they all looked very similar. When using a gramma fuzzer, you can get very varied inputs that still conform to a format, without needing to provide examples!

## To generate data:

Use case: fuzzing a JSON parser

`dharma -grammars ./json/dg`

`dharma -grammars ./json.dg -count 100`

Use case: generate lots of URLs that look weird, but are technically correctly-formatted. Can be used to ensure your URL parser is accurate

`dharma -grammars ./url.dg -count 100`