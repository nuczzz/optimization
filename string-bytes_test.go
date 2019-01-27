package optimization

import (
	"unsafe"
	"testing"
)

func Bytes2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func String2Bytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

var Str string
var Bs []byte

// go test -run=^^$ -bench=^BenchmarkNormalB2S$ -benchmem
func BenchmarkNormalB2S(b *testing.B) {
	bs := []byte{'t', 'e', 's', 't'}
	for i := 0; i < b.N; i++ {
		Str = string(bs)
	}
}

// go test -run=^^$ -bench=^BenchmarkOptimizationB2S$ -benchmem
func BenchmarkOptimizationB2S(b *testing.B) {
	bs := []byte{'t', 'e', 's', 't'}
	for i := 0; i < b.N; i++ {
		Str = Bytes2String(bs)
	}
}

// go test -run=^^$ -bench=^BenchmarkNormalS2B$ -benchmem
func BenchmarkNormalS2B(b *testing.B) {
	str := "test"
	for i := 0; i < b.N; i++ {
		Bs = []byte(str)
	}
}

// go test -run=^^$ -bench=^BenchmarkOptimizationS2B$ -benchmem
func BenchmarkOptimizationS2B(b *testing.B) {
	str := "test"
	for i := 0; i < b.N; i++ {
		Bs = String2Bytes(str)
	}
}
