package main

import (
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
} // sync.Pool of Person

//go:noinline
func inc(p *Person) { p.Age++ }

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Age = 1
			b.StopTimer()
			inc(p)
			b.StartTimer()
			personPool.Put(p)
		}
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = &Person{Age: 1}
			b.StopTimer()
			inc(p)
			b.StartTimer()
		}
	}
}

/*
func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
for j := 0; j < 10000; j++ {
p = personPool.Get().(*Person)
p.Age = 1
b.StopTimer(); inc(p); b.StartTimer()
personPool.Put(p)
}
}
}

*/
