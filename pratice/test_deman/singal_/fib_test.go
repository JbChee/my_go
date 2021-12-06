package singal_

import "testing"

//func benchmarkFib(b *testing.B, n int){
//
//	for i:=0;i<b.N; i++{
//		Fib(n)
//	}
//
//}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b,1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b,4)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b,8)
}


