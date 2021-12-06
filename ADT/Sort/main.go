package main

import (
	"fmt"

)

//冒泡排序
func bubblersort(list []int) []int{
	n := len(list)
	didswap :=false
	for i:=n-1;i>0;i--{
		for j :=0;j<i;j++{
			if list[j]> list[j+1]{
				list[j],list[j+1] = list[j+1],list[j]
				didswap = true
			}
		}
		if !didswap{
		}
	}
	return list
}


//选择排序
func selectsort(list []int) []int{
	n := len(list)

	for i :=0;i<n-1;i++{
		mindata := list[i]
		minindex := i
		for j:=i+1;j<n;j++{
			if mindata > list[j]{
				mindata = list[j]
				minindex = j
			}
		}
		if i !=minindex{
			list[i],list[minindex] = list[minindex],list[i]
		}
	}
	return list
}


//选择排序优化
func selectgoodsort(list []int) []int{
	n := len(list)

	for i:=0; i<n/2; i++ {
		minindex := i
		maxindex := i
		for j := i + 1; j < n-i; j++ {
			//找到最大值
			if list[j] > list[maxindex] {
				maxindex = j
				continue
			}
			//找到最小值
			if list[j] < list[minindex] {
				minindex = j
			}
		}
		//最大值在最后，最小值在最前
		if maxindex == i && minindex == n-i-1 {
			//互相交换
			list[maxindex], list[minindex] = list[minindex], list[maxindex]

			////}else if maxindex == i && minindex != n-i-1{//最大值在最前面
			//	//交换最大值  和尾部
			//	list[n-i-1], list[maxindex] = list[maxindex], list[n-i-1]
			//	//交换最小值  和头部
			//	list[i], list[minindex] = list[minindex], list[i]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minindex] = list[minindex], list[i]
			list[n-i-1], list[maxindex] = list[maxindex], list[n-i-1]
		}
	}


	return list



}

//插入排序
func insertsort(list []int) []int{
	n := len(list)
	for i :=1; i<=n-1; i++{
		deal := list[i]  //待排序
		j := i-1    //待排序左边

		if deal < list[j]{
			for j>=0 && deal < list[j]{
				list[j+1] = list[j]
				j -=1   //关键
			}
			list[j+1] = deal
		}
	}
	return list
}


//希尔排序
func shellsort(list []int) []int{
	n := len(list)

	//步长
	for step :=n/2; step>=1; step /= 2{
		for i:=step; i<n; i+=step{
			for j := i-step; j>=0; j -= step{
				if list[j+step] < list[j]{
					list[j+step],list[j] = list[j],list[j+step]
					continue
				}
				//break
			}
		}

	}

	return list
}

//快速排序
func quicksort(array []int,begin,end int){


	if begin <end{
		loc := quick(array,begin,end)

		quicksort(array, begin,loc-1)
		quicksort(array,loc+1,end)
	}
}

//优化快速排序递归   伪尾递归快速排序
func QuickSort3(array []int, begin, end int) {
	for begin < end {
		// 进行切分
		loc := quick(array, begin, end)

		// 那边元素少先排哪边
		if loc-begin < end-loc {
			// 先排左边
			QuickSort3(array, begin, loc-1)
			begin = loc + 1
		} else {
			// 先排右边
			QuickSort3(array, loc+1, end)
			end = loc - 1
		}
	}
}


func quick(array []int,begin , end int ) int{
	i := begin +1
	j := end

	for i<j {
		//比基准大的，都放右边
		if array[i] > array[begin]{
			array[i], array[j] = array[j],array[i]   //交换位置
			j--
		}else{
			i++
		}
	}
	//重合
	if array[j] >= array[begin]{
		i--
	}
	array[i],array[begin] = array[begin],array[i]


	return i
}

//归并排序
func mergesort(array []int, begin ,end int)  {
	//begin := 0
	//end := len(array)

	if end-begin >1{



		mid := begin + (end-begin+1)/2

		mergesort (array, begin, mid)
		mergesort (array , mid , end)

		merge(array , begin, mid, end)
	}

	//return array

}

func merge(array []int, begin , mid , end int){
	fmt.Println("begin",begin)
	fmt.Println("mid",mid)

	leftsize := mid - begin
	rightsize := end - mid
	newsize := leftsize+rightsize

	result := make([]int, 0 , newsize)
	fmt.Println("result", result)

	//指针
	l, r := 0,0
	for l < leftsize && r < rightsize {
		lvalue := array[begin+l]
		rvalue := array[mid+r]
		if lvalue < rvalue {
			result = append(result, lvalue)
			l++
		} else {
			result = append(result, rvalue)
			r++
		}
	}
		//剩余部分
	fmt.Println("l , r", l,r)
	result = append(result, array[begin+l : mid]...)
	result = append(result, array[mid+r : end]...)

	for i :=0; i < newsize; i++{
		array[begin+i] = result[i]
	}
	return

}


func main() {
	//data
	bubbler_list :=[]int{4,9,2,8,165,654,8,898,12,89,87,63,123,55,415}
	select_list :=[]int{4,9,2,8,165,654,8,898,12,89,87,63,123,55,415}
	selectgood_list :=[]int{4,9,2,8,165,1,654,8,898,12,89,87,63,123,55,415,99999}
	shellsort_list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	insertsort_list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}

	quicksort_list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}

	mergesort_list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}

	//mergesort_list := []int{5, 9, 6,3}



	//函数
	bubbler_ := bubblersort(bubbler_list)
	select_ := selectsort(select_list)
	selectgood_ := selectgoodsort(selectgood_list)
	shellsort_ := shellsort(shellsort_list)
	insertsort_ := insertsort(insertsort_list)

	//mergesort_ := mergesort(mergesort_list)


	quicksort(quicksort_list,0,len(quicksort_list)-1)
	mergesort(mergesort_list,0,len(mergesort_list))

	//查看结果
	fmt.Println(bubbler_)
	fmt.Println(select_)
	fmt.Println(selectgood_)
	fmt.Println(shellsort_)
	fmt.Println(insertsort_)

	fmt.Println(quicksort_list)

	fmt.Println(mergesort_list)
}