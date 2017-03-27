package main
import "testing"

func TestSumRegion(t *testing.T) {
    obj := Constructor([][]int{
        {3, 0, 1, 4, 2},
        {5, 6, 3, 2, 1},
        {1, 2, 0, 1, 5},
        {4, 1, 0, 1, 7},
        {1, 0, 3, 0, 5},
    })
    t.Log(len(obj.M))
    t.Log(obj.M)
    t.Log(obj.SumRegion(2, 1, 4, 3))
    t.Log(obj.SumRegion(1, 1, 2, 2))
    t.Log(obj.SumRegion(1, 2, 2, 4))
}