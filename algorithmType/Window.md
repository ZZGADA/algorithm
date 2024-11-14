# 滑动窗口
**核心思想**：
1. 窗口右指针右边移动是找可行解
2. 窗口左指针有移优化可行解
3. 最终对这个窗口范围找一个最优解 需要对多个窗口进行比较

**额外说明**：
1.  因为双指针其实也是利用了二段性质，当一个指针确定在某个位置，另外一个指针能够落在某个明确的分割点，使得左半部分满足，右半部分不满足。

#### 题目

1. 回文字符串（中心扩散）
```java
class Solution {
    public String longestPalindrome(String s) {
        int length = s.length();
        String resStr = s.substring(0, 1);
        int resLen = 0;

        for (int i = 0; i < length; i++) {
            // 区分奇数和偶数 
            // 此种所有返回的结果是不重复的子字符串
            String jiStr = partion(i - 1, i + 1, s);
            String ouStr = partion(i, i + 1, s);

            String tempRes = jiStr.length() > ouStr.length() ? jiStr : ouStr;
            if (resLen < tempRes.length()) {
                resLen = tempRes.length();
                resStr = tempRes;
            }
        }
        return resStr;

    }

    public String portion(int left, int right, String s) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            // 如果两个元素相等的情况就要向两边进行扩散
            left--;
            right++;
        }
        return s.substring(left + 1, right);
    }
}
````

--- 

2. 三数和 [leetcode 三数和](https://leetcode.cn/problems/3sum/)
```java
/**
 * 1. 排序 从小到大
 * 2. 一次遍历寻找起始位置 如果起始位置与上一个状态相同 那么就跳过
 * 3. 找到不重复的起始位置之后 建立滑窗 寻找int sum = nums[left] + nums[right] + nums[i] == 0
 * 4. 如果不等于0 判断大小 小于0 说明nums[left]小了 否则nums[right]大了
 * 5. 注意在第四步中 left++ right-- 对重复元素进行过滤 
 * @param nums
 * @return
 */
public List<List<Integer>> threeSum(int[] nums) {
   // 排序
   Arrays.sort(nums);
   List<List<Integer>> res = new ArrayList<>();
   int len = nums.length;
   int lenFirst = len - 2;

   for (int i = 0; i < lenFirst; i++) {
      // 初始元素>0 三数一定不可能等于0
      if (nums[i] > 0) {
         break;
      }
      // 初始值去重
      // 与上一个状态匹配 如果相同 那么就++ 因为上一个状态已经处理过了
      while (i < lenFirst && i > 0 && nums[i] == nums[i - 1]) {
         i++;
      }

      int left = i + 1, right = len - 1;

      while (left < right) {
         int sum = nums[left] + nums[right] + nums[i];
         if (sum == 0) {
            // 找到相等的后 继续寻找
            // 注意去重 防止状体重复
            res.add(Arrays.asList(nums[i], nums[left], nums[right]));
            while (left < right && nums[left] == nums[++left])
               ;
            while (left < right && nums[right] == nums[--right])
               ;
         } else if (sum < 0) {
            // 表示小了
            // 去重
            while (left < right && nums[left] == nums[++left])
               ;
         } else {
            while (left < right && nums[right] == nums[--right])
               ;
         }
      }
   }

   return res;
}

