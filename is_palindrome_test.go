package main

import "testing"

func BenchmarkIsPalindrome1(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !isPalindrome1(123456789987654321) {
			b.Fail()
		}
	}
	b.StopTimer()
}
