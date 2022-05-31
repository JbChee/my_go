package main

import "fmt"

type testData struct {
	Cv string `json:"cv"`
	Uids []int64 `json:"uids"`
}
func run(){

	a := 100

	data := testData{}
	data.Cv = "xxxx"

	//101
	defer func() {
		fmt.Println(a)
		fmt.Println(data.Uids)
	}()

	//100
	//defer func(a int) {
	//	fmt.Println(a)
	//}(a)

	//100
	//defer fmt.Println(a)
	a++
	data.Uids = []int64{100232,100264}




}


func main(){

	run()
}
