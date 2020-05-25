/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	filter := make(map[int]int)
	filter1 := make(map[int]int)
	for _, v := range nums1 {
		filter[v] = v
	}
	r := make([]int, 0)
	for _, v := range nums2 {
		if vv, ok := filter[v]; ok {
			if _, ok := filter1[vv]; ok {
				continue
			}
			filter1[vv] = vv
			r = append(r, vv)
		}
	}
	return r
}
// @lc code=end

