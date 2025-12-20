package regex

import (
	"github.com/dZev1/fundz-language/automata/finite_automata"
)

type Star struct {
	Inner Regex
}

func (s Star) String() string {
	return "(" + s.Inner.String() + ")*"
}

func (s Star) toNFA() *finiteautomata.NFA[int] {
	innerNFA := s.Inner.toNFA()
	nfa := &finiteautomata.NFA[int]{}
	offset := 1
	nfa.AddState(0, true)

	for state := range innerNFA.States {
		isAccepting := false
		if _, ok := innerNFA.FinalStates[state]; ok {
			isAccepting = true
		}
		nfa.AddState(state+offset, isAccepting)
	}

	for fromState, transitions := range innerNFA.Transitions {
		for symbol, toStates := range transitions {
			for toState := range toStates {
				nfa.AddTransition(fromState+offset, toState+offset, symbol)
			}
		}
	}

	nfa.AddTransition(0, 1, "")
	for finalState := range innerNFA.FinalStates {
		nfa.AddTransition(finalState+offset, 0, "")
	}

	return nfa
}