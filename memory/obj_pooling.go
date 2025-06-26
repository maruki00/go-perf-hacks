package memory

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

type Data struct {
	Value int
}

var dataPool = sync.Pool{
	New: func() any {
		return &Data{}
	},
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func BenchmarkIO(b *testing.B) {
	for b.Loop() {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		buf.WriteString("Hello, pooled world!")
		fmt.Println(buf.String())
		bufferPool.Put(buf)
	}
}

func BenchmarkObj(b *testing.B) {
	for b.Loop() {
		obj := dataPool.Get().(*Data)
		obj.Value = 42
		dataPool.Put(obj)
	}
	fmt.Println("Done")
}
