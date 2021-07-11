package main
import "testing"

func Benchmark_DynamicAllocation(b *testing.B) {
	slice := make([]string, 0)
	for i := 0; i < b.N; i++ {
		slice = append(slice, "Hey Bro!")
	}
	b.Log(len(slice))
}