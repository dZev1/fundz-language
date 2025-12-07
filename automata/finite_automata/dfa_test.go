package finiteautomata

import (
	"testing"

	"github.com/dZev1/fundz-language/automata/set"
)

func TestDFA_AddState(t *testing.T) {
	tests := []struct {
		name  string
		state int
		final bool
	}{
		{
			name:  "Add a non-final state",
			state: 4,
			final: false,
		},
		{
			name:  "Add a final state",
			state: 4,
			final: true,
		},
		{
			name:  "Add an already added state",
			state: 3,
			final: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfa := DFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]int{
					0: {"a": 1, "b": 2},
					1: {"a": 0, "b": 3},
					2: {"a": 2, "b": 2},
					3: {"a": 2, "b": 2},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			}
			previousSize := dfa.size()
			err := dfa.AddState(tt.state, tt.final)

			if err == ErrStateAlreadyIn {
				if dfa.size() != previousSize {
					t.Errorf("expected dfa of size: %v, but got: %v", previousSize, dfa.size())
				}
			} else {
				if dfa.size() != previousSize+1 {
					t.Errorf("expected dfa of size: %v but got: %v", previousSize+1, dfa.size())
				}
			}
		})
	}
}

func TestDFA_AddTransition(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fromState int
		toState   int
		symbol    string
		wantErr   bool
	}{
		{
			name:      "Add transition from 1 to 2 with symbol b",
			fromState: 1,
			toState:   2,
			symbol:    "b",
			wantErr:   false,
		},
		{
			name:      "Add transition from 1 to 2 with symbol not in alphabet",
			fromState: 1,
			toState:   2,
			symbol:    "c",
			wantErr:   false,
		},
		{
			name:      "Transition with symbol already exists",
			fromState: 2,
			toState:   3,
			symbol:    "b",
			wantErr:   true,
		},
		{
			name:      "Transition to unexisting state",
			fromState: 0,
			toState:   10,
			symbol:    "a",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfa := DFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]int{
					0: {"a": 1, "b": 2},
					1: {"a": 0},
					2: {"a": 2, "b": 2},
					3: {"a": 2, "b": 2},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			}
			gotErr := dfa.AddTransition(tt.fromState, tt.toState, tt.symbol)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddTransition() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddTransition() succeeded unexpectedly")
			}
		})
	}
}

func TestDFA_Accepts(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "Word accepted",
			word: "aaab",
			want: true,
		},
		{
			name: "Word not accepted",
			word: "aaa",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfa := DFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}, "b": {}},
				Transitions: map[int]map[string]int{
					0: {"a": 1, "b": 2},
					1: {"a": 0, "b": 3},
					2: {"a": 2, "b": 2},
					3: {"a": 2, "b": 2},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			}
			got := dfa.Accepts(tt.word)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Accepts() = %v, want %v", got, tt.want)
			}
		})
	}
}
