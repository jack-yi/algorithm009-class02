/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	r := make([]int, 0, len(digits)+1)
	for i:= len(digits) - 1; i >= 0 ; i-- {
		digits[i] += 1
		if digits[i] <= 9 {
			return digits
		}else {
			digits[i] = 0
			if i == 0 {
				r = append(r, 1)
				return append(r, digits...)
			}
		}
	}
	return r
}
// @lc code=end

