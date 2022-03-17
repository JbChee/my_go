package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Work(ctx context.Context, fn func()){
	fn()
	g := GoID()
	fmt.Println(g)
}

func init(){
	Work(context.TODO(), func()func()func()func(){
		return func()func()func(){
			return func()func(){
				return func(){
				}
			}
		}
	}()()())
}
func GoID() int {

	var buf [64]byte

	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))

	// 得到id字符串

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]

	id, err := strconv.Atoi(idField)

	if err != nil {

		panic(fmt.Sprintf("cannot get goroutine id: %v", err))

	}

	return id

}

func main() {


}
