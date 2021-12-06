package main

import "fmt"

//求并集
func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func main() {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	un := union(slice1, slice2)
	fmt.Println("slice1与slice2的并集为：", un)
	in := intersect(slice1, slice2)
	fmt.Println("slice1与slice2的交集为：", in)
	di := difference(slice1, slice2)
	fmt.Println("slice1与slice2的差集为：", di)

	di1 := difference(slice2, slice1)
	fmt.Println("slice2与slice1的差集为：", di1)
}
