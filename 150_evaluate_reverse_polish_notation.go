package main

import "strconv"

func evalRPN(tokens []string) int {
	numbers := []int{}

	for _, t := range tokens {
		number, err := strconv.Atoi(t)
		if err == nil { // not operation
			numbers = append(numbers, number)
			continue
		}

		last, prev_last := numbers[len(numbers)-2], numbers[len(numbers)-1]
		switch t {
		case "+":
			numbers[len(numbers)-2] = last + prev_last
		case "-":
			numbers[len(numbers)-2] = last - prev_last
		case "*":
			numbers[len(numbers)-2] = last * prev_last
		case "/":
			numbers[len(numbers)-2] = last / prev_last
		default: // shouldn't happen
		}
		numbers = numbers[:len(numbers)-1]
	}

	return numbers[0]
}
