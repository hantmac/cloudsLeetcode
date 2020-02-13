package main

import "fmt"

func threeSum(nums []int) [][]int {
	tmp := make([][]int, 0)
	s := 0
	nums = sort(nums)
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for i < j {
			s = nums[i] + nums[j] + nums[k]
			if s > 0 {
				j -= 1
				for i < j && nums[j] == nums[j+1] {
					j -= 1
				}
			}
			if s < 0 {
				i += 1
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
			}
			if s == 0 {
				t := []int{nums[i], nums[j], nums[k]}
				tmp = append(tmp, t)
				i += 1
				j -= 1
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
				for i < j && nums[j] == nums[j+1] {
					j -= 1
				}
			}
		}
	}
	fmt.Println(tmp)
	return tmp
}

func sort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	te := []int{-1, 0, 1, 2, -1, -4}

	threeSum(te)

}
