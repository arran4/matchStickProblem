package main

import "testing"

func BenchmarkIsANumber(b *testing.B) {
	input := []bool{
		// 1
		false,
		false, true,
		false,
		false, true,
		false,
		// 2
		true,
		false, true,
		true,
		true, false,
		true,
		// 3
		true,
		false, true,
		true,
		false, true,
		true,
		// 11
		false,
		true, true,
		false,
		true, true,
		false,
		// 8
		true,
		true, true,
		true,
		true, true,
		true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isANumber(input)
	}
}
