package algorithm_go

import (
	"math"
)

func wordBreakDP(s string, wordDict []string) bool {
	// 当个单词的最长长度 和 最小的长度
	maxLen := 0
	minLen := math.MaxInt
	dict := make(map[string]bool)
	for _, v := range wordDict {
		maxLen = max(maxLen, len(v))
		minLen = min(minLen, len(v))
		dict[v] = true
	}

	// dp[i] 表示前i个元素是否可以拆分
	n := len(s)
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 {
			dp[i] = 1
			continue
		}

		// 获得s[i-1] 开头的字符串切片 并判断是否在dict中
		for j := i - minLen; j >= 0 && j >= j-maxLen; j-- {
			ts := s[j:i]

			// 前j个元素 和  切片是否存在
			if dict[ts] && dp[j] == 1 {
				dp[i] = 1
			}
		}
	}

	return dp[n] == 1
}
