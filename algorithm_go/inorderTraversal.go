package algorithm_go

import (
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
