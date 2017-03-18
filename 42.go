package main

// func reverseSlice(height []int) []int {
//     r := make([]int, len(height))
//     for i := range height {
//         r[i] = height[len(height)-i-1]
//     }

//     return r
// }

func trap(height []int) int {
    if len(height)<2 {
        return 0
    }

    result := 0
    
    for i:=0; i < len(height); {
        if (height[i] == 0) {
            i++
            continue
        }
        area := 0
        j := i+1
        for ; j < len(height); j++ {
            if height[j] >= height[i] {
                result += (j-i-1) * height[i] - area
                i = j
                break
            } else {
                area += height[j]
            }
        }
        if j == len(height) {
            return result + trap(reverseSlice(height[i:]))
        }
    }

    return result
}

// func main() {
//     println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
//     println(trap([]int{9,8,2,6}))
// }
