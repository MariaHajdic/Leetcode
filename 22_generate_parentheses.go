package main

func generateParenthesis(n int) []string {
	result := []string{}
	stack := []byte{}

	var backtrack func(int, int)
	backtrack = func(open, close int) {
		if len(stack) == 2*n { // adding formed pairs
			result = append(result, string(stack))
			return
		}

		if open < n {
			stack = append(stack, '(')
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open {
			stack = append(stack, ')')
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}
