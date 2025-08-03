package algorithm_go

import "sort"

// ThreeSum 三数和
// 请找出所有和为0 且不元素不重复的三元组
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums) // 对数组进行排序
	size := len(nums)
	res := make([][]int, 0)

	for i := 0; i < size-2 && nums[i] <= 0; i++ {
		target := -nums[i]
		j := i + 1
		k := size - 1

		for j < k {
			tmpSum := nums[j] + nums[k]
			if tmpSum < target {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if tmpSum > target {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}

		for i < size-2 {
			if nums[i+1] == nums[i] {
				i++
			} else {
				break
			}
		}
	}
	return res

}
