package finiteautomata

import (
	"fmt"

	"github.com/dZev1/fundz-language/automata/set"
)

type NFA[T comparable] struct {
	States       set.Set[T]
	Alphabet     set.Set[string]
	Transitions  map[T](map[string]set.Set[T])
	InitialState T
	FinalStates  set.Set[T]
}

// size returns the number of states in the NFA
func (nfa *NFA[T]) size() int {
	return len(nfa.States)
}

func (nfa *NFA[T]) AddState(state T, final bool) error {
	if nfa.States == nil {
		nfa.States = make(set.Set[T])
	}

	if nfa.Transitions == nil {
		nfa.Transitions = make(map[T]map[string]set.Set[T])
	}

	if _, ok := nfa.States[state]; ok {
		return ErrStateAlreadyIn
	}


	nfa.States[state] = struct{}{}
	nfa.Transitions[state] =  make(map[string]set.Set[T])

	if final {
		if nfa.FinalStates == nil {
			nfa.FinalStates = make(set.Set[T])
		}

		nfa.FinalStates[state] = struct{}{}
	}
	
	return nil
}

func (nfa *NFA[T]) AddTransition(fromState, toState T, symbol string) error {
	if nfa.Transitions == nil {
		nfa.Transitions = make(map[T]map[string]set.Set[T])
	}

	if nfa.Alphabet == nil {
		nfa.Alphabet = make(set.Set[string])
	}

	if _, ok := nfa.States[fromState]; !ok {
		return fmt.Errorf("%w: %v", ErrStateNotIn, fromState)
	}

	if _, ok := nfa.States[toState]; !ok {
		return fmt.Errorf("%w: %v", ErrStateNotIn, toState)
	}

	if nfa.Transitions[fromState] == nil {
		nfa.Transitions[fromState] = make(map[string]set.Set[T])
	}

	if nfa.Transitions[fromState][symbol] == nil {
		nfa.Transitions[fromState][symbol] = make(set.Set[T])
	}
	
	nfa.Transitions[fromState][symbol][toState] = struct{}{}
	nfa.Alphabet[symbol] = struct{}{}

	return nil
}

// Accepts checks if the NFA accepts the given word
func (nfa *NFA[T]) Accepts(word string) bool {
	currentStates := make(set.Set[T])
	currentStates[nfa.InitialState] = struct{}{}

	for _, r := range word {
		nextStates := make(set.Set[T])
		symbol := string(r)

		for state := range currentStates {
			if transitions, ok := nfa.Transitions[state][symbol]; ok {
				for nextState := range transitions {
					nextStates[nextState] = struct{}{}
				}
			}
		}

		currentStates = nextStates

		if len(currentStates) == 0 {
			return false
		}
	}

	for state := range currentStates {
		if _, ok := nfa.FinalStates[state]; ok {
			return true
		}
	}

	return false
}


