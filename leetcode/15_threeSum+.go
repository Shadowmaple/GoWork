package main

import (
	"sort"
)

//执行用时 : 1472 ms, 在3Sum的Go提交中击败了77.63% 的用户
//内存消耗 : 82.4 MB, 在3Sum的Go提交中击败了45.48% 的用户
func threeSum(nums []int) [][]int {
	var data [][]int
	sort.Ints(nums)
	length := len(nums)

	for aIndex, a := range nums {
		if a > 0 {
			break
		}
		if aIndex > 0 && nums[aIndex] == nums[aIndex-1] {
			continue
		}
		bIndex := aIndex + 1
		cIndex := length - 1
		for ; bIndex < cIndex; {
			if a + nums[bIndex] == -nums[cIndex] {
				data = append(data, []int{a, nums[bIndex], nums[cIndex]})
				for bIndex < cIndex && nums[bIndex] == nums[bIndex+1] {
					bIndex++
				}
				for bIndex < cIndex && nums[cIndex] == nums[cIndex-1] {
					cIndex--
				}
				bIndex++; cIndex--
				continue
			}
			if a + nums[bIndex] < -nums[cIndex] {
				bIndex++
			} else {
				cIndex--
			}
		}
	}
	return data
}