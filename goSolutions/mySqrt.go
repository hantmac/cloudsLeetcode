package main

func mySqrt(x int) int {
	if x < 0 {
		return 0
	}

	start := 0
	end := x / 2 +1
	mid := 0
	sq := 0
	for start < end {
		mid = (start + end +1) /2
		sq = mid * mid
		if sq > x {
			end = mid -1
		}else {
			start = mid
		}
	}
	return start
}
