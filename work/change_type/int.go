package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type People struct {
	Name string
	Age  int
}

func (s People) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

//func (s *People) String() string {
//	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
//}

func main() {

	//类型转换
	var i int = 9
	var f float64
	f = float64(i)
	fmt.Printf("%T, %v\n", f, f)

	f = 10.8
	a := int(f)
	fmt.Printf("%T, %v\n", a, a)

	//断言类型

	var j interface{} = new(Student)
	s := j.(*Student)

	fmt.Println(s)

	//自定义打印字符串

	var p = &People{
		Name: "qcrao",
		Age:  18,
	}
	//实现string方法
	fmt.Println(p)
}
