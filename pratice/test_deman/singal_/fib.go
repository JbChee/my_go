package singal_

import "testing"

func Fib(n int) int{

	if n <2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}


func benchmarkFib(b *testing.B, n int){

	for i:=0;i<b.N; i++{
		Fib(n)
	}

}