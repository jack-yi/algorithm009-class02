/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	return dfs(0, num, num)
}

func dfs(i, j , target int) bool {
	if i > j { 
		return false
	}
	fmt.Println(i,j)
	num := i + ((j-i)/2)
	sqrt := num * num 
	if sqrt == target {
		return true 
	}else if sqrt > target {
		return dfs(i, num-1, target)
	}else {
		return dfs(num+1, j,target)
	}
}
// @lc code=end

