func twoSum(nums []int, target int) []int {
	var result []int
	for i, _ := range nums{
		for j:= i +1;j<len(nums); j = j + 1{
			if nums[i]+nums[j] == target{
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}
