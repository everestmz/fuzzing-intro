package handlers

import (
	"testing"
)

func TestRLE(t *testing.T) {
	testCase := "ABCDDDDDDDD"
	encoded := RunLengthEncode(testCase)
	decoded := RunLengthDecode(encoded)
	if decoded != testCase {
		t.Fatalf("original %s != decoded %s", testCase, decoded)
	}
}
