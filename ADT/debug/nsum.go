package main

import (
	"fmt"
	"sort"
)

//通用nSum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := nSum(nums, 4, 0, 0)
	return res

}
func nSum(nums []int, n, start, target int)[][]int{
	res := [][]int{}
	count := len(nums)
	if n <2 || n > count{
		return res
	}
	if n == 2{
		low, height := start, len(nums)-1
		for low < height{
			lval := nums[low]
			rval := nums[height]
			vsum := lval + rval
			if vsum == target{
				tem := []int{}
				tem = append(tem, nums[low], nums[height])
				res = append(res, tem)
				//去重
				for low < height && lval == nums[low] {
					low++
				}
				for low < height && rval == nums[height]{
					height--

				}
			}else if vsum > target{
				//去重
				for low < height && rval == nums[height]{
					height--

				}
			}else if vsum < target{
				//去重
				for low < height && lval == nums[low] {
					low++
				}
			}
		}

	}else{
		for i := start; i < count; i++{
			//递归
			arr := nSum(nums, n-1, i+1, target - nums[i])
			for _, v := range arr{
				v = append(v,nums[i])
				res = append(res, v)
			}
			for i < count-1 && nums[i] == nums[i+1]{
				i++
			}
		}
	}
	return res
}

//func nSum(nums []int, n, start, target int)[][]int{
//	res := [][]int{}
//	count := len(nums)
//	if n <2 || n > count{
//		return res
//	}
//	if n == 2{
//		low, height := start, len(nums)-1
//		for low < height{
//			lval := nums[low]
//			rval := nums[height]
//			vsum := lval + rval
//			if vsum == target{
//				tem := []int{}
//				tem = append(tem, nums[low], nums[height])
//				res = append(res, tem)
//				//去重
//				for low < height {
//					if lval == nums[low]{
//						low ++
//					}else{
//						break
//					}
//				}
//				for low < height {
//					if rval == nums[height]{
//						height --
//					}else{
//						break
//					}
//
//				}
//			}else if vsum > target{
//				//去重
//				for low < height {
//					if rval == nums[height]{
//						height --
//					}else{
//						break
//					}
//
//				}
//			}else if vsum < target{
//				//去重
//				for low < height {
//					if lval == nums[low]{
//						low ++
//					}else{
//						break
//					}
//
//				}
//			}
//		}
//
//	}else{
//		for i := start; i < count; i++{
//			//递归
//			arr := nSum(nums, n-1, i+1, target - nums[i])
//			for _, v := range arr{
//				v = append(v,nums[i])
//				res = append(res, v)
//			}
//			for i < count-1{
//				if nums[i] == nums[i+1]{
//					i++
//				}else{
//					break
//				}
//
//			}
//		}
//	}
//	return res
//}




func main() {
	nums := []int{2,2,2,2,2}
	res := threeSum(nums)
	fmt.Println(res)

}
