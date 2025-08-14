package algorithm_go

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	memo := make([]int8, n+1)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int8
	dfs = func(i int) (res int8) {
		if i == 0 { // 成功拆分！
			return 1
		}
		p := &memo[i]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(n) == 1
}

func wordBreakSelf(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	memo := make([]int8, n+1)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}

	// 从后向前递归
	// dfs(i) 表示能否把前缀 s[:i]（表示 s[0] 到 s[i−1] 这段子串）划分成若干段，使得每段都在 wordDict 中
	var dfs func(int) int8

	dfs = func(i int) (res int8) {
		if i == 0 {
			res = 1
			return
		}

		// 表示记忆过
		if memo[i] != -1 {
			res = memo[i]
			return
		}

		// 向前遍历
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				memo[i] = 1
				res = 1
				return
			}
		}

		memo[i] = 0
		return
	}

	// 0-n
	return dfs(n) == 1
}
