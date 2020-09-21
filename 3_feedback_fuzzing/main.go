package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/everestmz/fuzzingdemo/fuzzme"
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

	fmt.Println(fuzzme.BrokenMethod(inputData))
}
