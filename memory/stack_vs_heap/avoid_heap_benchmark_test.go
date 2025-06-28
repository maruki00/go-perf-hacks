package stackvsheap

import "testing"



func BenchmarkBad1(b *testing.B) {
	for b.Loop(){
		Bad1()
	}
}
func BenchmarkBad2(b *testing.B) {
	for b.Loop(){
		Bad1()
	}
}
func BenchmarkBad3(b *testing.B) {
	for b.Loop(){
		Bad1()
	}
}
func BenchmarGoodk(b *testing.B) {
	for b.Loop(){
		Good()
	}
}

