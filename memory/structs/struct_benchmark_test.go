package structs

import (
	"math"
	"testing"
)


func BenchmarkBadStruct(b *testing.B){
	for b.Loop(){
	items := make([]BadStruct, 1_000_000)
		for i,j := range items {
			j.Age = (i*1_000_000)% math.MaxInt
		}
	}
}



func BenchmarkGoodStruct(b *testing.B){
	for b.Loop(){
	items := make([]GoodStruct, 1_000_000)
		for i,j := range items {
			j.Age = (i*1_000_000)% math.MaxInt
		}
	}
}
