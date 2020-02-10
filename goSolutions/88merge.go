package main

import (
	"fmt"
)

func merge(nums1 []int,m int, nums2 []int,n int)  []int{
	copyNums1 := nums1
	nums3 := make([]int,0)

	var p1 = 0
	var p2 = 0

	for p1 < m && p2 <n {
		if copyNums1[p1] < nums2[p2] {
			nums3 = append(nums3,copyNums1[p1])
			p1++
		}else {
			nums3 = append(nums3,nums2[p2])
			p2 ++
		}
	}

	if p1 < m {
		for i := p1; i<m;i++{
			nums3 = append(nums3,copyNums1[i])
		}
	}
	if p2 < n {
		for j:= p2;j<n;j++{
			nums3 = append(nums3, nums2[j])
		}
	}
	nums1 = nums3
return nums1
}

func main()  {
	a := []int{1,2,3,0,0,0}
	m := 3
	n := 3
	b := []int{2,4,5}
fmt.Println(merge(a,m,b,n))
}