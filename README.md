# fsm-sim
Finite state machine simulator. Simulates both determinate and nondeterminate FSMs. Usage is
```
fsm-sim <definition-file> <input-string>
```
Only strings of {0, 1} are supported. Because I am lazy and that is all that is on my homework.

## FSM Definition Syntax
Each line defines one state and its transitions. First the state name is given. The only character in the state name that is important is
the first. If it is `(`, this state is the sole initial state. If it is `)`, this state is a final state. Then there is a list of transitions seperated by whitespace.
Each transition is of the form `<character>:<state-name>` are represents a transition to the named state on that character. The character
can either be 0, 1, or e. 0 and 1 are self explanatory. e is epsilon. The behavior is described in my textbook as follows:

>If a state with an ε symbol on an exiting arrow is encountered, something
>similar happens. Without reading any input, the machine splits into multiple
>copies, one following each of the exiting ε-labeled arrows and one staying at the
>current state. Then the machine proceeds nondeterministically as before.

So that is how it works in this simulator. Example definitions can be found in the examples folder.
