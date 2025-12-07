package finiteautomata

import (
	"errors"
	"fmt"

	"github.com/dZev1/fundz-language/automata/set"
)

var (
	ErrStateAlreadyIn = errors.New("state is already in automaton")
	ErrStateNotIn     = errors.New("state is not in automaton")
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
	dfa.States[state] = struct{}{}
	dfa.Transitions[state] = map[string]T{}

	if final {
		dfa.FinalStates[state] = struct{}{}
	}

	return nil
}

func (dfa *DFA[T]) AddTransition(fromState, toState T, symbol string) error {
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
	q := dfa.InitialState
	for _, r := range word {
		if _, ok := dfa.Transitions[q][string(r)]; ok {
			q = dfa.Transitions[q][string(r)]
		} else {
			return false
		}
	}
	_, ok := dfa.FinalStates[q]
	return ok
}