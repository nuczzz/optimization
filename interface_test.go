package optimization

import "testing"

type Interface interface {
	Add()
}

type Instance struct{
	count int
}

func (i *Instance) Add() {
	i.count++
}

func newInstance() Interface {
	return &Instance{}
}

func normalAdd(instance *Instance) {
	instance.Add()
}

func InterfaceAdd(instance Interface) {
	instance.Add()
}

//go test -run=^^$ -bench=^BenchmarkNormalAdd$ -benchmem
func BenchmarkNormalAdd(b *testing.B) {
	instance := &Instance{}
	for i := 0; i < b.N; i++ {
		normalAdd(instance)
	}
}

//go test -run=^^$ -bench=^BenchmarkInterfaceAdd$ -benchmem
func BenchmarkInterfaceAdd(b *testing.B) {
	instance := newInstance()
	for i := 0; i < b.N; i++ {
		InterfaceAdd(instance)
	}
}
