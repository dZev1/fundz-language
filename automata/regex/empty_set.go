package regex

import (
	"github.com/dZev1/fundz-language/automata/finite_automata"
)

type EmptySet struct{}

func (e EmptySet) String() string {
	return "∅"
}

func (e EmptySet) toNFA() *finiteautomata.NFA[int] {
	nfa := &finiteautomata.NFA[int]{}
	nfa.AddState(0, false)
	return nfa
}