package main

type state struct {
	name               string
	transitions        [2][]*state
	epsilonTransitions []*state
	final              bool
}
