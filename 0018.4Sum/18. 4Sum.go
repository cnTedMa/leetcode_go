package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	var uniqNums []int
	var res [][]int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		if uniqNums[i]*4 == target && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if uniqNums[i]*3+uniqNums[j] == target && counter[uniqNums[i]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if uniqNums[i]*2+uniqNums[j]*2 == target && counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			if uniqNums[i]+uniqNums[j]*3 == target && counter[uniqNums[j]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target && counter[uniqNums[i]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target && counter[uniqNums[j]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target && counter[uniqNums[k]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return res
}
