package main

func plusOne(digits []int) []int {
	var j = len(digits) - 1
	tmp := make([]int, 0)
	for j >= 0 {
		if digits[j]+1 < 10 {
			digits[j] = digits[j] + 1
			break
		} else {
			digits[j] = 0
		}
		j -= 1
	}
	if digits[0] == 0 {
		tmp = append(tmp,1)
		for _, i := range digits{
			tmp = append(tmp, i)
			return tmp
		}
	} else {
		return digits
	}

	return digits
}
