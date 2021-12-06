package main

import "fmt"

func double(x *int) {
	*x += *x
	x = nil
	fmt.Printf("x: %#v\n", x)
}
//
//func double1(x *int) {
//
//	x += x
//	x = nil
//}

type Data struct {
	Uid int64 `json:"uid"`
}

func changeStruct (d *Data){
	d.Uid = 100232

	d = nil
}


func main() {
	var a = 3
	fmt.Printf("a: %#v\n", a)
	double(&a)
	fmt.Println(a) // 6

	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false

	//tem := Data{}
	//
	//changeStruct(&tem)
	//
	//fmt.Println(tem.Uid)




}
