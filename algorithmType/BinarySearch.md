# 二分查找
1. 基础写法
```java
int search(int nums[], int size, int target) //nums是数组，size是数组的大小，target是需要查找的值
{
    int left = 0;
    int right = size - 1;	// 定义了target在左闭右闭的区间内，[left, right]
    while (left <= right) {	//当left == right时，区间[left, right]仍然有效
        int middle = left + ((right - left) / 2);//等同于 (left + right) / 2，防止溢出
        if (nums[middle] > target) {
            right = middle - 1;	//target在左区间，所以[left, middle - 1]
        } else if (nums[middle] < target) {
            left = middle + 1;	//target在右区间，所以[middle + 1, right]
        } else {	//既不在左边，也不在右边，那就是找到答案了
            return middle;
        }
    }
    //没有找到目标值
    return -1;
}

```

```java

public static void main(String[] args) {
   int left = -1;
   int right = length - 1; // 开区间 (-1, n-1)
   while (left + 1 < right) { // 开区间不为空
      int mid = left + (right - left) / 2;
      if (nums[mid] < nums[length- 1]) {
         // right 始终保持小于nums[length-1]
         right = mid;
      } else {
         left = mid;
      }
   }
   // right 为最小值
}
```




2. 搜索旋转数组
   [leetcode 33](https://leetcode.cn/problems/search-in-rotated-sorted-array/description/)
```java

class Solution {
    public int search(int[] nums, int target) {
        int length = nums.length;
        int left = 0;
        int right = length - 1; // 开区间 (-1, n-1)
        
        while (left<right){
            int mid = left + (right-left)/2;
            if(nums[mid]<nums[length-1]){
                // 寻找最小元素 维护右边界
                right = mid;
            }else{
                left = mid +1;
            }
        }

        if (target>nums[length-1]){
            // 在左边
            left = 0;
        }else{
            left = right;
            right = length-1;
        }

        while (left<=right){
            int mid = left + (right - left)/2;
            if(nums[mid] == target){
                return mid;
            }else if(nums[mid]>target){
                right = mid -1;
            }else{
                left = mid+1;
            }
        }
        return  -1;
    }
}


class Solution {
   public int search(int[] nums, int target) {
      int lo = 0, hi = nums.length - 1;
      while (lo <= hi) {
         int mid = lo + (hi - lo) / 2;
         if (nums[mid] == target) {
            return mid;
         }

         // 先根据 nums[0] 与 target 的关系判断目标值是在左半段还是右半段
         if (target >= nums[0]) {
            // 目标值在左半段时，若 mid 在右半段，则将 mid 索引的值改成 inf
            if (nums[mid] < nums[0]) {
               nums[mid] = Integer.MAX_VALUE;
            }
         } else {
            // 目标值在右半段时，若 mid 在左半段，则将 mid 索引的值改成 -inf
            if (nums[mid] >= nums[0]) {
               nums[mid] = Integer.MIN_VALUE;
            }
         }

         if (nums[mid] < target) {
            lo = mid + 1;
         } else {
            hi = mid - 1;
         }
      }
      return -1;
   }
}
```

3. 猴子吃桃子 [爱吃香蕉的柯柯](https://leetcode.cn/problems/koko-eating-bananas/description/)
```java
class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        // 吃香蕉的速度k 小于等于piles的最大值 大于多少就不知道了
        // 对于最小的点需要使用二分查找

        int maxNum = 0;
        int minNum = 1;
        int countH = 0;
        for (int i = 0; i < piles.length; i++) {
            maxNum = piles[i] > maxNum ? piles[i] : maxNum;
            minNum = piles[i] < minNum ? piles[i] : minNum;
        }

        // countH == h 时候才会停下来
        // 最坏情况是吃最大max
        int k = maxNum;
        while (minNum < maxNum) {
            countH = 0;
            int eatNum = minNum + (maxNum - minNum) / 2;
            for (int i = 0; i < piles.length; i++) {
                countH += piles[i] % eatNum == 0 ? piles[i] / eatNum : piles[i] / eatNum + 1; // 计算吃的数量
            }
            if (countH > h) {
                // 吃的少了 耗费的时间太多了
                minNum = eatNum + 1;
            } else if (countH <= h) {
                // 吃的太快了 最小速度k不是最小的
                // 找到了的话注意不是最优解 是可行解
                // maxNum = eatNum 向下找 然后下一步算出的值可能会大于h 这样就缩小了区间 
                // 维护右边界 然后找左边的可行解 如果有更小的就更新 
                maxNum = eatNum;
                k = eatNum;
            }
        }
        return k;
    }
}

```



4. 经典二分 [搜索插入位置](https://leetcode.cn/problems/search-insert-position/description/?envType=problem-list-v2&envId=binary-search)
```java

class Solution {
    public int searchInsert(int[] nums, int target) {
        int left = 0;
        int right = nums.length-1;
        int res = 0;

        while(left<right){
            int mid = left + (right-left)/2;
            if(nums[mid]>=target){
                // 维护右边界
                right = mid;
            }else{
                left = mid+1;
            }
        }
        return right==nums.length-1 && nums[right]<target? right+1:right;
    }
}
```