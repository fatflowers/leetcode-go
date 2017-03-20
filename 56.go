package main

import "sort"


type Interval struct {
    Start int
    End   int
}

func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}


type SortInterval []Interval

func (a SortInterval) Len() int           { return len(a) }
func (a SortInterval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortInterval) Less(i, j int) bool { 
	if a[i].Start < a[j].Start {
		return true
	} else if a[i].Start == a[j].Start {
		return a[i].End < a[j].Start
	}

	return false
}

func overlap(a, b *Interval) bool {
	return b.Start <= a.End
}

func mergeI(a, b *Interval) *Interval {
	result := &Interval{a.Start, maxInt(a.End, b.End)}
	return result
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    result, input := []Interval{}, make([]Interval, len(intervals))

    copy(input, intervals)

    sort.Sort(SortInterval(input))

    for i := 0; i < len(input); {
    	tmp := &input[i]
    	j := i+1
    	for ; j < len(input); j++ {
    		if overlap(tmp, &input[j]) {
    			tmp = mergeI(tmp, &input[j])
    		} else {
    			result = append(result, *tmp)
    			i = j
    			break
    		}
    	}
    	if j == len(input) {
    		result = append(result, *tmp)
    		break
    	}
    }

    return result
}





