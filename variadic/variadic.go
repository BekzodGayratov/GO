package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result

}
