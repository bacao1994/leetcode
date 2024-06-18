package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	characters := make(map[int32]int32)
	for _, c := range s {
		characters[c]++
	}

	for _, c := range t {
		count := characters[c]
		if count > 0 {
			characters[c]--
		} else {
			return false
		}
	}
	return true

}

//func main() {
//	fmt.Println(isAnagram("anagram", "nagaram"))
//}
