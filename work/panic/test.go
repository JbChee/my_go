package main

import (
	"fmt"
	"time"
)


//一个协程panic之后，是会导致所有的协程全部挂掉的，程序会整体退出
//最外层进程捕获recover也是不行的，只能在当前panic的协程，recover
func RunPanic(UserName string){

	//defer func() {
	//	if e := recover(); e != nil{
	//		fmt.Println("recover_panic")
	//	}
	//}()

	if UserName == ""{
		go panic("username = ''")
	}


}


func main() {
	defer func() {
		if e := recover(); e != nil{
			fmt.Println("recover_panic")
		}
	}()
	panic("")
	//go RunPanic("")

	time.Sleep(10*time.Second)
}
