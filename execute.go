package main

import (
	"fmt"

	"github.com/jamesh000/fsm-sim/stack"
)

type partialExe struct {
	fromState      string
	s              *state
	remainingInput string
}

func execute(initialState *state, input string) (bool, error) {
	exeStack := stack.Stack[partialExe]{}

	// initial execution
	exeStack.Push(partialExe{"initialization", initialState, input})

	for !exeStack.IsEmpty() {
		// get next execution
		e := exeStack.Pop()

		fmt.Printf("Moving to state %s from state %s with remaining input %q\n", e.s.name, e.fromState, e.remainingInput)

		if len(e.remainingInput) == 0 {
			fmt.Println("End of input reached on state", e.s.name)
			if e.s.final {
				fmt.Println("The input is valid!")
				return true, nil
			} else {
				fmt.Println("This path is dead due to end of input")
				fmt.Println()
			}

			// go to the next thing
			continue
		}

		for _, s := range e.s.epsilonTransitions {
			exeStack.Push(partialExe{e.s.name, s, e.remainingInput})
		}

		c := e.remainingInput[0]
		switch c {
		case '0':
			if len(e.s.transitions[0]) == 0 {
				fmt.Println("This path is dead due to nowhere to go")
				fmt.Println()
			}

			for _, s := range e.s.transitions[0] {
				exeStack.Push(partialExe{e.s.name, s, e.remainingInput[1:]})
			}
		case '1':
			for _, s := range e.s.transitions[1] {
				exeStack.Push(partialExe{e.s.name, s, e.remainingInput[1:]})
			}
		default:
			return false, fmt.Errorf("invalid character %q", c)
		}
	}

	fmt.Println("No more paths to take...")
	return false, nil
}
