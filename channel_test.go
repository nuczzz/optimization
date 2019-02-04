package optimization

import "testing"

const block = 500
const max = 500000
const bufferSize = 100

func normalChannel(data chan int, done chan struct{}) int {
	count := 0
	go func() {
		for x := range data {
			count += x
		}
		close(done)
	}()

	for i := 0; i < max; i++ {
		data <- i
	}

	close(data)
	<-done
	return count
}

func batchChannel(data chan [block]int, done chan struct{}) int {
	count := 0
	go func() {
		for a := range data {
			for _, x := range a {
				count += x
			}
		}
		close(done)
	}()

	for i := 0; i < max; i+=block {
		var b [block]int
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		data <- b
	}

	close(data)
	<-done
	return count
}

// go test -run=^^$ -test="^BenchmarkNormalChannel$" -benchmem
func BenchmarkNormalChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := make(chan int, bufferSize)
		done := make(chan struct{})
		b.StartTimer()

		_ = normalChannel(data, done)
	}
}

// go test -run=^^$ -bench="^BenchmarkBatchChannel$" -benchmem
func BenchmarkBatchChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := make(chan [block]int, bufferSize)
		done := make(chan struct{})
		b.StartTimer()

		_ = batchChannel(data, done)
	}
}
