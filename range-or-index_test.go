package optimization

import "testing"

const TraversalMax = 1000

func initList() []string {
	s := make([]string, TraversalMax)
	for i := 0; i < TraversalMax; i++ {
		s[i] = "https://github.com"
	}
	return s
}

func rangeTraversal(list []string) {
	for i, data := range list {
		_, _ = i, data
	}
}

func indexTraversal(list []string) {
	for i := 0; i < len(list); i++ {
		_, _ = i, list[i]
	}
}

// go test -run=^^$ -bench="^BenchmarkRangeTraversal$" -benchmem
func BenchmarkRangeTraversal(b *testing.B) {
	list := initList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rangeTraversal(list)
	}
}

// go test -run=^^$ -bench="^BenchmarkIndexTraversal$" -benchmem
func BenchmarkIndexTraversal(b *testing.B) {
	list := initList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexTraversal(list)
	}
}

// todo: benchmark test result is different from predicted
// todo: to confirm whether some optimization in GO 1.11
