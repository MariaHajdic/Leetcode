package main

func myAtoi(s string) int {
	characters := []byte(s)
	result, idx, negative := 0, 0, 1
	for ; idx < len(characters); idx++ { // pre-digits: space, -, 48(0)
		if characters[idx] == 43 && negative == 1 { // +
			idx++
			break
		}
		if characters[idx] == 45 { // negative
			negative = -1
			idx++
			break
		}
		if characters[idx] != 32 { // skipping whitespace
			break
		}
	}
	for ; idx < len(characters); idx++ { // reading digits
		if characters[idx] < 48 || characters[idx] > 57 {
			break
		}
		if result == 0 && characters[idx] == 48 {
			continue
		}
		result = result*10 + int(characters[idx]) - 48 // ascii
		if result == 2147483647 && negative == -1 {
			return -2147483647
		}
		if result >= 2147483648 && negative == -1 {
			return -2147483648
		}
		if result >= 2147483647 {
			return 2147483647
		}
	}
	return result * negative
}
