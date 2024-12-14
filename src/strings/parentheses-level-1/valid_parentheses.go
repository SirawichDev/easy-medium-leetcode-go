package main

import "fmt"

func main() {
	fmt.Println(ValidParentheses("("))
	fmt.Println(ValidParentheses("[({})]"))
	fmt.Println(ValidParentheses("("))
}

func ValidParentheses(str string) bool {
	var stack []string
	parenthesesMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	for i := 0; i < len(str); i++ {
		currentChar := string(str[i])
		if parenthesesMap[currentChar] != "" {
			stack = append(stack, string(parenthesesMap[currentChar]))
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != currentChar {
				return false
			}
			stack = stack[:len(stack)-1] // : mean from start to end of stack - 1

		}
	}

	return len(stack) == 0
}
