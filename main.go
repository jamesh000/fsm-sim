package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jamesh000/fsm-sim/stack"
)

type state struct {
	name               string
	transitions        [2][]*state
	epsilonTransitions []*state
	final              bool
}

type partialExe struct {
	s              *state
	remainingInput string
}

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

	execute(startState, str)
}

func tokenize(strIn string) [][]string {
	lines := strings.Split(strIn, "\n")

	tokensByLine := make([][]string, len(lines))

	for i, line := range lines {
		tokensByLine[i] = strings.Fields(line)
	}

	return tokensByLine
}

func parse(tokens [][]string) (*state, error) {
	// map from states to their names
	stateMap := make(map[string]*state)
	var startState *state

	for i, line := range tokens {
		s := new(state)

		if line[0][0] == '(' {
			// check for multiple starts
			if startState != nil {
				return nil, fmt.Errorf("second start state defined on line %d", i)
			}

			startState = s

			// clip the signifier
			line[0] = line[0][1:]
		} else if line[0][0] == ')' {
			s.final = true

			// clip the signifier
			line[0] = line[0][1:]
		}

		s.name = line[0]

		stateMap[s.name] = s
	}

	for i, line := range tokens {
		stateName := line[0]
		transitions := line[1:]

		// add transitions
		for _, t := range transitions {
			cTo := strings.Split(t, ":")
			c := cTo[0]
			target := stateMap[cTo[1]]

			if target == nil {
				return nil, fmt.Errorf("invalid state name %q on line %d", cTo[1], i)
			}

			// append the target to the appropriate transition array
			if c == "0" {
				stateMap[stateName].transitions[0] = append(stateMap[stateName].transitions[0], target)
			} else if c == "1" {
				stateMap[stateName].transitions[1] = append(stateMap[stateName].transitions[1], target)
			} else if c == "e" {
				stateMap[stateName].epsilonTransitions = append(stateMap[stateName].epsilonTransitions, target)
			} else {
				return nil, fmt.Errorf("invalid character %q on line %d", c, i)
			}
		}
	}

	return startState, nil
}

func execute(initialState *state, input string) {
	exeStack := stack.Stack[partialExe]{}

	// initial execution
	exeStack.Push(partialExe{initialState, input})

	for !exeStack.IsEmpty() {
		// get next execution
		e := exeStack.Pop()

		fmt.Println("Moving to state", e.s.name)

		if len(e.remainingInput) == 0 {
			fmt.Println("End of input reached on state", e.s.name)
			if e.s.final {
				fmt.Println("The input is valid!")
				return
			} else {
				fmt.Println("Path is dead")
				fmt.Println()
			}

			// go to the next thing
			continue
		}

		for _, s := range e.s.epsilonTransitions {
			exeStack.Push(partialExe{s, e.remainingInput})
		}

		c := e.remainingInput[0]
		switch c {
		case '0':
			for _, s := range e.s.transitions[0] {
				exeStack.Push(partialExe{s, e.remainingInput[1:]})
			}
		case '1':
			for _, s := range e.s.transitions[1] {
				exeStack.Push(partialExe{s, e.remainingInput[1:]})
			}
		}
	}

	fmt.Printf("No more paths to take...")
}
