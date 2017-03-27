package main

type NumMatrix struct {
    M [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := [][]int{}
    for i := range matrix {
        l := make([]int, len(matrix[i]))
        copy(l, matrix[i])
        m = append(m, l)
    }

    return NumMatrix{m}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    if this == nil || len(this.M) == 0 || len(this.M[0]) == 0 {
        return 0
    }

    if row1 > len(this.M)-1 || col1 > len(this.M[0])-1 || row2 < 0 || col2 < 0 {
        return 0
    }

    hrow1, hcol1, hrow2, hcol2 := row1, col1, row2, col2
    if row1 < 0 {
        hrow1 = 0
    }
    if hcol1 < 0 {
        hcol1 = 0
    }
    if hrow2 > len(this.M)-1 {
        hrow2 = len(this.M)-1
    }
    if hcol2 > len(this.M[0]) {
        hcol2 = len(this.M[0])
    }

    result := 0
    for hrow1 <= hrow2 {
        for c := hcol1; c <= hcol2; c++ {
            result += this.M[hrow1][c]
        }
        hrow1++
    }

    return result
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */