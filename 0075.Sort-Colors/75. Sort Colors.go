package leetcode

func sortColors(nums []int)  {
	zero, one, two := 0, 0, len(nums)-1
	for one <= two {
		if nums[one] == 1 {
			one++
		} else if nums[one] == 0 {
			tmp := nums[zero]
			nums[zero] = 0
			nums[one] = tmp
			one++
			zero++
		} else {
			tmp := nums[two]
			nums[two] = nums[one]
			nums[one] = tmp
			two--
		}
	}
}
