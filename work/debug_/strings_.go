package main

import "fmt"

type TT float64


func(t TT)String()string{
	return fmt.Sprintf("run...")
}

func main() {

	var t TT

	res := t.String()
	fmt.Println(res)

}
