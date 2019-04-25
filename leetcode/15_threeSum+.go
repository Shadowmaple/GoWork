package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var data [][]int

	for aPositoin, a := range nums {
		for bPosition := len(nums)-1; bPosition >= 0; bPosition-- {
			if bPosition == aPositoin {
				continue
			}
			b := nums[bPosition]
			c := -(a + b)
			flag := false
			for i, n := range nums {
				if i == aPositoin || i == bPosition {
					continue
				}
				if n == c {
					flag = true
					break
				}
			}
			if flag {
				data = append(data, []int{a, b, c})
			}
		}
	}
	return data
}