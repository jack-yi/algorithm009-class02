/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	s, e := 0 , len(nums)
	for s < e {
		if nums[s] == 0 {
			e--
			nums = append(nums[0:s], nums[s+1:]...)
			nums = append(nums, 0)
		}else {
			s++
		}

	}
}
// @lc code=end

