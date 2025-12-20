package regex

import (
    "github.com/dZev1/fundz-language/automata/finite_automata"
)

type Concat struct {
    Left  Regex
    Right Regex
}

func (c Concat) String() string {
    return "(" + c.Left.String() + c.Right.String() + ")"
}

func (c Concat) toNFA() *finiteautomata.NFA[int] {
    leftNFA := c.Left.toNFA()
    rightNFA := c.Right.toNFA()
    nfa := &finiteautomata.NFA[int]{}

    for state := range leftNFA.States {
        nfa.AddState(state, false)
    }

    offset := len(leftNFA.States)
    for state := range rightNFA.States {
        isAccepting := false
        if _, ok := rightNFA.FinalStates[state]; ok {
            isAccepting = true
        }
        nfa.AddState(state+offset, isAccepting)
    }

    for fromState, transitions := range leftNFA.Transitions {
        for symbol, toStates := range transitions {
            for toState := range toStates {
                nfa.AddTransition(fromState, toState, symbol)
            }
        }
    }

    for fromState, transitions := range rightNFA.Transitions {
        for symbol, toStates := range transitions {
            for toState := range toStates {
                nfa.AddTransition(fromState+offset, toState+offset, symbol)
            }
        }
    }

    for finalState := range leftNFA.FinalStates {
        nfa.AddTransition(finalState, rightNFA.InitialState+offset, "")
    }

    return nfa
}