package main

import "fmt"

func run(){

	a := 100

	//101
	//defer func() {
	//	fmt.Println(a)
	//}()

	//100
	//defer func(a int) {
	//	fmt.Println(a)
	//}(a)

	//100
	defer fmt.Println(a)
	a++




}


func main(){

	run()
}
