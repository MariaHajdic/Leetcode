package main

func minRemoveToMakeValid(s string) string {
	result := make([]rune, len(s))
	idx := 0
	// INPUT: r))a(r)a(r(a

	openCount := 0
	// Removing invalid ')' parentheses
	for _, ch := range s {
		if ch == '(' {
			openCount++
		} else if ch == ')' {
			if openCount == 0 {
				continue // skipping an invalid ')'
			}
			openCount--
		}
		result[idx] = ch
		idx++
	} // RES: ra(r)a(r(a // idx = 9

	openCount = 0
	finalIdx := idx
	for i := idx - 1; i >= 0; i-- {
		if result[i] == ')' {
			openCount++
		} else if result[i] == '(' {
			if openCount == 0 {
				continue
			}
			openCount--
		}
		result[finalIdx-1] = result[i]
		finalIdx--
	}

	return string(result[finalIdx:idx])
}
