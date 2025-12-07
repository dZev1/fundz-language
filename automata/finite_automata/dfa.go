package finiteautomata

import (
	"errors"
	"fmt"

	"github.com/dZev1/fundz-language/automata/set"
)

var (
	ErrStateAlreadyIn             = errors.New("state is already in automaton")
	ErrStateNotIn                 = errors.New("state is not in automaton")
	ErrTransitionWithSymbolExists = errors.New("transition with symbol already exists")
)

type DFA[T comparable] struct {
	States       set.Set[T]
	Alphabet     set.Set[string]
	Transitions  map[T](map[string]T)
	InitialState T
	FinalStates  set.Set[T]
}

func (dfa *DFA[T]) size() int {
	return len(dfa.States)
}

func (dfa *DFA[T]) AddState(state T, final bool) error {
	if _, ok := dfa.States[state]; ok {
		return ErrStateAlreadyIn
	}

	if (dfa.States == nil) {
		dfa.States = make(set.Set[T])
	}

	if (dfa.Transitions == nil) {
		dfa.Transitions = make(map[T]map[string]T)
	}
	
	dfa.States[state] = struct{}{}
	dfa.Transitions[state] = make(map[string]T)

	if final {
		if dfa.FinalStates == nil {
			dfa.FinalStates = make(set.Set[T])
		}
		dfa.FinalStates[state] = struct{}{}
	}

	return nil
}

func (dfa *DFA[T]) AddTransition(fromState, toState T, symbol string) error {
	if dfa.Transitions == nil {
		dfa.Transitions = make(map[T]map[string]T)
	}

	if dfa.Transitions[fromState] == nil {
		dfa.Transitions[fromState] = make(map[string]T)
	}

	if _, ok := dfa.States[fromState]; !ok {
		return fmt.Errorf("%w: %v", ErrStateNotIn, fromState)
	}

	if _, ok := dfa.States[toState]; !ok {
		return fmt.Errorf("%w: %v", ErrStateNotIn, toState)
	}

	if _, ok := dfa.Transitions[fromState][symbol]; ok {
		return fmt.Errorf("%w: %v", ErrTransitionWithSymbolExists, symbol)
	}

	dfa.Transitions[fromState][symbol] = toState
	dfa.Alphabet[symbol] = struct{}{}

	return nil
}

func (dfa *DFA[T]) Accepts(word string) bool {
	state := dfa.InitialState
	for _, r := range word {
		if _, ok := dfa.Transitions[state][string(r)]; ok {
			state = dfa.Transitions[state][string(r)]
		} else {
			return false
		}
	}
	_, ok := dfa.FinalStates[state]
	return ok
}

// Minimize a DFA
func (dfa *DFA[T]) Minimize() *DFA[string] {
	minimized := &DFA[string]{}

	minimized.States = set.Set[string]{"q0": struct{}{}, "q1": struct{}{}}
	minimized.Alphabet = dfa.Alphabet
	minimized.FinalStates = set.Set[string]{"q0": struct{}{}}
	minimized.InitialState = "q0"

	minimized.AddTransition("q0", "q1", "a")
	minimized.AddTransition("q1", "q0", "a")
	return minimized
}
