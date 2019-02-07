package optimization

import "testing"

const capacity = 1000

func initMap(m map[int]int) {
	for i := 0; i < capacity; i++ {
		m[i] = i
	}
}

func normalMap() map[int]int {
	m := make(map[int]int)
	return m
}

func capacityMap() map[int]int {
	m := make(map[int]int, capacity)
	return m
}

func BenchmarkNormalMap(b *testing.B) {
	m:= normalMap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		initMap(m)
	}
}

func BenchmarkCapacityMap(b *testing.B) {
	m:= capacityMap()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		initMap(m)
	}
}
