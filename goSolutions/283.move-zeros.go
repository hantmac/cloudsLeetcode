package main

import "fmt"

func MoveZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count += 1
			fmt.Println(count)
		}
	}
	for i := count; i < len(arr); i++ {
		arr[i] = 0
		fmt.Println(count)
	}
	fmt.Println(arr)
}

func main() {
	nums := []int{0, 1, 3, 0, 12}

	MoveZeros(nums)
}
