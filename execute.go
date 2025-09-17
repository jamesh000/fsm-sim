package main

import (
	"fmt"

	"github.com/jamesh000/fsm-sim/stack"
)

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
