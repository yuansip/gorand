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
		assert.InDelta(t, elementMap[k], float64(v)/float64(loopCount), 0.03)
	}
}

func doBenchmark(b *testing.B, n, m int) {
	elementCount := m
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
		RandomSelectN(elements, n, false)
	}
}

func BenchmarkRandomSelect1_100(b *testing.B) {
	doBenchmark(b, 1, 100)
}

func BenchmarkRandomSelect1_1000(b *testing.B) {
	doBenchmark(b, 1, 1000)
}

func BenchmarkRandomSelect1_10000(b *testing.B) {
	doBenchmark(b, 1, 10000)
}

func BenchmarkRandomSelect50_100(b *testing.B) {
	doBenchmark(b, 50, 100)
}

func BenchmarkRandomSelect500_1000(b *testing.B) {
	doBenchmark(b, 500, 1000)
}

func BenchmarkRandomSelect5000_10000(b *testing.B) {
	doBenchmark(b, 5000, 10000)
}
func BenchmarkRandomSelect99_100(b *testing.B) {
	doBenchmark(b, 99, 100)
}

func BenchmarkRandomSelect999_1000(b *testing.B) {
	doBenchmark(b, 999, 1000)
}

func BenchmarkRandomSelect9999_10000(b *testing.B) {
	doBenchmark(b, 9999, 10000)
}

func TestRandomSelectEvenly(t *testing.T) {
	elementCount := 10
	elements := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		elements[i] = i
	}

	loopCount := 100000
	mp := make(map[int]int)

	for i := 0; i < loopCount; i++ {
		selection := SelectIntEvenly(elements)
		mp[selection]++
	}

	expectProbability := float64(1) / float64(elementCount)
	for _, v := range mp {
		assert.InDelta(t, expectProbability, float64(v)/float64(loopCount), 0.02)
	}
}

func doEvenlyBenchmark(b *testing.B, n, m int) {
	elementCount := m
	elements := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectNIntEvenly(elements, n)
	}
}

func BenchmarkRandomSelectEvenly1_100(b *testing.B) {
	doEvenlyBenchmark(b, 1, 100)
}

func BenchmarkRandomSelectEvenly1_1000(b *testing.B) {
	doEvenlyBenchmark(b, 1, 1000)
}

func BenchmarkRandomSelectEvenly1_10000(b *testing.B) {
	doEvenlyBenchmark(b, 1, 10000)
}

func BenchmarkRandomSelectEvenly50_100(b *testing.B) {
	doEvenlyBenchmark(b, 50, 100)
}

func BenchmarkRandomSelectEvenly500_1000(b *testing.B) {
	doEvenlyBenchmark(b, 500, 1000)
}

func BenchmarkRandomSelectEvenly5000_10000(b *testing.B) {
	doEvenlyBenchmark(b, 5000, 10000)
}
func BenchmarkRandomSelectEvenly99_100(b *testing.B) {
	doEvenlyBenchmark(b, 99, 100)
}

func BenchmarkRandomSelectEvenly999_1000(b *testing.B) {
	doEvenlyBenchmark(b, 999, 1000)
}

func BenchmarkRandomSelectEvenly9999_10000(b *testing.B) {
	doEvenlyBenchmark(b, 9999, 10000)
}
