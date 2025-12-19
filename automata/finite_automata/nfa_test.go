package finiteautomata

import (
	"reflect"
	"testing"

	"github.com/dZev1/fundz-language/automata/set"
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

func TestNFA_Determinize(t *testing.T) {
	tests := []struct {
		name string
		nfa  *NFA[int]
		want *DFA[string]
	}{
		{
			name: "Determinize a NFA without instant transitions",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}, 2: {}},
					},
					1: {
						"a": set.Set[int]{2: {}},
					},
					2: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{2: {}},
			},
			want: &DFA[string]{
				States:       set.Set[string]{"q0": {}, "q1": {}, "q2": {}, "q3": {}},
				Alphabet:     set.Set[string]{"a": {}},
				Transitions:  map[string]map[string]string{"q0": {"a": "q1"}, "q1": {"a": "q2"}, "q2": {"a": "q3"}, "q3": {"a": "q3"}},
				InitialState: "q0",
				FinalStates:  set.Set[string]{"q1": {}, "q2": {}, "q3": {}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nfa.Determinize()
			
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Determinize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNFA_lambdaClosure(t *testing.T) {
	tests := []struct {
		name   string
		nfa    *NFA[int]
		states set.Set[int]
		want   set.Set[int]
	}{
		{
			name: "Lambda closure of a set of only one state",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}, 2: {}},
					},
					1: {
						"a": set.Set[int]{2: {}},
					},
					2: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{2: {}},
			},
			states: set.Set[int]{0:{}},
			want: set.Set[int]{0:{}},
		},
		{
			name: "Lambda closure of a set of multiple states",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}},
						"":  set.Set[int]{2: {}},
					},
					1: {
						"a": set.Set[int]{3: {}},
					},
					2: {
						"": set.Set[int]{3: {}},
					},
					3: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			states: set.Set[int]{0:{}},
			want: set.Set[int]{0:{},2:{},3:{}},
		},
		{
			name: "Lambda closure of an empty set",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}},
						"":  set.Set[int]{2: {}},
					},
					1: {
						"a": set.Set[int]{3: {}},
					},
					2: {
						"": set.Set[int]{3: {}},
					},
					3: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			states: set.Set[int]{},
			want:   set.Set[int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.nfa.lambdaClosure(tt.states)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lambdaClosure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNFA_move(t *testing.T) {
	tests := []struct {
		name   string
		nfa    *NFA[int]
		states set.Set[int]
		symbol string
		want   set.Set[int]
	}{
		{
			name: "Move from a set of states with a symbol",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}, 2: {}},
					},
					1: {
						"a": set.Set[int]{3: {}},
					},
					2: {
						"a": {3: {}},
					},
					3: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			states: set.Set[int]{0: {}, 1: {}},
			symbol: "a",
			want:   set.Set[int]{1: {}, 2: {}, 3: {}},
		},
		{
			name: "Move from an empty set of states",
			nfa: &NFA[int]{
				States:   set.Set[int]{0: {}, 1: {}, 2: {}, 3: {}},
				Alphabet: set.Set[string]{"a": {}},
				Transitions: map[int]map[string]set.Set[int]{
					0: {
						"a": set.Set[int]{1: {}, 2: {}},
					},
					1: {
						"a": set.Set[int]{3: {}},
					},
					2: {
						"a": {3: {}},
					},
					3: {
						"a": {0: {}},
					},
				},
				InitialState: 0,
				FinalStates:  set.Set[int]{3: {}},
			},
			states: set.Set[int]{},
			symbol: "a",
			want:   set.Set[int]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.nfa.move(tt.states, tt.symbol)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}
