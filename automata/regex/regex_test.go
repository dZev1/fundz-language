package regex

import (
	"reflect"
	"testing"

	finiteautomata "github.com/dZev1/fundz-language/automata/finite_automata"
)

func TestRegexToString(t *testing.T) {
	tests := []struct {
		name string
		reg  Regex
		want string
	}{
		{
			name: "Lambda regex to string",
			reg:  Lambda{},
			want: "λ",
		},
		{
			name: "Empty set regex to string",
			reg:  EmptySet{},
			want: "∅",
		},
		{
			name: "Symbol regex to string",
			reg:  Symbol{Value: "a"},
			want: "a",
		},
		{
			name: "Concat regex to string",
			reg:  Concat{Left: Symbol{Value: "a"}, Right: Symbol{Value: "b"}},
			want: "(ab)",
		},
		{
			name: "Union regex to string",
			reg:  Union{Left: Symbol{Value: "a"}, Right: Symbol{Value: "b"}},
			want: "(a|b)",
		},
		{
			name: "Star regex to string",
			reg:  Star{Inner: Symbol{Value: "a"}},
			want: "(a)*",
		},
		{
			name: "Plus regex to string",
			reg:  Plus{Inner: Symbol{Value: "a"}},
			want: "(a)+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.reg.String(); got != tt.want {
				t.Errorf("Regex.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexToDFA(t *testing.T) {
	tests := []struct {
		name string
		reg  Regex
		want *finiteautomata.DFA[string]
	}{
		{
			name: "Symbol regex to DFA",
			reg:  Symbol{Value: "a"},
			want: &finiteautomata.DFA[string]{
				States:       map[string]struct{}{"q0": {}, "q1": {}},
				Alphabet:     map[string]struct{}{"a": {}},
				Transitions:  map[string]map[string]string{"q0": {"a": "q1"}, "q1": {}},
				InitialState: "q0",
				FinalStates:  map[string]struct{}{"q1": {}},
			},
		},
		{
			name: "Lambda regex to DFA",
			reg:  Lambda{},
			want: &finiteautomata.DFA[string]{
				States:       map[string]struct{}{"q0": {}},
				Alphabet:     map[string]struct{}{},
				Transitions:  map[string]map[string]string{"q0": {}},
				InitialState: "q0",
				FinalStates:  map[string]struct{}{"q0": {}},
			},
		},
		{
			name: "Empty set regex to DFA",
			reg:  EmptySet{},
			want: &finiteautomata.DFA[string]{
				States:       map[string]struct{}{"q0": {}},
				Alphabet:     map[string]struct{}{},
				Transitions:  map[string]map[string]string{"q0": {}},
				InitialState: "q0",
				FinalStates:  map[string]struct{}{},
			},
		},
		{
			name: "Plus regex to DFA",
			reg:  Plus{Inner: Symbol{Value: "a"}},
			want: &finiteautomata.DFA[string]{
				States:       map[string]struct{}{"q0": {}, "q1": {}},
				Alphabet:     map[string]struct{}{"a": {}},
				Transitions:  map[string]map[string]string{"q0": {"a": "q1"}, "q1": {"a": "q1"}},
				InitialState: "q0",
				FinalStates:  map[string]struct{}{"q1": {}},
			},
		},
		{
			name: "Concat regex to DFA",
			reg:  Concat{Left: Symbol{Value: "a"}, Right: Symbol{Value: "b"}},
			want: &finiteautomata.DFA[string]{
				States:       map[string]struct{}{"q0": {}, "q1": {}, "q2": {}},
				Alphabet:     map[string]struct{}{"a": {}, "b": {}},
				Transitions:  map[string]map[string]string{"q0": {"a": "q1"}, "q1": {"b": "q2"}, "q2": {}},
				InitialState: "q0",
				FinalStates:  map[string]struct{}{"q2": {}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.reg.toNFA().Determinize().Minimize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Regex.toDFA() = %v, want %v", got, tt.want)
			}
		})
	}
}
