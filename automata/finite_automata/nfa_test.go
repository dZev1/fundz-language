package finiteautomata

import (
	"github.com/dZev1/fundz-language/automata/set"
	"testing"
)

func TestNFA_AddState(t *testing.T) {
	tests := []struct {
		nfa   *NFA[int]
		name  string
		state int
		final bool
	}{
		{
			nfa:   &NFA[int]{},
			name:  "Add state to empty automaton",
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

func TestNFA_AddTransition(t *testing.T) {
	tests := []struct {
		name      string
		fromState int
		toState   int
		symbol    string
	}{
		{
			name:      "Add valid transition",
			fromState: 1,
			toState:   2,
			symbol:    "b",
		},
		{
			name:      "Add transition from non-existing state",
			fromState: 5,
			toState:   2,
			symbol:    "a",
		},
		{
			name:      "Add transition to non-existing state",
			fromState: 1,
			toState:   6,
			symbol:    "a",
		},
		{
			name:      "Add non-deterministic transition",
			fromState: 1,
			toState:   3,
			symbol:    "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nfa := NFA[int]{
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
			}

			err := nfa.AddTransition(tt.fromState, tt.toState, tt.symbol)

			if tt.name == "Add valid transition" || tt.name == "Add non-deterministic transition" {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				} else {
					if _, ok := nfa.Transitions[tt.fromState][tt.symbol]; !ok {
						t.Errorf("expected transition from state %v with symbol %v to exist", tt.fromState, tt.symbol)
					} else {
						if _, ok := nfa.Transitions[tt.fromState][tt.symbol][tt.toState]; !ok {
							t.Errorf("expected transition to state %v to exist", tt.toState)
						}
					}
				}
			} else {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			}
		})
	}
}

func TestNFA_Accepts(t *testing.T) {
	nfa := NFA[int]{
		States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
		Alphabet: set.Set[string]{"a": {}, "b": {}},
		Transitions: map[int]map[string]set.Set[int]{
			0: {
				"a": {1: {}, 2: {}},
				"b": {2: {}},
			},
			1: {
				"a": {0: {}},
				"b": {3: {}},
			},
			2: {
				"a": {2: {}, 3: {}},
				"b": {2: {}},
			},
			3: {
				"a": {2: {}},
				"b": {2: {}},
			},
		},
		InitialState: 0,
		FinalStates:  set.Set[int]{3: {}},
	}

	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "Accepted word 'ab'",
			word: "ab",
			want: true,
		},
		{
			name: "Rejected word 'aab'",
			word: "aab",
			want: false,
		},
		{
			name: "Accepted word 'aa'",
			word: "aa",
			want: true,
		},
		{
			name: "Rejected word 'bbab'",
			word: "bbab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nfa.Accepts(tt.word)
			if got != tt.want {
				t.Errorf("NFA.Accepts(%v) = %v; want %v", tt.word, got, tt.want)
			}
		})
	}
}
