package main

func isValid(s string) bool {
	var stack []uint8

	for i := 0; i < len(s); i++ {
		if s[i] == ')' && len(stack) > 0 {
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && len(stack) > 0 {
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && len(stack) > 0 {
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}
