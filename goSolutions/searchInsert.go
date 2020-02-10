package main

import "fmt"

//binary search
func searchInsert(nums []int, target int) int {
	var start = 0
	var end = len(nums)
	var mid = 0
	for start < end {
		mid = (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			end = mid - 1
		}
	}

	return 0
}

func main() {
	var nums = []int{1,3,5,6}
	fmt.Println(searchInsert(nums, 2))
}
