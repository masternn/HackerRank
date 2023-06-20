package main

import "fmt"

func main() {
	ar3 := []int32{2, 1, 5, 3, 4}
	arCaotic := []int32{2, 5, 1, 3, 4}

	minimumBribes(ar3)
	minimumBribes(arCaotic)
}

func minimumBribes(q []int32) {
	// Write your code here
	var bribes int32
	var i, j, d int32
	t := int32(len(q))

	for i = (t - 1); i >= 0; i-- {
		if (q[i] - (i + 1)) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		d = (q[i] - 2)
		if d < 0 {
			d = 0
		}
		for j = d; j <= (i - 1); j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	fmt.Println(bribes)
}
