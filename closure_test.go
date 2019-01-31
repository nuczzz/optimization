package optimization

import "testing"

func function(i int) {
	t := i
	t++
}

// go test -run=^^$ -bench=^BenchmarkNormal$ -benchmem
func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		function(i)
	}
}

// go test -run=^^$ -bench=^BenchmarkAnonymous$ -benchmem
func BenchmarkAnonymous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func(i int) {
			t := i
			t++
		}(i)
	}
}

// go test -run=^^$ -bench=^BenchmarkClosure$ -benchmem
func BenchmarkClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			t := i
			t++
		}()
	}
}
