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

func io_test(t *testing.T) {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString("Hello, pooled world!")
	fmt.Println(buf.String())
	bufferPool.Put(buf) 
}

func obj_test(t *testing.T) {
	for range 1000000 {
		obj := dataPool.Get().(*Data) 		
		obj.Value = 42                
		dataPool.Put(obj)             
	}
	fmt.Println("Done")
}
