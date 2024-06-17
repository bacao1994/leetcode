package main

func containsDuplicate(nums []int) bool {
	dupMap := make(map[int]bool)

	for _, num := range nums {
		if dupMap[num] {
			return true
		} else {
			dupMap[num] = true
		}
	}
	return false
}

//func main() {
//	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
//}
