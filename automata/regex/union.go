package regex

import (
    "github.com/dZev1/fundz-language/automata/finite_automata"
)

type Union struct {
    Left  Regex
    Right Regex
}

func (u Union) String() string {
    return "(" + u.Left.String() + "|" + u.Right.String() + ")"
}

func (u Union) toNFA() *finiteautomata.NFA[int] {
    leftNFA := u.Left.toNFA()
    rightNFA := u.Right.toNFA()

    nfa := &finiteautomata.NFA[int]{}
    startState := 0
    nfa.AddState(startState, false)
    
    leftOffset := 1
    for state := range leftNFA.States {
        isAccepting := false
        if _, ok := leftNFA.FinalStates[state]; ok {
            isAccepting = true
        }
        nfa.AddState(state+leftOffset, isAccepting)
    }

    rightOffset := leftOffset + len(leftNFA.States)
    for state := range rightNFA.States {
        isAccepting := false
        if _, ok := rightNFA.FinalStates[state]; ok {
            isAccepting = true
        }
        nfa.AddState(state+rightOffset, isAccepting)
    }

    for fromState, transitions := range leftNFA.Transitions {
        for symbol, toStates := range transitions {
            for toState := range toStates {
                nfa.AddTransition(fromState+leftOffset, toState+leftOffset, symbol)
            }
        }
    }

    for fromState, transitions := range rightNFA.Transitions {
        for symbol, toStates := range transitions {
            for toState := range toStates {
                nfa.AddTransition(fromState+rightOffset, toState+rightOffset, symbol)
            }
        }
    }

    nfa.AddTransition(startState, leftNFA.InitialState+leftOffset, "")
    nfa.AddTransition(startState, rightNFA.InitialState+rightOffset, "")

    return nfa
}