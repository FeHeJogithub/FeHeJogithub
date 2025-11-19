package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
