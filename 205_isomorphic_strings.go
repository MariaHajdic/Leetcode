package main

import "slices"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := map[rune]rune{}
	reverseMapping := map[rune]rune{}

	sRunes := []rune(s)
	tRunes := []rune(t)

	for i, letter := range sRunes {
		if _, ok := mapping[letter]; !ok {
			mapping[letter] = tRunes[i] // map s -> t
			if _, alreadyExists := reverseMapping[mapping[letter]]; alreadyExists {
				return false
			}
			reverseMapping[mapping[letter]] = letter
		}
		sRunes[i] = mapping[letter]
	}

	return slices.Equal(sRunes, tRunes)

	// for i, letter := range s {
	//     if _, ok := mapping[letter]; !ok {
	//         mapping[letter] = rune(t[i]) // map s[i] -> t[i]
	//     }
	//     letter = mapping[letter]
	// }

	/*    var result strings.Builder
	    replaced := make([]rune, len(s))
	    for i, letter := range s {
	        mapped, ok := mapping[letter]
	        if !ok {
	            mapping[letter] = rune(t[i]) // map s[~i] -> t[i]
	            replaced[i] = mapping[letter]
	        } else {
	            replaced[i] = mapped
	        }
	    }

	    // Compare the []rune with t
	    return string(replaced) == t
	}*/

	// for _, letter := range s {}

	// if _, ok := tToS[charT]; ok {

	// stringS := map[rune]int{}
	// stringT := map[rune]int{}

	// for _, letter := range s {
	//     stringS[letter]++
	// }
	// for _, letter := range t {
	//     stringT[letter]++
	// }

	// for i, _ := range s {
	//     if stringS[rune(s[i])] != stringT[rune(t[i])] {
	//         return false
	//     }
	// }

	// return true
}
