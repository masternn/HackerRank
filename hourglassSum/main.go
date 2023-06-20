package main

import "fmt"

func main() {
	ar := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 9, 2, -4, -4, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, -1, -2, -4, 0},
	}

	fmt.Println(hourglassSum(ar))
}

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	if len(arr) == 0 {
		return 0
	}
	var sum int32
	first := true
	var biggest int32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum += arr[j][i] + arr[j][(i+1)] + arr[j][(i+2)]
			sum += arr[(j + 1)][(i + 1)]
			sum += arr[(j + 2)][i] + arr[(j + 2)][(i+1)] + arr[(j + 2)][(i+2)]
			if first {
				biggest = sum
				first = false
			} else if sum > biggest {
				biggest = sum
			}
			sum = 0
		}
	}
	return biggest
}
