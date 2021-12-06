package main

import (
	"fmt"
	"sync"
)


func listfunc() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func slicefunc() {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}



func test_range(){
	var wg sync.WaitGroup
	wg.Add(10)
	var as = []int{1,2,3,4,5,6,7,8,9,10}
	for _,v:=range as {
		go func() {
			defer wg.Done()
			v+=1
			fmt.Println(v)

		}()
	}

	wg.Wait()
}
func test_range1(){
	var wg sync.WaitGroup
	wg.Add(10)
	var as = []int{1,2,3,4,5,6,7,8,9,10}
	for _,v:=range as {
		go func(i int) {
			defer wg.Done()
			i+=1
			fmt.Println(i)

		}(v)
	}

	wg.Wait()
}


func main() {
	//runabc(4,4)
	test_range() // 变量覆盖
	//test_range1()
}
