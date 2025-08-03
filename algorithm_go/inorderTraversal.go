package algorithm_go

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

// flatten 将二叉树展开为链表`
func flatten(root *TreeNode) {

}
