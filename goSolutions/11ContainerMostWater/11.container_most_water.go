package main

import "fmt"

func MostWater(height []int) int {
	max_contain := 0
	aera := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			aera = (j - i) * min(height[j], height[i])
			max_contain = max(aera, max_contain)
		}
	}
	return max_contain

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	x := MostWater(arr)

	fmt.Println(x)
}

func MostWater2(height []int) int {
	max_contain := 0
	min_height := 0
	aera := 0
	j := len(height) - 1
	i := 0
	for i < j {
		if height[i] < height[j] {
			min_height = height[i]
			i++
		} else {
			min_height = height[j]
			j--
		}
		aera = (j - i + 1) * min_height
		max_contain = max(max_contain, aera)
	}
	return max_contain
}
