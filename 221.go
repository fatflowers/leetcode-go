package main

func minInt(x, y, z int) int {
	var min int
	if x < y {
		min = x
	} else {
		min = y
	}

	if min < z {
		return min
	} else {
		return z
	}
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maximal := 0
	markRow := make([]int, len(matrix[0]))
	for i := range matrix[0] {
		markRow[i] = int(matrix[0][i])
		if markRow[i] == 1 {
			maximal = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		tmpRow := make([]int, len(matrix[0]))
		tmpRow[len(matrix[0]) - 1] = int(matrix[i][len(matrix[0]) - 1])
		for j := len(matrix[0]) - 2; j >= 0; j-- {
			if matrix[i][j] == 0 {
				tmpRow[j] = 0
			} else {
				tmpRow[j] = minInt(tmpRow[j+1], markRow[j+1], markRow[j]) + 1
				if maximal < tmpRow[j] {
					maximal = tmpRow[j]
				}
			}
		}
		copy(markRow, tmpRow)
	}

	return maximal * maximal
}

//func main() {
//	println(maximalSquare([][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}))
//	println(maximalSquare([][]byte{"10100","10111","11111","10010"}))
//	println(maximalSquare([][]byte{{1, 1}, {1, 0}}))
//}