```

--- 

3. 最多k个重复元素的最长子串 [最多k个重复元素的最长子串](https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/description/)
```java
class Solution {
    // 滑动窗口 两个map
    public int maxSubarrayLength(int[] nums, int k) {
        // 最后一次出现的位置 + 统计出现的次数
        HashMap<Integer, Integer> mapCount = new HashMap<>();
        int i = 0;
        int res = 0;

        for (int j = 0; j < nums.length; j++) {
            // 如果小于 窗口扩大
            if (mapCount.getOrDefault(nums[j], 0) < k) {
                mapCount.put(nums[j], mapCount.getOrDefault(nums[j], 0) + 1);
                res = Math.max(j - i + 1, res);
                
            } else {
                // 超过了
                while(nums[i]!=nums[j]){
                    mapCount.put(nums[i],mapCount.get(nums[i])-1);
                    i++;
                }
                i++;
            }
        }
        return res;
    }
}
```

4. 至少有k个重复字符的最长子串 [至少有k个重复字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/description/)
    1. [题解](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solutions/624045/xiang-jie-mei-ju-shuang-zhi-zhen-jie-fa-50ri1/)
    2. 题目说明了只包含小写字母（26 个，为有限数据），我们可以枚举最大长度所包含的字符类型数量，答案必然是 [1, 26]，即最少包含 1 个字母，最多包含 26 个字母。  
       你会发现，当确定了长度所包含的字符种类数量时，区间重新具有了**二段性质**
       当我们使用双指针的时候：
        * 右端点往右移动必然会导致字符类型数量增加（或不变）
        * 左端点往右移动必然会导致字符类型数量减少（或不变）

```java
class Solution {
    public int longestSubstring(String s, int k) {
        int ans = 0;
        int n = s.length();
        char[] cs = s.toCharArray();
        int[] cnt = new int[26];
        // 考虑1，2，3，4...... 个字符
        for(int p = 1;p<=26;p++){
            Arrays.fill(cnt, 0);
            // tot 代表 [j, i] 区间所有的字符种类数量；sum 代表满足「出现次数不少于 k」的字符种类数量
            for (int i = 0, j = 0, tot = 0, sum = 0; i < n; i++) {
                int u = cs[i] - 'a';
                cnt[u]++;
                // 如果添加到 cnt 之后为 1，说明字符总数 +1
                if (cnt[u] == 1) tot++;
                // 如果添加到 cnt 之后等于 k，说明该字符从不达标变为达标，达标数量 + 1
                if (cnt[u] == k) sum++;
                // 当区间所包含的字符种类数量 tot 超过了当前限定的数量 p，那么我们要删除掉一些字母，即「左指针」右移
                // 两段性质
                // 右端点往右移动必然会导致字符类型数量增加（或不变）
                // 左端点往右移动必然会导致字符类型数量减少（或不变）
                // 移动窗口 更新左端点 减少一个字符
                while (tot > p) { 
                    int t = cs[j++] - 'a';
                    cnt[t]--;
                    // 如果添加到 cnt 之后为 0，说明字符总数-1
                    if (cnt[t] == 0) tot--;
                    // 如果添加到 cnt 之后等于 k - 1，说明该字符从达标变为不达标，达标数量 - 1
                    if (cnt[t] == k - 1) sum--;
                }
                // 区间内字符种类数量 与复合条件的字符种类数量相等 那么就可以更新答案
                if (tot == sum) ans = Math.max(ans, i - j + 1);
            }
        }
        return ans;
    }
}
```
5. 2进制数组全部等于1 的最小操作 ==> 限定来窗口大小为3  [滑窗+贪心](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/?envType=daily-question&envId=2024-10-18)

```java


class Solution {
    public int minOperations(int[] nums) {
        // 感觉是滑动窗口+贪心 尝试一下,窗口大小限定为 3

        int right = 0;
        int res = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 0) {
                // 为0 才会进行反转
                // 如果窗口终点大于数组的尾下标 那么就是不可行的
                if (i + 2 > nums.length - 1) {
                    return -1;
                }
                res ++;
                right = i;
                while (right <= i + 2) {
                    nums[right] = 1 -nums[right] == 0?1 :0;
                    right++;
                }
            }
        }
        return res;
    }
}

class Solution {
   public int minOperations(int[] nums) {
      int cnt = 0;
      for(int i = 0; i < nums.length - 2; i++){
         if(nums[i] == 0){
            nums[i] = 1;
            nums[i + 1] = 1 - nums[i + 1];
            nums[i + 2] = 1 - nums[i + 2];
            cnt++;
         }
      }

      if(nums[nums.length - 2] == 0 || nums[nums.length - 1] == 0){
         return -1;
      }
      return cnt;
   }
}
```




---

