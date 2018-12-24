// Package gorand provides functions to select element from slice randomly according to specific element probability distribution.
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
func RandomSelectN(pes []ProbabilityElement, n int) []ProbabilityElement {
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
	selections := make([]ProbabilityElement, 0, n)
	for i := 0; i < n; i++ {
		index := randomSelect(pesNew, totalProbability)
		selections = append(selections, pesNew[index])
		totalProbability -= pesNew[index].GetProbability()
		candidateCount--
		pesNew[index], pesNew[candidateCount] = pesNew[candidateCount], pesNew[index]
		pesNew = pesNew[:candidateCount]
	}
	sort.Slice(selections, func(i, j int) bool {
		return selections[i].GetProbability() > selections[j].GetProbability()
	})
	return selections
}

// RandomSelect selects one element from slice randomly according to the elements' probability distribution
func RandomSelect(pes []ProbabilityElement) ProbabilityElement {
	selections := RandomSelectN(pes, 1)
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
