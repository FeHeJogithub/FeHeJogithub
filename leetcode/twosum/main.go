package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i, num}
		}
		m[num] = i
	}
	return nil

}

func main() {
	nums := []int{6, 3, 11, 2}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
	fmt.Println(twoSum(nums, target))

}

/*func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}*/
