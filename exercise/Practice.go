package exercise

import "fmt"

func Practice() {
	nums := []int{1, 2, 3, 4}
	newNums := nums[:len(nums)-1]
	nums = append(nums, 5)
	fmt.Println(nums)
	fmt.Println(newNums)
}
