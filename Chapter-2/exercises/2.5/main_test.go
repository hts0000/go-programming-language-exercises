package main

import "testing"

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(99199999)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(99199999)
	}
}
