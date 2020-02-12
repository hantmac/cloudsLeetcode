package main

func plusOne(digits []int) []int {
	for j := len(digits) - 1; j >= 0; j-- {
		digits[j] += 1
		digits[j] = digits[j] % 10
		if digits[j] != 0 {
			return digits
		}
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
