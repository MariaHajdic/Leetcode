package main

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{} // key = num, value = it's idx
	for i, num := range nums {
		if _, ok := numsMap[num]; ok {
			if num*2 == target {
				return []int{numsMap[num], i}
			}
			continue
		}
		numsMap[num] = i
	}
	for num := range numsMap {
		num2 := target - num
		if idx, ok := numsMap[num2]; ok && numsMap[num] != numsMap[num2] {
			return []int{numsMap[num], idx}
		}
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
//     numsMap := map[int]int{}
//     for i, _ := range nums {
//         if _, ok := numsMap[nums[i]]; ok {
//             if nums[i] * 2 == target {
//                 return []int{numsMap[nums[i]], i}
//             }
//             continue
//         }
//         numsMap[nums[i]] = i
//     }
//     for num := range numsMap {
//         num2 := target - num
//         delete(numsMap, num)
//         idx, ok := numsMap[num2]
//         if ok {
//             return []int{numsMap[num], idx}
//         }
//     }
//     return nil
// }
