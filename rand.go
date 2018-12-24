// Package gorand provides functions to select element from slice randomly.
package gorand

import (
	"math/rand"
	"sort"
	"time"
)

// ProbabilityElement is a interface that can be random selected from slice
type ProbabilityElement interface {
	GetValue() interface{}
	GetProbability() float64
}

// RandomSelectN selects n elements from slice randomly according to the elements' probability distribution
// sortSelections indicates whether to sort the output selections by probability
func RandomSelectN(pes []ProbabilityElement, n int, sortSelections bool) []ProbabilityElement {
	if n <= 0 || len(pes) == 0 {
		return nil
	}
	pesNew := make([]ProbabilityElement, len(pes))
	totalProbability := 0.0
	candidateCount := 0
	for i := range pes {
		pesNew[i] = pes[i]
		if pes[i].GetProbability() > 0 {
			totalProbability += pes[i].GetProbability()
			candidateCount++
		}
	}
	if candidateCount == 0 {
		return nil
	}
	sort.Slice(pesNew, func(i, j int) bool {
		return pesNew[i].GetProbability() > pesNew[j].GetProbability()
	})
	pesNew = pesNew[:candidateCount]
	if n >= len(pesNew) {
		return pesNew
	}
	candidates := pesNew
	for i := 0; i < n; i++ {
		index := randomSelect(candidates, totalProbability)
		totalProbability -= candidates[index].GetProbability()
		candidateCount--
		candidates[index], candidates[candidateCount] = candidates[candidateCount], candidates[index]
		candidates = candidates[:candidateCount]
	}
	selections := pesNew[len(pesNew)-n : len(pesNew)]
	if sortSelections {
		sort.Slice(selections, func(i, j int) bool {
			return selections[i].GetProbability() > selections[j].GetProbability()
		})
	}
	return selections
}

// RandomSelect selects one element from slice randomly according to the elements' probability distribution
func RandomSelect(pes []ProbabilityElement) ProbabilityElement {
	selections := RandomSelectN(pes, 1, false)
	if len(selections) == 0 {
		return nil
	}
	return selections[0]
}

func randomSelect(pes []ProbabilityElement, totalProbability float64) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	f := r.Float64() * totalProbability
	accumulatedProbability := 0.0
	for i := range pes {
		accumulatedProbability += pes[i].GetProbability()
		if accumulatedProbability >= f {
			return i
		}
	}
	return 0
}

// RandomSelectNEvenly selects n elements from slice evenly.
func RandomSelectNEvenly(elements []interface{}, n int) []interface{} {
	elementsLength := len(elements)
	if n <= 0 || elementsLength == 0 {
		return nil
	}
	if n >= elementsLength {
		return elements
	}
	candidates := elements
	candidateCount := elementsLength
	for i := 0; i < n; i++ {
		index := randomSelectEvenly(candidates)
		candidateCount--
		candidates[index], candidates[candidateCount] = candidates[candidateCount], candidates[index]
		candidates = candidates[:candidateCount]
	}
	return elements[elementsLength-n : elementsLength]
}

// RandomSelectEvenly selects one element from slice evenly
func RandomSelectEvenly(elements []interface{}) interface{} {
	selections := RandomSelectNEvenly(elements, 1)
	if len(selections) == 0 {
		return nil
	}
	return selections[0]
}

func randomSelectEvenly(elements []interface{}) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len(elements))
}
