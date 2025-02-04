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
            // 至少一个到达边界 
            while (i <= j && nums[i] < target) {
                i++;
            }

            while (i < j && nums[j] > target) {
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

top k 问题 最大堆写法   
```java
class Solution {
    public int findKthLargest(int[] nums, int k) {
        return topK(nums, k);
    }

    public int topK(int[] nums, int k) {
        int end = nums.length - 1;

        buildHeap(nums, end); // 最大堆初始化
        for (int i = 0; i < k; i++) {
            swap(nums, 0, end--);
            maxHeapify(nums, 0, end);
        }

        return nums[end + 1];
    }

    // 构建大根堆
    public void buildHeap(int[] nums, int end) {
        for (int i = end / 2; i >= 0; i--) {
            maxHeapify(nums, i, end);
        }
    }

    // 构建大根堆的过程
    // 向父节点进行比较
    public void maxHeapify(int[] nums, int index, int end) {
        int left = index * 2 + 1;
        int right = index * 2 + 2;

        int largeIndex = index;
        if (left > end)
            return;

        // 交换得到当前节点的最大堆
        largeIndex = nums[index] > nums[left] ? index : left;
        largeIndex = right > end ? largeIndex : nums[largeIndex] > nums[right] ? largeIndex : right;
        if (largeIndex == index)
            return;
        swap(nums, largeIndex, index);

        // 递归向下 继续构造最大堆
        maxHeapify(nums, largeIndex, end);
    }

    public void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

}
```

---

4. 堆排序 [堆排序](https://leetcode.cn/problems/sort-an-array/)
```java
class Solution {
    public int[] sortArray(int[] nums) {
        heapSort(nums);
        return nums;
    }

    // 堆排序
    public void heapSort(int[] nums) {
        int len = nums.length - 1;

        // 初始化建堆 
        buildMaxHeap(nums,len);
        for(int i = len;i>=1;--i){
            swap(nums,i,0); // 将堆顶元素 至于最尾 
            len -=1;    // 调整带排序数组的大小 最后一个就不动了
            maxHeapify(nums,0,len); // 向下调整 将最大元素不断置于父节点
        }
    }

    // 初始化 建最大堆 
    // 从第一个非叶子节点开始 然后将最大元素调整到堆顶
    // 每一次swap 都要递归向下 才能保证每一个节点都是 最大堆
    public void buildMaxHeap(int[] nums, int len) {
        // 最后一个非叶子节点的索引是 len/2 -1
        // 从最后一个非叶子节点向前 每一个节点都要遍历 进行调整
        for (int i = len / 2 ; i >= 0; i--) {
            maxHeapify(nums,i,len);
        }
    }

    // 将父节点调整为最大堆 不断向下递推
    public void maxHeapify(int[] nums, int i, int len) {
        while (2 * i + 1 <= len) {
            // 左子树 & 右子树
            int lson = 2 * i + 1;
            int rson = 2 * i + 2;
            int large;

            // 大根堆 选择子树最大元素  判断左子树
            if (lson <= len && nums[lson] > nums[i]) {
                large = lson;
            } else {
                large = i;
            }

            // 大根堆 选择子树最大元素  判断右子树
            // 注意判断的节点 是判断当前三个节点中最大的元素
            if (rson <= len && nums[rson] > nums[large]) {
                large = rson;
            }

            // 交换 得到最大堆
            if (large != i) {
                swap(nums, i, large);
                i = large;
            } else {
                // 无需交换 当前子树就是最大堆 
                break;
            }
        }
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
```

--- 
5. [归并排序](https://leetcode.cn/problems/sort-an-array/solutions/178305/pai-xu-shu-zu-by-leetcode-solution/)
```java
class Solution {
    int[] temp;

    public int[] sortArray(int[] nums) {
        temp = new int[nums.length];
        mergeSort(nums, 0, nums.length - 1);
        return nums;
    }

    public void mergeSort(int[] nums, int left, int right) {
        if (left >= right) {
            return;
        }

        int mid = (right - left) / 2 + left;
        int i = left, j = mid + 1, k = left;
        mergeSort(nums, i, mid);
        mergeSort(nums, j, right);

        while (i <= mid && j <= right) {
            if (nums[i] <= nums[j]) {
                temp[k++] = nums[i++];
            } else {
                temp[k++] = nums[j++];
            }
        }

        // 右侧没有更小的元素
        while(i<=mid){
            temp[k++] = nums[i++];
        }
        while(j<=right){
            temp[k++] = nums[j++];
        }

        for (k = left; k <= right; k++) {
            nums[k] = temp[k];
        }
    }
}
```