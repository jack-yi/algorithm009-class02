/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]bool)
	var counter int
	for k , v:= range nums {
		if counter > len(nums)/2 {
			return nums[k-1]
		}
		if m[v] {
			continue
		}
		counter = 0
		m[v] = true
		for _, vv := range nums {
			if v == vv {
				counter++
			}
		}
	}
	return 0
}

// @lc code=end

