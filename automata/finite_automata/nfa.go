package finiteautomata

import "github.com/dZev1/fundz-language/automata/set"

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