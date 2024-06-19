package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		containerLength := right - left
		area := containerLength * min(height[left], height[right])
		res = max(area, res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
