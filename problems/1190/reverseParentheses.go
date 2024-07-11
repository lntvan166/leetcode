// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/?envType=daily-question&envId=2024-07-11
package exercises1190

func ReverseParentheses(s string) string {
	return reverseParentheses(s)
}

func reverseParentheses(s string) string {
	stack := []int{}
	pairs := make(map[int]int)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairs[i], pairs[j] = j, i
		}
	}
	var res []byte
	for i, step := 0, 1; i < len(s); i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pairs[i]
			step = -step
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
