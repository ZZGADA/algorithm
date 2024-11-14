# 排序
1. 快排  [找到topK的元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/)
```java

    public int findKthLargest(int[] nums, int k) {
        quickSort(0, nums.length - 1, nums);
        return nums[nums.length - k];
    }

    public void quickSort(int left, int right, int[] nums) {
        if (left < right) {
            int mid = pivot(left, right, nums);
            quickSort(left, mid - 1, nums);
            quickSort(mid + 1, right, nums);
        }

    }

    public int pivot(int left, int right, int[] nums) {
        int targetIndex = right--;
        int j = left - 1;
        for (int i = left; i <= right; i++) {
            if (nums[i] < nums[targetIndex]) {
                j++;
                swap(nums, i, j);
            }
        }
        swap(nums, ++j, targetIndex);
        return j;
    }

    private void swap(int[] nums, int i, int j) {
        if (i == j) {
            return;
        }
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
```