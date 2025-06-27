package memprealloc

import (
	"testing"
)

const MAX_LEN = 1_000_000

func BenchmarkBadOne(b *testing.B) {
	for b.Loop() {
		items := make([]int, 0)
		for i := range MAX_LEN {
			items = append(items, i*100)
		}
	}
}

func BenchmarkGoodOne(b *testing.B) {
	items := make([]int, MAX_LEN)
	for b.Loop() {
		for i := range MAX_LEN {
			items = append(items, i*100)
		}
	}
}
