package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	inputData := string(input)

	vec := []int{0, 1, 2}

	i := 0
	if strings.Contains(inputData, "foo") {
		i++
	}
	if strings.Contains(inputData, "bar") {
		i++
	}
	if strings.Contains(inputData, "omg") {
		i++
	}

	fmt.Println(vec[i])
}
