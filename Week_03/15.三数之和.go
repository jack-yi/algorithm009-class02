/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {  return [][]int{} }
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i:=0; i < len(nums); i++ {
		l, r := i+1, len(nums) - 1
		if l < len(nums)  && (nums[i] > 0 || nums[i] + nums[l] > 0) { break }
		if i > 0 && nums[i-1] == nums[i] { continue }

		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++ 
				continue
			}
			if  r < len(nums)-2 && nums[r+1] == nums[r] {
				r-- 
				continue
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			}else if sum < 0 {
				l++
			}else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return ret
}
// @lc code=end

