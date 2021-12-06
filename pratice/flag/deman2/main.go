package main


import (
	"flag"
	"fmt"
)
var cli int

func flag2(){
	fmt.Println("cli",cli)
	p :=new(int8)
	fmt.Println("p",p)

	p1 := new([]int)
	fmt.Println("p1",p1)
	//p[0] = 1

	b := new([8]int)
	fmt.Println("b",b)
	b[0]=1
	fmt.Println("b",b)

	var al []int
	fmt.Println("al",al)

	flag.IntVar(&cli,"cli",88,"just for deman")  //name 名字要和 定义的变量名一致



}

func main() {
	name :=flag.String("name","haha","这是一个描述")
	age :=flag.Int("age",10,"descript")
	index :=flag.Int("index",10,"descript")

	flag2()
	flag.Parse()
	fmt.Println("flag.args:",flag.Args())

	fmt.Println(name,age,index)
	fmt.Println(flag.Lookup("name").Value)
	fmt.Println("age:",*age)
	fmt.Println("cli:",cli)

}
