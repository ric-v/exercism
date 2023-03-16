package main

import "testing"

// benchmark Reverse
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("cool")
	}
}
