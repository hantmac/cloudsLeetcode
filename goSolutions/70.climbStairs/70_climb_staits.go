package main

func ClimbStatis(n) int {
	a := make([]int, 0)
	res := 0
	if n <= 2 {
		return n
	}
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 2)
	for i := 3; i <= n; i++ {

		a = append(a, a[i-1]+a[i-2])
		res = a[i]
	}
	return res
}
