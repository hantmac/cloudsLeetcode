package main

import (
	"fmt"
)

func addBin(x int, y int, flagss int) (sum, addFlag int) {
	sum = (x + y + flagss) % 2
	addFlag = (x + y + flagss) / 2
	return
}

func binstr2slice(binstr string) []int {
	tmp := make([]int, len(binstr), len(binstr))
	for i := 0; i < len(binstr); i++ {
		if binstr[i] == '0' {
			tmp[i] = 0
		} else {
			tmp[i] = 1
		}

	}
	return tmp
}

func binslice2str(binstr []int) string {
	s := ""
	for _, v := range binstr {
		s = s + fmt.Sprint(v)
	}
	return s
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	aInt := binstr2slice(a)
	bInt := binstr2slice(b)

	maxL := len(a)
	addFlag := 0
	sum := make([]int, maxL, maxL)
	for i := 1; i <= maxL; i++ {
		if i <= len(b) {
			sum[len(a)-i], addFlag = addBin(aInt[len(a)-i], bInt[len(a)-i], addFlag)
		} else {
			sum[len(a)-i], addFlag = addBin(aInt[len(a)-i], 0, addFlag)
		}
	}

	if addFlag == 1 {
		sum = append([]int{1}, sum...)
	}

	return binslice2str(sum)
}
