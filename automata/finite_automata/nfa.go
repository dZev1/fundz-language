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
	
	if symbol != "" {
		nfa.Alphabet[symbol] = struct{}{}
	}

	return nil
}

func (nfa *NFA[T]) Determinize() (DFA[string], error) {
	result := DFA[string]{}

	result.Alphabet = nfa.Alphabet

	return result, nil
}

func (nfa *NFA[T]) String() string {
	result := "NFA:\n"
	result += "States: "
	for state := range nfa.States {
		result += fmt.Sprintf("%v ", state)
	}
	result += "\nAlphabet: "
	for symbol := range nfa.Alphabet {
		result += fmt.Sprintf("%v ", symbol)
	}
	result += "\nTransitions:\n"
	for fromState, transitions := range nfa.Transitions {
		for symbol, toStates := range transitions {
			for toState := range toStates {
				result += fmt.Sprintf("  %v --%v--> %v\n", fromState, symbol, toState)
			}
		}
	}
	result += fmt.Sprintf("Initial State: %v\n", nfa.InitialState)
	result += "Final States: "
	for state := range nfa.FinalStates {
		result += fmt.Sprintf("%v ", state)
	}
	result += "\n"
	return result
}


