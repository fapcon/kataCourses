package main

import (
	"fmt"
	"testing"
)

func BenchmarkHashMap_CasheSlice(b *testing.B) {
	m := NewHashMap()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CasheSlice()
	}
}
func BenchmarkHashMap_CasheList(b *testing.B) {
	m := NewHashMap()
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.CasheList()
	}
}
