package main

import "strconv"

func compareVersion(version1 string, version2 string) int {
    a1, err := strconv.ParseFloat(version1, 32)
    if err != nil {
        println(err.Error())
        return 0
    }

    a2, err := strconv.ParseFloat(version2, 32)
    if err != nil {
        println(err.Error())
        return 0
    }

    if a1 > a2 {
        return 1
    } else if a1 == a2 {
        return 0
    } else {
        return -1
    }
}