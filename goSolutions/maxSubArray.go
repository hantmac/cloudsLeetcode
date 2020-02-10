package main

import (
	"fmt"
	"math"
	"runtime"
)

func maxSubArray(nums []int) int {
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + int(math.Max(float64(nums[i-1]), 0))
	}

	for _, v := range nums {
		if res < v {
			res = v
		}
	}
	return res
}

var quit chan int = make(chan int)

func loop(id int) { // id: 该goroutine的标号
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(3) // 最多同时使用2个核

	for i := 0; i < 3; i++ { //开三个goroutine
		go loop(i)
	}

	for i := 0; i < 3; i++ {
		<- quit
	}
}


