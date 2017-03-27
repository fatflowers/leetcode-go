package main

// type NumArray struct {
//     L []int
// }


// func Constructor(nums []int) NumArray {
//     l := make([]int, len(nums))
//     copy(l, nums)
//     return NumArray{l}
// }


// func (this *NumArray) Update(i int, val int)  {
//     if i < 0 || i > len(this.L)-1 {
//         return
//     }

//     this.L[i] = val
// }


// func (this *NumArray) SumRange(i int, j int) int {
//     start, end := i, j    
//     if start > len(this.L) - 1 || end < 0 {
//         return 0
//     }

//     if start < 0 {
//         start = 0
//     }    
//     if end > len(this.L) - 1 {
//         end = len(this.L) - 1
//     }

//     result := 0
//     for start <= end {
//         result += this.L[start]
//         start++
//     }

//     return result
// }


// /**
//  * Your NumArray object will be instantiated and called as such:
//  * obj := Constructor(nums);
//  * obj.Update(i,val);
//  * param_2 := obj.SumRange(i,j);
//  */