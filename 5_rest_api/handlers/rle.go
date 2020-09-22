package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//RunLengthDecode expands a string in run length encoding
func RunLengthDecode(s string) string {
	//var b strings.Builder
	var code = regexp.MustCompile("^[0-9]*[A-Za-z ]")
	var b bytes.Buffer
	for n := 0; n < len(s); {
		short := code.FindString(s[n:])
		var long string
		var err error
		if len(short) == 1 {
			long = short
		} else {
			long, err = Expand(short)
			if err != nil {
				return "Encoding failed as I " + err.Error()
			}
		}
		//b.WriteString(long)
		b.Write([]byte(long))
		n += len(short)
	}
	return b.String()
}

//RunLengthEncode compresses a string using runlength encoding
func RunLengthEncode(s string) string {
	var result string
	var prev byte = 0
	var current byte = 0

	count := 1
	for i := 0; i < len(s); i++ {
		current = s[i]
		if current != prev {
			if prev != 0 {
				result += fmt.Sprint(count) + string(prev)
			}
			// Oops!
			prev = current
		} else {
			count += 1
		}
	}
	result += fmt.Sprint(count) + string(current)
	return result
}

//Expand takes a code like "6X" and returns "XXXXXX"
//The code must be in form number-letter
func Expand(s string) (string, error) {
	if len(s) < 2 {
		return "", errors.New("cannot convert " + s + " too short")
	}
	number, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return "", errors.New("cannot convert " + s + " No number")
	}
	return strings.Repeat(s[len(s)-1:], number), nil
}

//Contract takes a string like "XXXXXX" and returns 6X
func Contract(s string) string {
	return strconv.Itoa(len(s)) + s[:1]
}
