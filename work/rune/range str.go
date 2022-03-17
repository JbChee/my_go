package main

import "fmt"

func range_str(){
	a := "0Abcdefg\x41\xfe"

	for _, i := range a{
		fmt.Printf("str = %#x\n" , i)

		fmt.Printf("str = %#v" , i)
		fmt.Println()
	}
	fmt.Println()
	for _, j := range []byte(a){

		fmt.Println("byte str = ", j)
	}
}

func space_switch(){
	isSpace := func(char byte) bool {
		switch char {
		case ' ':    // 空格符会直接 break，返回 false // 和其他语言不一样
		// fallthrough    // 返回 true
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t'))    // true
	fmt.Println(isSpace(' '))    // false
}

func main() {

	range_str()

	space_switch()

}
