package main

import "fmt"

/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

https://neetcode.io/problems/duplicate-integer/question
*/

func hasDuplicate(nums []int) bool {
    hashmap := make(map[int]int)

    for _, num := range nums {
        if _, ok := hashmap[num]; !ok {
            hashmap[num] = 1
			fmt.Println(hashmap)
        } else {
            return true
        }
    }
    return false
}

func main() {
	fmt.Println(hasDuplicate([]int{3, 4, 3, 1, 2, 6}))
}
