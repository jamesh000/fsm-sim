package main

import (
	"fmt"
	"strings"
)

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
				return nil, fmt.Errorf("second start state defined on line %d", i+1)
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
				return nil, fmt.Errorf("invalid state name %q on line %d", cTo[1], i+1)
			}

			// append the target to the appropriate transition array
			if c == "0" {
				stateMap[stateName].transitions[0] = append(stateMap[stateName].transitions[0], target)
			} else if c == "1" {
				stateMap[stateName].transitions[1] = append(stateMap[stateName].transitions[1], target)
			} else if c == "e" {
				stateMap[stateName].epsilonTransitions = append(stateMap[stateName].epsilonTransitions, target)
			} else {
				return nil, fmt.Errorf("invalid character %q on line %d", c, i+1)
			}
		}
	}

	return startState, nil
}
