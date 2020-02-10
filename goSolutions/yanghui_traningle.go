package main

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	var tmp = []int{}
	for i := 1; i <= numRows; i++ {

		tmp = []int{}
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, res[i-2][j-1]+res[i-2][j])
			}
		}
		res = append(res, tmp)

	}

	return res

}
