package regex

import (
    "github.com/dZev1/fundz-language/automata/finite_automata"
)

type Lambda struct{}

func (l Lambda) String() string {
    return "λ"
}

func (l Lambda) toNFA() *finiteautomata.NFA[int] {
    nfa := &finiteautomata.NFA[int]{}
    nfa.AddState(0, false)
    nfa.AddState(1, true)
    nfa.AddTransition(0, 1, "")
    return nfa
}