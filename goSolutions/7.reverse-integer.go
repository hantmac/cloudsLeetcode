/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.22%)
 * Total Accepted:    643.4K
 * Total Submissions: 2.6M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
package main

import (
	"fmt"
)

func reverse(x int) int {
	var min = 0x80000000
	var max = 0x7fffffff
	var tmp = 0

	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > max || tmp > min {
		return 0
	}
	return tmp
}

func main() {
	a := 1
	fmt.Println(reverse(a))
}
