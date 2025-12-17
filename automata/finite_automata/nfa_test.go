package finiteautomata

import (
	"testing"
	"github.com/dZev1/fundz-language/automata/set"
)

func TestNFA_AddState(t *testing.T) {
	tests := []struct {
		nfa *NFA[int]
		name  string
		state int
		final bool
	}{
		{
			nfa: &NFA[int]{},
			name: "Add state to empty automaton",
			state: 0,
			final: true,
		},
		{
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {"a": {1: {}}, "b": {2: {}}},
					1: {"a": {0: {}}, "b": {3: {}}},
					2: {"a": {2: {}}, "b": {2: {}}},
					3: {"a": {2: {}}, "b": {2: {}}},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			name:  "Add a non-final state",
			state: 4,
			final: false,
		},
		{
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {"a": {1: {}}, "b": {2: {}}},
					1: {"a": {0: {}}, "b": {3: {}}},
					2: {"a": {2: {}}, "b": {2: {}}},
					3: {"a": {2: {}}, "b": {2: {}}},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			name:  "Add a final state",
			state: 4,
			final: true,
		},
		{
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {"a": {1: {}}, "b": {2: {}}},
					1: {"a": {0: {}}, "b": {3: {}}},
					2: {"a": {2: {}}, "b": {2: {}}},
					3: {"a": {2: {}}, "b": {2: {}}},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			name:  "Add an already added state",
			state: 3,
			final: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nfa := tt.nfa
			previousSize := nfa.size()
			err := nfa.AddState(tt.state, tt.final)

			if err == ErrStateAlreadyIn {
				if nfa.size() != previousSize {
					t.Errorf("expected nfa of size: %v, but got: %v", previousSize, nfa.size())
				}
			} else {
				if nfa.size() != previousSize+1 {
					t.Errorf("expected nfa of size: %v but got: %v", previousSize+1, nfa.size())
				}
			}
		})
	}
}

/*
func TestNFA_AddTransition(t *testing.T) {
	tests := []struct {
		nfa *NFA[int]
		name string
		fromState int
		toState int
		symbol string	
	}{
		{
			nfa 
		}
	}
}
	*/