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

5. 寻找峰值元素 [寻找峰值元素](https://leetcode.cn/problems/find-peak-element/?envType=problem-list-v2&envId=binary-search)
```java

class Solution {
    public int findPeakElement(int[] nums) {
        // 一定出现一个峰值 则一定存在单调区间 那么根据单调性进行判断
        // 判断元素是否是峰值元素是判断 元素是否都大于左右两边 
        int length = nums.length;
        int left = 0;
        int right = length-1;
        while (left < right) {
            int mid = (right-left)/2 + left;
            if(mid == left){
                // 如果只剩下两个数了 选大的那一个
                return nums[left]>nums[right]?left:right;
            }

            if(nums[mid] > nums[mid-1]){
                // 现在在单调增区间
                if(mid+1<length && nums[mid] > nums[mid+1]){
                    // 找到峰值
                    return mid;
                }else{
                    left = mid+1;
                }
            }else{
                // mid 元素 小于 左边的元素 峰值在左边
                // 在递减区间
                right = mid;
            }
        }

        return left;
    }
}
```

6. 两数之和-输入有序数组 [两数之和-输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/?envType=problem-list-v2&envId=binary-search)
```java
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        // 非递减顺序排列 ---> 增序列 && 有相同元素
        // 注意下标元素从1 开始
        int[] res = new int[2];
        int length = numbers.length;

        for (int i = 0; i < length; i++) {
            if (numbers[i]<0 ||numbers[i] <= target) {
                int left = i + 1;
                int right = length - 1;
                int addTarget = target - numbers[i];

                // 向后寻找符合目标值的元素
                while (left < right) {
                    int mid = (right - left) / 2 + left;
                    if (numbers[mid] >= addTarget) {
                        right = mid;
                    }else{
                        left = mid+1;
                    }
                }

                if(numbers[right] == addTarget){
                    res[0] = i+1;
                    res[1] = right+1;
                    break;
                }
            }else{
                break;
            }

        }

        return res;
    }
}
```
7. [第N位数字](https://leetcode.cn/problems/nth-digit/description/?envType=problem-list-v2&envId=binary-search)
```java
class Solution {
    public int findNthDigit(int n) {
        // 1-9 (10^1 - 10^0) *1
        // 10-99 (10^2 - 10^1) *2
        // 100 - 999 (10^3 - 10^2) *3
        // 1000 - 9999 (10^4 - 10^3) *4

        // 寻找范围区间
        int zhi = 1;
        long numLength = 0, maxEnd = 0;
        while (n > numLength) {
            numLength += (long) (Math.pow(10, zhi) - Math.pow(10, zhi - 1)) * zhi;
            zhi++;
        }
        zhi--; // 区间内数字都是zhi位
        numLength -= (Math.pow(10, zhi) - Math.pow(10, zhi - 1)) * zhi;
        n -= numLength; // 确定区间内的第n个数字


        int flag = (n - 1) / zhi;   //区间内的第flag数
        int mod = (n - 1) % zhi;    // 第flag数的第mod位

        // 确定区间范围
        int start = (int) Math.pow(10, zhi - 1);
        int end = (int) Math.pow(10, zhi) - 1;

        // 开始2分查找  寻找区间内的第flag数
        int left = start;
        int right = end;
        int res = 0;
        while (left <= right) {
            int mid = (right - left) / 2 + left;
            int midflag =  mid - start;
            if (flag < midflag) {
                right = mid - 1;
            } else if (flag > midflag) {
                left = mid + 1;
            } else {
                res = Integer.toString(mid).charAt(mod) - '0';
                break;
            }
        }
        return res;
    }
}
```

8. [我的日程表1](https://leetcode.cn/problems/my-calendar-i/description/?envType=problem-list-v2&envId=binary-search)
```java


class MyCalendar {
    private List<Integer> listStart;
    private List<Integer> listEnd;

    public MyCalendar() {
        listStart = new ArrayList<>();
        listEnd = new ArrayList<>();
        listStart.add(0);
        listEnd.add(0);
    }

    // 可添加到日程 不会导致重复预定 返回true 否则返回false
    // 区间范围是左闭 右开
    public boolean book(int startTime, int endTime) {
        // 偏移量 + 长度
        // 寻找左边第一个节点的起始时间和时长
        boolean res = false;
        int leftFirstIndex = findLeftFirst(startTime);
        int rightFirstIndex = leftFirstIndex + 1;
        // System.out.println(leftFirstIndex);

        int sst = rightFirstIndex >= listStart.size() ? Integer.MAX_VALUE : listStart.get(rightFirstIndex);
        int eet = listEnd.get(leftFirstIndex);
        if (startTime >= eet && endTime <= sst) {
            listStart.add(leftFirstIndex + 1, startTime);
            listEnd.add(leftFirstIndex + 1, endTime);
            res = true;
        }

        // for (int s : listStart) {
        //     System.out.printf("%d ", s);
        // }
        // System.out.println();
        // System.out.println(leftFirstIndex);

        return res;
    }

    // 寻找左边第一个节点的起始时间和时长
    public int findLeftFirst(int startTime) {
        int left = 0;
        int right = listStart.size() - 1;
        int leftFirstIndex = 0;
        while (left < right) {
            int mid = (right - left) / 2 + left;
            if (startTime >= listStart.get(mid)) {
                leftFirstIndex = mid;
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        if (startTime >= listStart.get(right)) {
            leftFirstIndex = right;
        }

        return leftFirstIndex;
    }
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar obj = new MyCalendar();
 * boolean param_1 = obj.book(startTime,endTime);
 */
```

9. [在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)
```java
class Solution {
    int resLeft = Integer.MAX_VALUE;
    int resRight = Integer.MIN_VALUE;

    public int[] searchRange(int[] nums, int target) {
        // 非递减序列
        ll(nums, target, 0, nums.length - 1);

        resLeft = resLeft == Integer.MAX_VALUE ? -1 : resLeft;
        resRight = resRight == Integer.MIN_VALUE ? -1 : resRight;
        return new int[] { resLeft, resRight };
    }

    // 左边找最小的
    public void ll(int[] nums, int target, int left, int right) {
        if (left <= right) {
            int mid = (right - left) / 2 + left;
            if (target < nums[mid]) {
                ll(nums,target,left,mid-1);
            }else if(target == nums[mid]){
                resLeft = Math.min(resLeft,mid);
                resRight = Math.max(resRight,mid);
                ll(nums,target,left,mid-1);
                ll(nums,target,mid+1,right);
            }else{
                ll(nums,target,mid+1,right);
                ll(nums,target,mid+1,right);
            }
        }

    }
}
```

10. [搜索插入位置](https://leetcode.cn/problems/search-insert-position/description/)
```java
class Solution {
    public int searchInsert(int[] nums, int target) {
        // 寻找小于 target的第一个数
        int left = 0, right = nums.length - 1, mid = 0;
        while (left <= right) {
            mid = (right - left) / 2 + left;
            if (target <= nums[mid]) {
                right = mid - 1;
            } else {
                // 寻找小于 target的第一个数
                left = mid + 1;
            }
        }

        return left;
    }
}
```

11.[寻找比目标字母大的最小字母](https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/)
```java
class Solution {
    public char nextGreatestLetter(char[] letters, char target) {
        // 非递减排序
        // 返回大于target的第一个字符
        // 没有一个字符在字典上大于 'z'，所以我们返回 letters[0]
        int left = 0, right = letters.length - 1;
        while (left < right) {
            int mid = (right - left) / 2 + left;
            if (target >= letters[mid]) {
                left = mid + 1;
            } else {
                // mid是大于target的数
                right = mid;
            }
        }

        return letters[right] <= target ? letters[0] : letters[right];
    }
}
```