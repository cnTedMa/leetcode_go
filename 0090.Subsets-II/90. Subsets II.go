package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var c []int
	var res [][]int
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(c, nums, &res, k, 0)
	}
	return res
}

func generateSubsetsWithDup(c, nums []int, res *[][]int, k, start int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(c, nums, res, k, i+1)
		c = c[:len(c)-1]
	}
	return
}
