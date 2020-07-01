package main

import "fmt"

//超时
func threeSum(nums []int) [][]int {
	var data [][]int
	for index, a := range nums {
		for i := index + 1; i < len(nums)-1; i++ {
			b := nums[i]
			for j := i + 1; j < len(nums); j++ {
				c := nums[j]
				if a+b == -c {
					//排除重复
					flag := true
					for _, k := range data {
						//单独排除[0,0,0]这一可能
						if a == b && a == 0 {
							if k[0] == k[1] && k[0] == k[2] && k[0] == 0 {
								flag = false
								break
							}
							continue
						}
						if contain(a, k) && contain(b, k) && contain(c, k) {
							flag = false
							break
						}
					}
					//若无此三元组则加入
					if flag {
						num := []int{a, b, c}
						data = append(data, num)
					}
				}
			}
		}
	}
	return data
}

//整型数据是否在数组中
func contain(n int, array []int) bool {
	for _, number := range array {
		if n == number {
			return true
		}
	}
	return false
}

func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(num))
}
