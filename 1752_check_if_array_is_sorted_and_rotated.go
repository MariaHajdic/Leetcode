package main

func check(nums []int) bool {
	countDrop := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		nextIdx := (i + 1) % n // Wrap around with modulo
		if nums[i] > nums[nextIdx] {
			countDrop++
			if countDrop > 1 {
				return false // More than one drop found, not valid
			}
		}
	}

	return true // Exactly one drop was found
}

/* Not working */
// func check(nums []int) bool {
//     prev := nums[0]
//     idx := 1
//     for ; idx < len(nums); idx++ {
//         if nums[idx] < prev {
//             idx++
//             break
//         }
//         prev = nums[idx]
//     }
//     for ; idx < len(nums); idx++ {
//         if nums[idx] < prev {
//             return false
//         }
//         prev = nums[idx]
//         if prev > nums[0] {
//             return false
//         }
//     }

//     return true
// }
