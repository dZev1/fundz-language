package regex

import (
	"github.com/dZev1/fundz-language/automata/finite_automata"
)

type Regex interface {
	toNFA() *finiteautomata.NFA[int]
	String() string
}

func ToNFA(r Regex) *finiteautomata.NFA[int] {
	return r.toNFA()
}

func BuildAutomatonFromRegex(r Regex) *finiteautomata.DFA[string] {
	nfa := ToNFA(r)
	return nfa.Determinize().Minimize()
}

func RegexToString(r Regex) string {
	return r.String()
}