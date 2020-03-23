package main

import (
	"fmt"
	"math"
	//"strconv"
)

func main() {

	// Log 2
	fmt.Println(   math.Log2(240000))

	// Binary search
	fmt.Println(binarySearch([]int{1,4,5,6,8,9,10,12}, 7))


}

func binarySearch(list []int, item int) int  {
	low := 0
	high := len(list) - 1


	for low <= high {
		half := (low + high)/2
		guess := list[half]

		if guess == item {
			return half
		}

		if guess > item {
			high = half - 1
		} else {
			low = half + 1
		}
	}

	return -1


}