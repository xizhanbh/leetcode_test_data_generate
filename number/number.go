package number

func getPositiveNum(nums []int) []int{
	res := make([]int,0, len(nums))
	n := len(nums)-1
	for i := 0; i<= n;i++ {
		if nums[i] > 0 {
			res = append(res, nums[i])
		}
	}
	return res
}

func getNegativeNum(nums []int) []int{
	res := make([]int,0, len(nums))
	n := len(nums)-1
	for i := 0; i<= n;i++ {
		if nums[i] < 0 {
			res = append(res, nums[i])
		}
	}
	return res
}
