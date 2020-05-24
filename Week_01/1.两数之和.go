 /*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	r := make([]int, 0, 2)
	m := make(map[int]int)
	for secondIndex , val := range nums {
		if firstIndex , ok := m[val]; ok {
			return append(r, firstIndex, secondIndex)
		}
		m[target-val] = secondIndex
	}
	return r
}

// @lc code=end
