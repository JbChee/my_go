package main

import (
	"context"
)

func Work(ctx context.Context, fn func()){
	fn()
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


func main() {


}
