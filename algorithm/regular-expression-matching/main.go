package main

import "fmt"

func main() {
	fmt.Println(isMatch("ba", "b*b."))
}

func isMatch(s string, p string) bool {
	var dp [][]bool
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(p)+1))
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // 匹配串为空，则只有结尾为'*'才可以一个个抵消
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if p[j-1] == '*' {
					if s[i-1] == p[j-2] || p[j-2] == '.' {
						zero := dp[i][j-2]  // 出现0次
						one := dp[i-1][j-2] // 出现一次
						many := dp[i-1][j]  // 出现多次
						dp[i][j] = zero || one || many
					} else {
						dp[i][j] = dp[i][j-2] // 抵消一次,和上面出现0次一样【s[i-1] != p[j-2]】
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
