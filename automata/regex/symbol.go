package regex

import (
	"github.com/dZev1/fundz-language/automata/finite_automata"
)

type Symbol struct {
	Value string
}

func (s Symbol) String() string {
	return s.Value
}

func (s Symbol) toNFA() *finiteautomata.NFA[int] {
	nfa := &finiteautomata.NFA[int]{}
	nfa.AddState(0, false)
	nfa.AddState(1, true)
	nfa.AddTransition(0, 1, s.Value)
	return nfa
}