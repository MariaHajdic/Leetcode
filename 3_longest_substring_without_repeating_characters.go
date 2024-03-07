package main

func lengthOfLongestSubstring(s string) int {
	currentUnique := make(map[byte]int, 128) // ascii symbols
	max, currentMax, cutout := 0, 0, 0

	for idx, symbol := range []byte(s) {
		if foundIdx, ok := currentUnique[symbol]; !ok || foundIdx < cutout {
			currentUnique[symbol] = idx
			currentMax++
			continue
		}
		if currentMax > max {
			max = currentMax
		}
		currentMax = idx - currentUnique[symbol]
		cutout = currentUnique[symbol]
		currentUnique[symbol] = idx
	}
	if currentMax > max {
		max = currentMax
	}

	return max
}

// func lengthOfLongestSubstring(s string) int {
//     currentUnique := make(map[byte]int, 128) // ascii symbols
//     max, currentMax := 0, 0

//     for idx, symbol := range []byte(s) {
//         if _, ok := currentUnique[symbol]; !ok {
//             currentUnique[symbol] = idx
//             currentMax++
//             continue
//         }
//         if currentMax > max {
//             max = currentMax
//         }
//         currentMax = idx - currentUnique[symbol]
//         for el := range currentUnique {
//             if currentUnique[el] < currentUnique[symbol] {
//                 delete(currentUnique, el)
//             }
//         }
//         currentUnique[symbol] = idx
//     }
//     if currentMax > max {
//         max = currentMax
//     }

//     return max
// }
