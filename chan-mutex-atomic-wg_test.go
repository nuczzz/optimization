package optimization

import (
	"sync"
	"testing"
	"sync/atomic"
)

var chanCount int32

func chanAdd(channel chan struct{}) {
	<-channel
	chanCount++
	channel <- struct{}{}
}

func mutexAdd(mutex sync.Locker) {
	mutex.Lock()
	chanCount++
	mutex.Unlock()
}

func atomicAdd() {
	atomic.AddInt32(&chanCount, 1)
}

func waitGroupAdd(wg *sync.WaitGroup) {
	wg.Add(1)
	chanCount++
	wg.Done()
}

// go test -run=^^$ -bench=^BenchmarkChan$ -benchmem
func BenchmarkChan(b *testing.B) {
	channel := make(chan struct{}, 1)
	channel <- struct{}{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanAdd(channel)
	}
}

// go test -run=^^$ -bench=^BenchmarkMutex$ -benchmem
func BenchmarkMutex(b *testing.B) {
	mutex := &sync.Mutex{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mutexAdd(mutex)
	}
}

// go test -run=^^$ -bench=^BenchmarkAtomic$ -benchmem
func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicAdd()
	}
}

// go test -run=^^$ -bench=^BenchmarkWg$ -benchmem
func BenchmarkWg(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		waitGroupAdd(wg)
	}
}
