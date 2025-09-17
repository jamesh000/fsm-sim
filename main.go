package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Wrong number of arguments, needs 2")
		return
	}

	fmtDefFile := args[0]
	str := args[1]

	fmtDefString, err := os.ReadFile(fmtDefFile)

	if err != nil {
		panic(err)
	}

	tokens := tokenize(string(fmtDefString))

	fmt.Println(tokens)

	startState, err := parse(tokens)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Parsed successfully, startstate with arrays", *startState, startState.transitions[0][0], startState.transitions[1][0])

	result, err := execute(startState, str)

	if err != nil {
		panic(err)
	}

	if result {
		fmt.Println("Very cool")
	} else {
		fmt.Println("How unfortunate")
	}
}
