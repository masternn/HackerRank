package main

import "fmt"

func main() {
	ar := []int32{1, 3, 5, 2, 4, 6, 7}

	fmt.Println(minimumSwaps(ar))
}

func minimumSwaps(arr []int32) int32 {
	var swaps, t, i int32
	tot := int32(len(arr))

	for i < tot {
		if (i + 1) != arr[i] {
			if arr[i] > arr[(arr[i]-1)] || (i+1) > arr[i] {
				t = arr[i]
				arr[i] = arr[(arr[i] - 1)]
				arr[(t - 1)] = t
				swaps++
				i -= 1
			}
		}
		i++
	}
	return swaps
}
