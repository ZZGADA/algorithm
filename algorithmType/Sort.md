# 排序
1. 快排  [找到topK的元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/)
```java

class Solution {
    public int[] sortArray(int[] nums) {
        int left = 0;
        int right = nums.length - 1;

        quickSort(nums, left, right);
        return nums;
    }

    public void quickSort(int[] nums, int left, int right) {
        if (left < right) {
            int pivot = partition(nums,left,right);
            quickSort(nums,left,pivot-1);
            quickSort(nums,pivot+1,right);
        }
    }

    public int partition(int[] nums, int left, int right) {
        int target = nums[right];
        int i, j;

        for (i = left - 1, j = left; j < right; j++) {
            if (nums[j]<target){
                i++;        // 记录 最前面的小于 target的元素
                if (i!=j){
                    swap(nums,i,j);
                }
            }
        }

        //最终 nums[++i]是大于等于target的
        swap(nums,++i,right);
        return i;
    }

    /**
     * 交换
     */
    public void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
        return;
    }
}
```

--- 
2. 快排+随机选择交换 [快排+随机选择交换](https://leetcode.cn/problems/sort-an-array/solutions/178305/pai-xu-shu-zu-by-leetcode-solution/)
随机数的目的是加快排序
```java
class Solution {
    public int[] sortArray(int[] nums) {
        randomizedQuicksort(nums, 0, nums.length - 1);
        return nums;
    }

    public void randomizedQuicksort(int[] nums, int l, int r) {
        if (l < r) {
            int pos = randomizedPartition(nums, l, r);
            randomizedQuicksort(nums, l, pos - 1);
            randomizedQuicksort(nums, pos + 1, r);
        }
    }

    public int randomizedPartition(int[] nums, int l, int r) {
        int i = new Random().nextInt(r - l + 1) + l; // 随机选一个作为我们的主元
        swap(nums, r, i);
        return partition(nums, l, r);
    }

    public int partition(int[] nums, int l, int r) {
        int pivot = nums[r];
        int i = l - 1;
        for (int j = l; j <= r - 1; ++j) {
            if (nums[j] <= pivot) {
                i = i + 1;
                swap(nums, i, j);
            }
        }
        swap(nums, i + 1, r);
        return i + 1;
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
```
--- 
3. topK[topK](https://leetcode.cn/problems/kth-largest-element-in-an-array/?envType=problem-list-v2&envId=sorting)
```java

import java.util.Random;

class Solution {
    private Random random = new Random();

    public int findKthLargest(int[] nums, int k) {
        return topK(nums, 0, nums.length - 1, k);
    }

    private int topK(int[] nums, int left, int right, int k) {
        int pivot, countOfLargerElements;
        while (true) {
            pivot = partition(nums, left, right);
            countOfLargerElements = nums.length - pivot;
            if (countOfLargerElements > k) {
                left = pivot + 1;
            } else if (countOfLargerElements < k) {
                right = pivot - 1;
            } else {
                return nums[pivot];
            }
        }
    }

    private int partition(int[] nums, int left, int right) {
        int r = random.nextInt(right - left + 1) + left;
        swap(nums, r, right);

        int target = nums[right];
        int i = left, j = right - 1;
        while (true) {
            while (i <= j && nums[i] < target) {
                i++;
            }

            while (i <= j && nums[j] > target) {
                j--;
            }
            if (i >= j)break;

            swap(nums,i++,j--);
        }

        swap(nums,i,right);
        return i;
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;

    }
}

```