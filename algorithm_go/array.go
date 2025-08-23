package algorithm_go

import "slices"

// 合并区间
func merge(intervals [][]int) [][]int {
	// 需要先排序
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})

	// intervals 现在具有递增性
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}

	// tmp 保存当前的最大区间
	tmp := []int{intervals[0][0], intervals[0][1]}
	for _, arr := range intervals {
		if arr[0] > tmp[1] {
			// 超出最大区间
			t := make([]int, 2)
			copy(t, tmp)
			res = append(res, t)
			tmp = []int{arr[0], arr[1]}
		} else {
			tmp[1] = max(tmp[1], arr[1])
		}
	}

	// 最后总是剩余一个tmp 待放入
	res = append(res, tmp)

	return res
}
