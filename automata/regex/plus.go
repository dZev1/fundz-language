package regex

import (
    "github.com/dZev1/fundz-language/automata/finite_automata"
)

type Plus struct {
    Inner Regex
}

func (p Plus) String() string {
    return "(" + p.Inner.String() + ")+"
}

func (p Plus) toNFA() *finiteautomata.NFA[int] {
    innerNFA := p.Inner.toNFA()
    nfa := &finiteautomata.NFA[int]{}
    startState := 0
    nfa.AddState(startState, false)
    
    offset := 1
    for state := range innerNFA.States {
        isAccepting := false
        nfa.AddState(state+offset, isAccepting)
    }

    endState := offset + len(innerNFA.States)
    nfa.AddState(endState, true)

    for fromState, transitions := range innerNFA.Transitions {
        for symbol, toStates := range transitions {
            for toState := range toStates {
                nfa.AddTransition(fromState+offset, toState+offset, symbol)
            }
        }
    }
    
    nfa.AddTransition(startState, innerNFA.InitialState+offset, "")

    for finalState := range innerNFA.FinalStates {
        nfa.AddTransition(finalState+offset, endState, "")
        nfa.AddTransition(finalState+offset, innerNFA.InitialState+offset, "")
    }

    return nfa
}