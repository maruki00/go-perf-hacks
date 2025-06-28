package io

import "testing"




func BenchmarkBad(b *testing.B) {
	for b.Loop(){
		Bad()
	}
}


func BenchmarkGood(b *testing.B) {
	for b.Loop(){
		Good()
	}
}
