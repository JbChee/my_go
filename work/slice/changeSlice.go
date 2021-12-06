package main

import "fmt"

func main() {
	slice := make([]int, 2, 3)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	fmt.Printf("slice: %v, addr: %p, &slice = %p \n", slice, slice, &slice)

	changeSlice(slice)
	fmt.Printf("slice: %v, addr: %p, &slice = %p \n", slice, slice, &slice)
}

func changeSlice(s []int) {
	s = append(s, 3)
	s = append(s, 4)
	s[1] = 111
	fmt.Printf("func s: %v, addr: %p, &s = %p \n", s, s, &s)
}
