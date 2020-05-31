/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	r := make([]string, 0)
	for i:= 1; i <=n ; i++ {
		if (i/3 != 0 && i%3 == 0) && (i / 5 != 0 && i%5==0) {
			r = append(r, "FizzBuzz")
		}else if i / 5 != 0 && i%5==0 {
			r = append(r, "Buzz")
		}else if i/3 != 0 && i%3 == 0 {
			r = append(r, "Fizz")
		}else {
			r = append(r, strconv.Itoa(i))
		}
	}
	return r
}
// @lc code=end
 