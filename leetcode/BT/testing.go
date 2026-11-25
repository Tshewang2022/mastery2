package main

import (
	"testing"
	"unsafe"
)

// have to learn more about the heap and stack memory allocation in go
var t *[5]int64
var s []byte

func f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = &[5]int64{}
	}
}

func g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]byte, 32769)
	}
}

func main() {
	println(unsafe.Sizeof(*t)) // 40
	rf := testing.Benchmark(f)
	println(rf.AllocedBytesPerOp()) // 48
	rg := testing.Benchmark(g)
	println(rg.AllocedBytesPerOp()) // 40960
}
