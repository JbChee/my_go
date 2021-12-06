package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}
func bytes2string(b []byte) string{
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
func main() {

	a1 := "123456789"
	b1 := "abcde"
	fmt.Println(string2bytes(a1))
	fmt.Println(string2bytes(b1))


	var a string
	a = "111"

	b := (*int)(unsafe.Pointer(&a))

	c := unsafe.Pointer(&a)

	d := &a

	mp := make(map[int]int,10)
	fmt.Printf("len = %v",len(mp))


	fmt.Printf("a = %#v",a)
	fmt.Printf("b = %#v",b)
	fmt.Printf("c = %#v",c)
	fmt.Printf("d = %#v",d)

	fmt.Println()
	fmt.Printf("%s\n", "hello world")


}
