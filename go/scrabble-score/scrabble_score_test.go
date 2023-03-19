package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Score(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

// BenchmarkScore-8         1088914               919.1 ns/op            80 B/op          9 allocs/op
// BenchmarkScore-8         1777758               681.8 ns/op            80 B/op          9 allocs/op
// BenchmarkScore-8          961029              1101 ns/op              80 B/op          9 allocs/op
func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode. long benchmark")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func BenchmarkScoreParallel(b *testing.B) {
	if testing.Short() {
		// long sentence benchmark
		b.Skip("skipping benchmark in short mode. The quick brown fox jumps over the lazy dog")
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range scrabbleScoreTests {
				Score(test.input)
			}
		}
	})
}
