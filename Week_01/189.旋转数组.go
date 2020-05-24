/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
   var temp, pre int
	k = k % len(nums)
	for i:= 0 ; i < k ; i++ {
		pre = nums[len(nums) - 1]
		for j := 0; j < len(nums); j++ {
			temp = nums[j]
			nums[j] = pre
			pre = temp
		}
	}
}
// @lc code=end

