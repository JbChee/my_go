package main

import "fmt"

func main() {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), raw[0])    // 10000 10000 0xc420080000

	rawSlice := raw[:3]
	fmt.Println(len(rawSlice), cap(rawSlice), &rawSlice[0])

	//修改了原切片，新切片会改变。因为新切片引用的是原切片的数组地址
	raw[0] = 100

	fmt.Println(len(raw), cap(raw), raw[0])
	fmt.Println(len(rawSlice), cap(rawSlice), &rawSlice[0])
	fmt.Println(len(rawSlice), cap(rawSlice), rawSlice[0])


	//往左边切，容量降低
	rawSliceLow := raw[10:]
	fmt.Println(len(rawSliceLow), cap(rawSliceLow), &rawSliceLow[0])


	str := ""
	fmt.Println(str[4:])


}
