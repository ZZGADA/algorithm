package algorithm_go

import (
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	// 非空判断
	for root != nil || len(stack) > 0 {
		// 非空判断
		// 节点全部入队
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 抛出节点
		// 获取中序遍历 结点 然后向右遍历
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// generateParenthesis 生成有效括号
func generateParenthesis(n int) []string {
	// 回溯法
	// 右括号在递归的过程中 数量一定比左括号少
	sb := make([]string, 0)
	res := make([]string, 0)
	var dfs func(leftCnt int, rightCnt int)

	dfs = func(leftCnt int, rightCnt int) {

		// 括号数刚刚好一致的时候
		if leftCnt == n && leftCnt == rightCnt {

			res = append(res, strings.Join(sb, ""))
		}

		if leftCnt < n {
			sb = append(sb, "(")

			// 递归 然后回溯
			dfs(leftCnt+1, rightCnt)
			sb = sb[:len(sb)-1]
		}

		// 右边的括号数 必须小于左括号数
		if rightCnt < leftCnt {
			sb = append(sb, ")")

			dfs(leftCnt, rightCnt+1)
			sb = sb[:len(sb)-1]
		}
	}
	dfs(0, 0)

	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	size := len(candidates)
	var dfs func(flag int, cnt int, tmp []int)

	dfs = func(flag int, cnt int, tmp []int) {
		for i := flag; i < size; i++ {
			if cnt == target {
				dst := append([]int(nil), tmp...)
				res = append(res, dst)
				break
			}

			if cnt > target || candidates[i] > target {
				break
			}

			tCnt := cnt + candidates[i]
			if cnt+candidates[i] <= target {
				tmp = append(tmp, candidates[i])
				dfs(i+1, tCnt, tmp)
				if i+1 == size && tCnt == target {
					dst := append([]int(nil), tmp...)
					res = append(res, dst)
				}
				tmp = tmp[:len(tmp)-1]

				x := candidates[i]
				for i < size && candidates[i] == x {
					i++
				}
				i--
			} else {
				break
			}

		}
	}

	dfs(0, 0, []int{})
	return res
}

func combinationSum3(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	size := len(candidates)
	var dfs func(flag int, cnt int, tmp []int)

	// 本质上是一个递归树
	// 然后进行选和不选
	// 选的时候要过滤掉和当前元素一样的情况，因为同一个开头的树可能会有同样的组合
	dfs = func(flag int, cnt int, tmp []int) {
		if cnt == target {
			t := append([]int(nil), tmp...)
			res = append(res, t)
			return
		}

		if cnt > target || flag >= size {
			return
		}

		// 选
		tmp = append(tmp, candidates[flag])
		cnt += candidates[flag]
		dfs(flag+1, cnt, tmp)

		// 回溯
		cnt -= candidates[flag]
		tmp = tmp[:len(tmp)-1]

		// 不选
		flag++
		for flag < size && candidates[flag] == candidates[flag-1] {
			flag++
		}
		dfs(flag, cnt, tmp)
	}

	dfs(0, 0, []int{})
	return res
}
