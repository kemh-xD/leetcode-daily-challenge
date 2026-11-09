package array


//Should optimize this solution
//The time complexity of this implementation is O(N*N)
func TwoSum(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//Using a map
//Time complexity here O(N)
func TwoSum2(nums []int, target int) []int{

	var numsMap map[int]int = make(map[int]int)
	for position, value := range nums{
		numsMap[value] = position
	}

	for position, value := range nums{
		complement := target - value
		if j, ok := numsMap[complement]; ok && j != position {
			return []int{position, j}
		}
	}
	return []int{}
}