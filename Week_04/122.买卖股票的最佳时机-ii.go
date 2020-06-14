/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var ret int
    for i:=1 ; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}
// @lc code=end

