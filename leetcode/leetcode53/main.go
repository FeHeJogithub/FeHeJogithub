package main

import "fmt"

func maxSubArray(nums []int) (int, int, int) {
	currentSum := nums[0]
	maxSum := nums[0]

	start, end := 0, 0 // límites del mejor subarray
	tempStart := 0     // inicio temporal

	for i := 1; i < len(nums); i++ {
		// Si nums[i] es mejor que currentSum + nums[i], empezamos nuevo subarray
		if nums[i] > currentSum+nums[i] {
			currentSum = nums[i]
			tempStart = i
		} else {
			currentSum += nums[i]
		}

		// Si encontramos una suma mejor, actualizamos maxSum y el rango
		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}

	return maxSum, start, end
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum, start, end := maxSubArray(nums)

	fmt.Println("Array:", nums)
	fmt.Println("Máxima suma:", maxSum)
	fmt.Println("Subarray que la produce:", nums[start:end+1])
}
