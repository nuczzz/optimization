package optimization

import "testing"

const cap = 1024

func initArray() [cap]int {
	var array [cap]int
	for i := 0; i < cap; i++ {
		array[i] = i
	}
	return array
}

func initSlice() []int {
	var slice = make([]int, cap)
	for i := 0; i < cap; i++ {
		slice[i] = i
	}
	return slice
}

// go test -run=^^$ -bench=^BenchmarkArray$ -benchmem
func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = initArray()
	}
}

// go test -run=^^$ -bench=^BenchmarkSlice$ -benchmem
func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = initSlice()
	}
}
