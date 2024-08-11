package main

func isValid(s string) bool {
	characters := []byte(s)
	stack := make([]byte, len(characters))
	iter := 0

	for i := 0; i < len(characters); i++ {
		c := characters[i]
		switch c {
		case '(', '[', '{':
			stack[iter] = c
			iter++
		case ')':
			if iter == 0 || stack[iter-1] != '(' {
				return false
			}
			iter--
		case ']':
			if iter == 0 || stack[iter-1] != '[' {
				return false
			}
			iter--
		case '}':
			if iter == 0 || stack[iter-1] != '{' {
				return false
			}
			iter--
		}
	}
	return iter == 0
}

// func isValid(s string) bool {
//     characters := []byte(s)
//     stack := make([]byte, len(characters))
//     iter := 0

//     // Map to check matching brackets
//     matchingBrackets := map[byte]byte{
//         ')': '(',
//         ']': '[',
//         '}': '{',
//     }

//     for i := 0; i < len(characters); i++ {
//         c := characters[i]
//         switch c {
//         case '(', '[', '{':
//             stack[iter] = c
//             iter++
//         case ')', ']', '}':
//             if iter == 0 || stack[iter-1] != matchingBrackets[c] {
//                 return false
//             }
//             iter--
//         }
//     }

//     return iter == 0
// }
