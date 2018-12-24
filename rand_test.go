package gorand

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Selectable struct {
	Value       int
	Probability float64
}

func (s Selectable) GetValue() interface{} {
	return s.Value
}

func (s Selectable) GetProbability() float64 {
	return s.Probability
}

func TestRandomSelect(t *testing.T) {
	elementCount := 4
	elements := make([]ProbabilityElement, 0, elementCount)
	totalProbability := 0.0
	for i := 0; i < elementCount; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		f := r.Float64()
		elements = append(elements, Selectable{
			Value:       i,
			Probability: f,
		})
		totalProbability += f
	}
	elementMap := make(map[int]float64)
	for i := 0; i < elementCount; i++ {
		elementMap[i] = elements[i].GetProbability() / totalProbability
	}

	loopCount := 100000
	mp := make(map[int]int)

	for i := 0; i < loopCount; i++ {
		selection := RandomSelect(elements)
		mp[selection.GetValue().(int)]++
	}

	for k, v := range mp {
		assert.InDelta(t, elementMap[k], float64(v)/float64(loopCount), 0.02)
	}
}

func doBenchmark(b *testing.B, n int) {
	elementCount := n
	elements := make([]ProbabilityElement, 0, elementCount)
	for i := 0; i < elementCount; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		f := r.Float64()
		elements = append(elements, Selectable{
			Value:       i,
			Probability: f,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandomSelect(elements)
	}
}

func BenchmarkRandomSelect100(b *testing.B) {
	doBenchmark(b, 100)
}

func BenchmarkRandomSelect1000(b *testing.B) {
	doBenchmark(b, 1000)
}

func BenchmarkRandomSelect10000(b *testing.B) {
	doBenchmark(b, 10000)
}
