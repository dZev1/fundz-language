package finiteautomata

import (
	"errors"
	"fmt"

	"github.com/dZev1/fundz-language/automata/set"
)

// Error definitions
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

// size returns the number of states in the DFA
func (dfa *DFA[T]) size() int {
	return len(dfa.States)
}

// AddState adds a new state to the DFA. If final is true, the state is added to the set of final states.
func (dfa *DFA[T]) AddState(state T, final bool) error {
	if _, ok := dfa.States[state]; ok {
		return ErrStateAlreadyIn
	}

	if dfa.States == nil {
		dfa.States = make(set.Set[T])
	}

	if dfa.Transitions == nil {
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

// AddTransition adds a transition from fromState to toState with the given symbol symbol
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

// returns true if the word is accepted by the DFA
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

// Minimize a DFA using Hopcroft's algorithm
func (dfa *DFA[T]) Minimize() *DFA[string] {
	dfa.RemoveUnreachableStates()
	
	finalPartition := dfa.FinalStates
	nonFinalPartition := make(set.Set[T])
	for state := range dfa.States {
		if _, ok := finalPartition[state]; !ok {
			nonFinalPartition[state] = struct{}{}
		}
	}

	partitions := []set.Set[T]{}
	if len(nonFinalPartition) > 0 {
		partitions = append(partitions, nonFinalPartition)
	}
	if len(finalPartition) > 0 {
		partitions = append(partitions, finalPartition)
	}

	// Step 2: Refine partitions until no more refinement is needed
	changed := true
	for changed {
		changed = false
		newPartitions := []set.Set[T]{}

		for _, partition := range partitions {
			refined := refinePartition(partition, partitions, dfa)
			for _, subPartition := range refined {
				if len(subPartition) > 0 {
					newPartitions = append(newPartitions, subPartition)
				}
			}
			if len(refined) > 1 {
				changed = true
			}
		}

		partitions = newPartitions
	}

	minimized := &DFA[string]{}
	minimized.States = make(set.Set[string])
	minimized.Alphabet = dfa.Alphabet
	minimized.FinalStates = make(set.Set[string])
	minimized.Transitions = make(map[string]map[string]string)

	initialPartitionIndex := -1
	for i, partition := range partitions {
		if _, ok := partition[dfa.InitialState]; ok {
			initialPartitionIndex = i
			break
		}
	}

	if initialPartitionIndex > 0 {
		initialPartition := partitions[initialPartitionIndex]
		partitions = append([]set.Set[T]{initialPartition}, append(partitions[:initialPartitionIndex], partitions[initialPartitionIndex+1:]...)...)
	}

	partitionMap := make(map[T]int)
	for i, partition := range partitions {
		for state := range partition {
			partitionMap[state] = i
		}
	}

	for i := 0; i < len(partitions); i++ {
		stateName := fmt.Sprintf("q%d", i)
		minimized.States[stateName] = struct{}{}
		minimized.Transitions[stateName] = make(map[string]string)
	}

	minimized.InitialState = "q0"

	for state := range dfa.FinalStates {
		partitionID := partitionMap[state]
		minimized.FinalStates[fmt.Sprintf("q%d", partitionID)] = struct{}{}
	}

	for fromState := range dfa.States {
		fromPartitionID := partitionMap[fromState]
		fromStateName := fmt.Sprintf("q%d", fromPartitionID)

		for symbol, toState := range dfa.Transitions[fromState] {
			toPartitionID := partitionMap[toState]
			toStateName := fmt.Sprintf("q%d", toPartitionID)
			minimized.Transitions[fromStateName][symbol] = toStateName
		}
	}

	return minimized.NormalizeStates()
}

// refinePartition splits a partition based on transitions to other partitions
func refinePartition[T comparable](partition set.Set[T], partitions []set.Set[T], dfa *DFA[T]) []set.Set[T] {
	if len(partition) <= 1 {
		return []set.Set[T]{partition}
	}

	var symbols []string
	for symbol := range dfa.Alphabet {
		symbols = append(symbols, symbol)
	}

	subPartitions := []set.Set[T]{partition}

	for _, symbol := range symbols {
		newSubPartitions := []set.Set[T]{}

		for _, subPartition := range subPartitions {
			split := splitBySymbol(subPartition, symbol, partitions, dfa)
			newSubPartitions = append(newSubPartitions, split...)
		}

		subPartitions = newSubPartitions
	}

	return subPartitions
}

// splitBySymbol splits a partition based on which partition the symbol leads to
func splitBySymbol[T comparable](partition set.Set[T], symbol string, partitions []set.Set[T], dfa *DFA[T]) []set.Set[T] {
	if len(partition) <= 1 {
		return []set.Set[T]{partition}
	}

	groups := make(map[int]set.Set[T])

	for state := range partition {
		targetPartition := -1
		if transitions, ok := dfa.Transitions[state]; ok {
			if target, ok := transitions[symbol]; ok {
				targetPartition = findPartition(target, partitions)
			}
		}

		if groups[targetPartition] == nil {
			groups[targetPartition] = make(set.Set[T])
		}
		groups[targetPartition][state] = struct{}{}
	}

	result := []set.Set[T]{}
	for _, group := range groups {
		if len(group) > 0 {
			result = append(result, group)
		}
	}

	return result
}

// findPartition returns the index of the partition containing the state
func findPartition[T comparable](state T, partitions []set.Set[T]) int {
	for i, partition := range partitions {
		if _, ok := partition[state]; ok {
			return i
		}
	}
	return -1
}

func (dfa *DFA[T]) String() string {
	result := "DFA:\n"
	result += "States: "
	for state := range dfa.States {
		result += fmt.Sprintf("%v ", state)
	}
	result += "\nAlphabet: "
	for symbol := range dfa.Alphabet {
		result += fmt.Sprintf("%s ", symbol)
	}
	result += "\nTransitions:\n"
	for fromState, transitions := range dfa.Transitions {
		for symbol, toState := range transitions {
			result += fmt.Sprintf("  %v --%s--> %v\n", fromState, symbol, toState)
		}
	}
	result += fmt.Sprintf("Initial State: %v\n", dfa.InitialState)
	result += "Final States: "
	for state := range dfa.FinalStates {
		result += fmt.Sprintf("%v ", state)
	}
	result += "\n"
	return result
}

func (dfa *DFA[T]) NormalizeStates() *DFA[string] {
	stateMap := make(map[T]string)
	newStates := make(set.Set[string])
	newTransitions := make(map[string]map[string]string)
	var index int = 0
	for state := range dfa.States {
		newState := fmt.Sprintf("q%d", index)
		stateMap[state] = newState
		newStates[newState] = struct{}{}
		index++
	}
	for fromState, transitions := range dfa.Transitions {
		newFromState := stateMap[fromState]
		newTransitions[newFromState] = make(map[string]string)
		for symbol, toState := range transitions {
			newToState := stateMap[toState]
			newTransitions[newFromState][symbol] = newToState
		}
	}
	normalized := &DFA[string]{
		States:       newStates,
		Alphabet:     dfa.Alphabet,
		Transitions:  newTransitions,
		InitialState: stateMap[dfa.InitialState],
		FinalStates:  make(set.Set[string]),
	}
	for finalState := range dfa.FinalStates {
		normalized.FinalStates[stateMap[finalState]] = struct{}{}
	}
	return normalized
}

// RemoveUnreachableStates removes states that cannot be reached from the initial state
func (dfa *DFA[T]) RemoveUnreachableStates() *DFA[T] {
    reachable := make(set.Set[T])
    queue := []T{dfa.InitialState}
    reachable[dfa.InitialState] = struct{}{}

    for len(queue) > 0 {
        state := queue[0]
        queue = queue[1:]

        if transitions, ok := dfa.Transitions[state]; ok {
            for _, toState := range transitions {
                if _, ok := reachable[toState]; !ok {
                    reachable[toState] = struct{}{}
                    queue = append(queue, toState)
                }
            }
        }
    }

    newDFA := &DFA[T]{
        States:       make(set.Set[T]),
        Alphabet:     dfa.Alphabet,
        Transitions:  make(map[T]map[string]T),
        InitialState: dfa.InitialState,
        FinalStates:  make(set.Set[T]),
    }

    for state := range reachable {
        newDFA.States[state] = struct{}{}
        newDFA.Transitions[state] = make(map[string]T)

        if _, ok := dfa.FinalStates[state]; ok {
            newDFA.FinalStates[state] = struct{}{}
        }
    }

    for state := range reachable {
        if transitions, ok := dfa.Transitions[state]; ok {
            for symbol, toState := range transitions {
                if _, ok := reachable[toState]; ok {
                    newDFA.Transitions[state][symbol] = toState
                }
            }
        }
    }

    return newDFA
}