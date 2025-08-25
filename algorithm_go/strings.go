package algorithm_go

import "strings"

func longestPalindrome(s string) string {
	// 中心扩展
	// 每一个字符都可能是奇数开头的 也可以是偶数开头的
	// 分类讨论

	res := ""
	if len(s) == 0 {
		return res
	}
	res = string(s[0])

	for i := range s {
		// 奇数回文
		left, right := i-1, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
		// 偶数回文
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	return res

}

func reverseWords(s string) string {
	sl := strings.Split(s, " ")

	sb := new(strings.Builder)
	// strings.TrimSpace() 会去除字符串前后的所有空白字符（包括空格、制表符、换行符等）
	for i := len(sl) - 1; i >= 0; i-- {
		if strings.TrimSpace(sl[i]) != "" {
			sb.WriteString(sl[i])
			sb.WriteString(" ")
		}
	}
	res := sb.String()

	if len(res) == 0 {
		return res
	}
	return res[:len(res)-1]
}
