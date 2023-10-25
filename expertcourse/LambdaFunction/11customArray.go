// https://recursionist.io/dashboard/problems/436
package main

import (
	"fmt"
)

func main() {
	customArray := func(f func(int) int, arr []int) {
		for _, val := range arr {
			fmt.Println(f(val))
		}
	}

	cube := func(n int) int {
		return n * n * n
	}

	splitAndAdd := func(n int) int {
		result := 0
		for n > 0 {
			result += n % 10
			n /= 10
		}
		return result
	}
	customArray(cube, []int{3, 11, 24, 31})
	customArray(splitAndAdd, []int{3, 11, 24, 31})

}
