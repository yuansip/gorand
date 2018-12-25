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
	selections := pesNew[len(pesNew)-n:]
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

// Slice is a type, typically a collection, that satisfies gorand.RandomSelectNEvenly
type Slice interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// SubSlice returns the sub slice of the collection in range of begin and end.
	SubSlice(begin, end int) Slice
	// At returns the element at index i of the collection
	At(i int) interface{}
}

// Convenience types for common cases

// IntSlice attaches the methods of Slice to []int.
type IntSlice []int

// Len is the number of elements in the slice of ints
func (p IntSlice) Len() int { return len(p) }

// Swap swaps the elements with indexes i and j.
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SubSlice returns the sub slice of the ints in range of begin and end.
func (p IntSlice) SubSlice(begin, end int) Slice { return p[begin:end] }

// At returns the element at index i of the slice of ints
func (p IntSlice) At(i int) interface{} { return p[i] }

// StringSlice attaches the methods of Slice to []string.
type StringSlice []string

// Len is the number of elements in the slice of strings
func (p StringSlice) Len() int { return len(p) }

// Swap swaps the elements with indexes i and j.
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SubSlice returns the sub slice of the strings in range of begin and end.
func (p StringSlice) SubSlice(begin, end int) Slice { return p[begin:end] }

// At returns the element at index i of the slice of strings
func (p StringSlice) At(i int) interface{} { return p[i] }

// Float64Slice attaches the methods of Slice to []float64.
type Float64Slice []float64

// Len is the number of elements in the slice of float64s
func (p Float64Slice) Len() int { return len(p) }

// Swap swaps the elements with indexes i and j.
func (p Float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SubSlice returns the sub slice of the float64s in range of begin and end.
func (p Float64Slice) SubSlice(begin, end int) Slice { return p[begin:end] }

// At returns the element at index i of the slice of float64
func (p Float64Slice) At(i int) interface{} { return p[i] }

// InterfaceSlice attaches the methods of Slice to []interface{}.
type InterfaceSlice []interface{}

// Len is the number of elements in the slice of interfaces
func (p InterfaceSlice) Len() int { return len(p) }

// Swap swaps the elements with indexes i and j.
func (p InterfaceSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// SubSlice returns the sub slice of the interfaces in range of begin and end.
func (p InterfaceSlice) SubSlice(begin, end int) Slice { return p[begin:end] }

// At returns the element at index i of the slice of interfaces
func (p InterfaceSlice) At(i int) interface{} { return p[i] }

// RandomSelectNEvenly selects n elements from slice evenly. The result is selected from elements in place.
func RandomSelectNEvenly(slice Slice, n int) Slice {
	elementsLength := slice.Len()
	if n <= 0 || elementsLength == 0 {
		return nil
	}
	if n >= elementsLength {
		return slice
	}
	candidates := slice
	candidateCount := elementsLength
	for i := 0; i < n; i++ {
		index := randomSelectEvenly(candidates)
		candidateCount--
		candidates.Swap(index, candidateCount)
		candidates = candidates.SubSlice(0, candidateCount)
	}
	return slice.SubSlice(elementsLength-n, elementsLength)
}

// RandomSelectEvenly selects one element from slice evenly
func RandomSelectEvenly(slice Slice) interface{} {
	selections := RandomSelectNEvenly(slice, 1)
	if selections.Len() == 0 {
		return nil
	}
	return selections.At(0)
}

// Convenience wrappers for common cases

// SelectIntEvenly select an int evenly from a slice of ints.
func SelectIntEvenly(slice []int) int {
	return RandomSelectEvenly(IntSlice(slice)).(int)
}

// SelectNIntEvenly select n ints evenly from a slice of ints.
func SelectNIntEvenly(slice []int, n int) []int {
	return RandomSelectNEvenly(IntSlice(slice), n).(IntSlice)
}

// SelectStringEvenly select a string evenly from a slice of strings.
func SelectStringEvenly(slice []string) string {
	return RandomSelectEvenly(StringSlice(slice)).(string)
}

// SelectNStringEvenly select n strings evenly from a slice of strings.
func SelectNStringEvenly(slice []string, n int) []string {
	return RandomSelectNEvenly(StringSlice(slice), n).(StringSlice)
}

// SelectFloat64Evenly select a float64 evenly from a slice of float64s.
func SelectFloat64Evenly(slice []float64) float64 {
	return RandomSelectEvenly(Float64Slice(slice)).(float64)
}

// SelectNFloat64Evenly select n float64s evenly from a slice of float64s.
func SelectNFloat64Evenly(slice []float64, n int) []float64 {
	return RandomSelectNEvenly(Float64Slice(slice), n).(Float64Slice)
}

// SelectInterfaceEvenly select an interface evenly from a slice of interfaces.
func SelectInterfaceEvenly(slice []interface{}) interface{} {
	return RandomSelectEvenly(InterfaceSlice(slice))
}

// SelectNInterfaceEvenly select n interfaces evenly from a slice of interfaces.
func SelectNInterfaceEvenly(slice []interface{}, n int) []interface{} {
	return RandomSelectNEvenly(InterfaceSlice(slice), n).(InterfaceSlice)
}

func randomSelectEvenly(slice Slice) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(slice.Len())
}
