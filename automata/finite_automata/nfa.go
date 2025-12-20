package finiteautomata

import (
	"fmt"
	"slices"

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

func (nfa *NFA[T]) Determinize() *DFA[string] {
	result := &DFA[string]{
		States: make(set.Set[string]),
		Alphabet: make(set.Set[string]),
		Transitions: make(map[string]map[string]string),
		FinalStates: make(set.Set[string]),
	}
	
	for symbol := range nfa.Alphabet {
		result.Alphabet[symbol] = struct{}{}
	}
	
	currIndex := 0
	initialState, isFinal := nfa.lambdaClosure(set.Set[T]{nfa.InitialState:{}})
	initialStateID := fmt.Sprintf("q%d", currIndex)
	
	mapNFAtoDFA := make(map[string]string)
	setToID := make(map[string]string)
	
	initialStateStr := serializeSet(initialState)
	setToID[initialStateStr] = initialStateID
	mapNFAtoDFA[initialStateStr] = initialStateID
	
	result.States[initialStateID] = struct{}{}
	if isFinal {
		result.FinalStates[initialStateID] = struct{}{}
	}

	currIndex++
	result.InitialState = initialStateID
	
	queue := []set.Set[T]{initialState}
	processed := make(map[string]bool)

	for len(queue) > 0 {
		states := queue[0]
		queue = queue[1:]
		
		statesStr := serializeSet(states)
		if processed[statesStr] {
			continue
		}
		processed[statesStr] = true

		dfaFromState := setToID[statesStr]
		
		for symbol := range nfa.Alphabet {
			reachableStates, isFinal := nfa.move(states, symbol)

			// Skip empty transitions - don't create dead states
			if len(reachableStates) == 0 {
				continue
			}

			reachableStr := serializeSet(reachableStates)
			
			if _, exists := setToID[reachableStr]; !exists {
				dfaState := fmt.Sprintf("q%d", currIndex)
				setToID[reachableStr] = dfaState
				result.States[dfaState] = struct{}{}
				if isFinal {
					result.FinalStates[dfaState] = struct{}{}
				}
				queue = append(queue, reachableStates)
				currIndex++
			}
			
			dfaToState := setToID[reachableStr]
			result.AddTransition(dfaFromState, dfaToState, symbol)
		}
	}

	return result
}

func serializeSet[T comparable](s set.Set[T]) string {
	var strs []string
	for state := range s {
		strs = append(strs, fmt.Sprintf("%v", state))
	}
	slices.Sort(strs)
	return "{" + fmt.Sprintf("%v", strs) + "}"
}

func (nfa *NFA[T]) lambdaClosure(states set.Set[T]) (set.Set[T], bool) {
	result := make(set.Set[T])

	isFinal := false

	stack := make([]T, 0)
	visited := make(set.Set[T])

	for state := range states {
		stack = append(stack, state)
		result[state] = struct{}{}
		visited[state] = struct{}{}
	}

	for len(stack) > 0 {
		currentState := stack[len(stack)-1]
		if _, ok := nfa.FinalStates[currentState]; ok {
			isFinal = true
		}
		stack = stack[:len(stack)-1]
		
		transitions, ok := nfa.Transitions[currentState]
		if !ok {
			continue
		}
		
		lambdaTransitions, ok := transitions[""]
		if !ok {
			continue
		}
		
		for toState := range lambdaTransitions {
			if _, seen := visited[toState]; !seen {
				result[toState] = struct{}{}
				visited[toState] = struct{}{}
				stack = append(stack, toState)
			}
		}
	}
	
	return result, isFinal
}


func (nfa *NFA[T]) move(states set.Set[T], symbol string) (set.Set[T], bool) {
	reachable := make(set.Set[T])

	for state := range states {
		transitions, ok := nfa.Transitions[state]
		if !ok {
			continue
		}

		toStates, ok := transitions[symbol]
		if !ok {
			continue
		}

		for toState := range toStates {
			set.Add(&reachable, toState)
		}
	}
	reachableClosure, isFinal := nfa.lambdaClosure(reachable)
	return reachableClosure, isFinal
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


func (nfa *NFA[T]) NormalizeStates(prefix string) {
	stateMap := make(map[T]T)
	newStates := make(set.Set[T])
	newTransitions := make(map[T]map[string]set.Set[T])
	var index int = 0

	for state := range nfa.States {
		newState := any(fmt.Sprintf("%s%d", prefix, index)).(T)
		stateMap[state] = newState
		newStates[newState] = struct{}{}
		index++
	}
	for fromState, transitions := range nfa.Transitions {
		newFromState := stateMap[fromState]
		newTransitions[newFromState] = make(map[string]set.Set[T])
		for symbol, toStates := range transitions {
			newTransitions[newFromState][symbol] = make(set.Set[T])
			for toState := range toStates {
				newToState := stateMap[toState]
				newTransitions[newFromState][symbol][newToState] = struct{}{}
			}
		}
	}
	nfa.States = newStates
	nfa.Transitions = newTransitions
	nfa.InitialState = stateMap[nfa.InitialState]
	newFinalStates := make(set.Set[T])
	for finalState := range nfa.FinalStates {
		newFinalStates[stateMap[finalState]] = struct{}{}
	}
	nfa.FinalStates = newFinalStates
}