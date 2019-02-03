package optimization

import (
	"testing"
	"reflect"
	"unsafe"
)

type Anonymous struct {
	D int
}

func normalIncrease(v interface{}) {
	es := reflect.ValueOf(v).Elem()
	e := es.FieldByName("D")
	x := e.Int()
	x++
	e.SetInt(x)
}

func TestNormalIncrease(t *testing.T) {
	anonymous := &Anonymous{}
	normalIncrease(anonymous)
	t.Log(anonymous.D)
}

var cache = make(map[*uintptr]map[string]uintptr)

func ptrIncrease(v interface{}) {
	itab := *(**uintptr)(unsafe.Pointer(&v))
	m, ok := cache[itab]
	if !ok {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	offset, ok := m["D"]
	if !ok {
		t := reflect.TypeOf(v).Elem()
		x, _ := t.FieldByName("D")
		offset = x.Offset
		m["D"] = offset
	}

	p := (*[2]uintptr)(unsafe.Pointer(&v))
	px := (*int)(unsafe.Pointer(p[1]+offset))
	*px++
}

func TestPtrIncrease(t *testing.T) {
	anonymous := &Anonymous{}
	ptrIncrease(anonymous)
	t.Log(anonymous.D)
}

// go test -run=^^$ -bench="^BenchmarkNormalIncrease$" -benchmem
func BenchmarkNormalIncrease(b *testing.B) {
	instance := &Anonymous{}
	for i := 0; i < b.N; i++ {
		normalIncrease(instance)
	}
}

// go test -run=^^$ -bench="^BenchmarkPtrIncrease$" -benchmem
func BenchmarkPtrIncrease(b *testing.B) {
	instance := &Anonymous{}
	for i := 0; i < b.N; i++ {
		ptrIncrease(instance)
	}
}
