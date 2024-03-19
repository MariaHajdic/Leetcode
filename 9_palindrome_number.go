package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	newNum, oldNum := 0, x
	for oldNum != 0 {
		digit := oldNum % 10
		newNum = newNum*10 + digit
		oldNum /= 10
	}
	return newNum == x
}
