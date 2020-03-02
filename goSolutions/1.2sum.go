package main

// 哈希表一次遍历，时间复杂度为o(n)
func TwoSum(nums []int, target int) []int {
	d := map[int]int{}
	res := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if ok := ContainsKey(d, nums[i]); ok {
			s := IndexOfMap(d, nums[i])
			res = append(res, i)
			res = append(res, s)
			return res
		} else {
			d[target-nums[i]] = i
		}
	}
	return nil
}

func ContainsKey(m map[int]int, key int) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}

	return false
}

func IndexOfMap(m map[int]int, value int) int {
	for k, v := range m {
		if v == value {
			return k
		}
	}

	return 0
}

//暴力解法，两层遍历，时间复杂度为O(n^2)

func TwoSum2(nums []int, target int) []int {
	res := make([]int,0)
	for i :=0;i<len(nums);i++{
		for j := i+1;j<len(nums);j++{
			if target - nums[i] = nums[j]{
				res = append(res,i)
				res = append(res,j)
				return res
			}
		}
	}
	return nil
}
