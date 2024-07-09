package main

import "strings"

func numberToWords(num int) string {
	divisors := []struct {
		value int
		name  string
	}{
		{1000000000, "Billion"},
		{1000000, "Million"},
		{1000, "Thousand"},
	}

	res := []string{}

	for _, divisor := range divisors {
		if num/divisor.value > 0 {
			res = append(res, parse_triplet(num/divisor.value), divisor.name)
			num %= divisor.value
		}
	}
	if num > 0 {
		res = append(res, parse_triplet(num))
	} else if num == 0 && len(res) == 0 {
		return "Zero"
	}

	return strings.Join(res, " ")
}

func parse_triplet(num int) string {
	res := []string{}

	if num >= 100 {
		res = append(res, get_digit(num/100), "Hundred")
		num %= 100
	}
	if num >= 20 {
		res = append(res, get_tens(num))
		num %= 10
	} else if num >= 10 {
		res = append(res, get_teens(num))
		num = 0
	}
	if num > 0 {
		res = append(res, get_digit(num))
	}

	return strings.Join(res, " ")
}

func get_tens(digits int) string {
	digits /= 10
	switch digits {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 8:
		return "Eighty"
	default:
		return strings.Join([]string{get_digit(digits), "ty"}, "")
	}
}

func get_teens(digits int) string {
	switch digits {
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 15:
		return "Fifteen"
	case 18:
		return "Eighteen"
	default:
		return strings.Join([]string{get_digit(digits % 10), "teen"}, "")
	}
}

func get_digit(digit int) string {
	switch digit {
	case 9:
		return "Nine"
	case 8:
		return "Eight"
	case 7:
		return "Seven"
	case 6:
		return "Six"
	case 5:
		return "Five"
	case 4:
		return "Four"
	case 3:
		return "Three"
	case 2:
		return "Two"
	case 1:
		return "One"
	default:
		return "Zero"
	}
}
