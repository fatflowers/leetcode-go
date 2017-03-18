package main

//import "fmt"

func getColumnMember(matrix [][]int, col, start, end int) []int {
	result := []int{}
	for start < end {
		result = append(result, matrix[start][col])
		start++
	}

	return result
}

func getReversedColumnMember(matrix [][]int, col, start, end int) []int {
	result := []int{}
	for start <= end-1 {
		result = append(result, matrix[end-1][col])
		end--
	}

	return result
}

func getReversedSlice(slice []int) []int {
	result := []int{}

	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}

	return result
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	} else if len(matrix) == 1 {
		return matrix[0]
	} else if len(matrix[0]) == 1 {
		return getColumnMember(matrix, 0, 0, len(matrix))
	}

	w, h := len(matrix[0]), len(matrix)

	iW, iH := 0, 0

	for iW < w/2 && iH < h/2 {
		if iW < w/2 {
			result = append(result, matrix[iH][iW:w-iW]...)
		}

		if iH < h/2 {
			result = append(result, getColumnMember(matrix, w-iW-1, iH+1, h-iH)...)
		}

		if iW < w/2 {
			result = append(result, getReversedSlice(matrix[h-iH-1][iW:w-iW-1])...)
		}

		if iH < h/2 {
			result = append(result, getReversedColumnMember(matrix, iW, iH+1, h-iH-1)...)
		}

		iH++
		iW++
	}

	if w%2 == 1 && iW == w/2 && iH != h/2 {
		result = append(result, getColumnMember(matrix, iW, iH, h-iH)...)
	} else if h%2 == 1 && iH == h/2 && iW != w/2 {
		result = append(result, matrix[iH][iW:w-iW]...)
	} else if h%2 == 1 && w%2 == 1 {
		result = append(result, matrix[h/2][w/2])
	}

	return result
}

//func main() {
//	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
//	fmt.Printf("%v\n", spiralOrder(m))
//
//	m = [][]int{{1, 2}, {3, 4}}
//	fmt.Printf("%v\n", spiralOrder(m))
//	m = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}
//	fmt.Printf("%v\n", spiralOrder(m))
//}
