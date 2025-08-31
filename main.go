package main

import (
	"sort"
	"strings"
)

func main() {

	s := "pwwkew"

	// 记录区间出现的字符串
	rem := make(map[rune]bool)
	left := 0
	res := 0

	for _, rv := range s {

		if rem[rv] {
			// 字符串元素存在
			// 需要缩小滑动窗口
			for !rem[rune(s[left])] {
				delete(rem, rune(s[left]))
				left++
			}
			left++
		} else {
			// 否则增大
			rem[rv] = true
		}

		res = max(res, len(rem))

	}
}

func simplifyPath(path string) string {
	// path 绝对符合unix风格
	left, right := 0, 0
	size := len(path)
	stack := make([]string, 0)
	for right < size {
		// 抵达分隔符号
		if path[right:right+1] == "/" {
			middle := path[left:right]
			switch middle {
			case "..":
				if len(stack) > 0 {
					// 抛出元素
					stack = stack[:len(stack)-1]
				}
			case "/", ".", "":
				// 直接跳过
			default:
				// 入栈 && 指针滑动
				stack = append(stack, path[left:right])
			}
			left = right + 1
		}
		right++
	}

	// right  抵达最后一位进行单独处理
	middle := path[left:right]
	switch middle {
	case "..":
		if len(stack) > 0 {
			// 抛出元素
			stack = stack[:len(stack)-1]
		}
	case "/", ".", "":
		// 直接跳过
	default:
		// 入栈 && 指针滑动
		stack = append(stack, path[left:right])
	}

	res := strings.Join(stack, "/")

	return "/" + res
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	prefix = strs[0]

	for _, s := range strs {
		// 提前终止
		if prefix == "" {
			break
			var i int
			// 比对截至目前为止的最长公共前缀
			for i = 0; i < len(prefix); i++ {
				if prefix[i] != s[i] {
					break
				}
			}
			prefix = prefix[:i]
		}

		return prefix
	}
}
