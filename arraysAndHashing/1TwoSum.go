package main

func twoSum(nums []int, target int) []int {
	mapNumbers := make(map[int]int)

	for i, num := range nums {
		if _, has := mapNumbers[target-num]; has {
			return []int{i, mapNumbers[target-num]}
		} else {
			mapNumbers[num] = i
		}
	}
	return []int{}
}
