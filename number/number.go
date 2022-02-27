package number

// 得到组合中的所有正数
func GetPositiveNum(nums []int) []int {
	res := make([]int, 0, len(nums))
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		if nums[i] > 0 {
			res = append(res, nums[i])
		}
	}
	return res
}

// 得到组合中的所有负数
func GetNegativeNum(nums []int) []int {
	res := make([]int, 0, len(nums))
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		if nums[i] < 0 {
			res = append(res, nums[i])
		}
	}
	return res
}

func GetMinVaule(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	n := len(nums) - 1
	for i := 1; i <= n; i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}

func GetMaxValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	n := len(nums) - 1
	for i := 1; i <= n; i++ {
		if nums[i] > res && nums[i] < 100000 {
			res = nums[i]
		}
	}
	return res
}
