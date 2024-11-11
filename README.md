## ZZGEDA的算法学习记录

### BFS

1. 拓扑排序+BFS  
   将有向无环图转变为线性排序的一个过程。重点：结点和入度和出度的关系  
   [LeetCode 207 课程表](https://leetcode.cn/problems/course-schedule/?envType=problem-list-v2&envId=breadth-first-search)
```java
class Solution {
    // 将一个有向无环图转换为线性的排序称之为拓扑排序
    // 如果存在一条有向边 A --> B，则这条边给 A 增加了 1 个出度，给 B 增加了 1 个入度。

    // 每次只能选入度为 0 的课，因为它不依赖别的课，是当下你能上的课。
    // 假设选了 0，课 3 的先修课少了一门，入度由 2 变 1。

    // 让入度为 0 的课入列，它们是能直接选的课。
    // 然后逐个出列，出列代表着课被选，需要减小相关课的入度。
    // 如果相关课的入度新变为 0，安排它入列、再出列……直到没有入度为 0 的课可入列。

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        int[] ans = new int[numCourses];// 存每个结点的入度
        List<List<Integer>> res = new ArrayList<>();// 存结点之间依赖关系
        Queue<Integer> queue = new LinkedList<>();

        // 初始化二维List集合
        for (int i = 0; i < numCourses; i++)
            res.add(new ArrayList<>());

        // 遍历每一个结点
        for (int[] temp : prerequisites) {
            ans[temp[0]]++;// 给需要依赖的结点入度
            res.get(temp[1]).add(temp[0]);  // 追加出度和入度的关系
        }

        //先把所有入度为0的结点加入队列
        for (int i = 0; i < numCourses; i++)
            if (ans[i] == 0)
                queue.add(i);

        while (!queue.isEmpty()) {
            int pre = queue.poll();
            numCourses--;   // 记录所有结点
            for (int relateNode : res.get(pre)) {
                if (--ans[relateNode] == 0) {
                    // 入度-1
                    queue.add(relateNode);
                }
            }
        }
        return numCourses == 0;

    }
}
```

--- 

### 滑动窗口
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


### dp

1. 零钱兑换  
   **完全背包问题**：每一次的状态更新是从所有可行结果中转换而来的，dp[i-coins[j]]。然后+1，表示从之前的状态 +1个硬币
```java

class Solution {
    /**
     * 贪心不行：举个例子，假设硬币面值为 [1, 3, 4]，目标金额是 6。贪心策略会选择面值 4 的硬币，然后剩下 2，再选择两个 1，总共需要 3
     * 枚硬币。但实际
     * 上，选择两个 3 面值的硬币只需要 2 枚硬币，这是最优解。
     */

    public int coinChange(int[] coins, int amount) {
        // 转换为背包问题
        // 装满11的背包就要从装满10的背包转移过来
        // 每一次转移的时候 都需要遍历全部的硬币数量 保证可以从不同的状态进行转移
        int length = coins.length;
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, amount+1);  // 最坏情况 coins全是1 
        dp[0] = 0;

        // 遍历所有状态
        for (int i = 1; i <= amount; i++) {
            // 遍历所有coin
            for (int j = 0;j<length;j++){
                // 找到合适的可以放进来的coin 
                if(i-coins[j]>=0){
                    dp[i] = Math.min(dp[i-coins[j]]+1,dp[i]);
                }
            }
        }

        return dp[amount]!= amount+1 ? dp[amount]:-1;
    }
}
```
--- 

2. 零钱兑换2 [其实就是跳跃游戏](https://leetcode.cn/problems/coin-change-ii/description/)
题解看这个： [灵茶山艾府](https://leetcode.cn/problems/coin-change-ii/solutions/2706227/shi-pin-wan-quan-bei-bao-cong-ji-yi-hua-o3ew0/)

其实这题的关键就在于要把coins的遍历放在外面，因为要维持有序性，不然就会出现1，2和2，1重复计数的问题。   
将coins放在外面，那么内部遍历amount，保证当下一个coin被遍历的时候，上一个状态一定是可行解，同时有序。
```java
class Solution {
    public int change(int amount, int[] coins) {
        int[] dp = new int[amount+1];
        dp[0] = 1;

        for(int coin : coins){
            for(int i =coin ;i<=amount;i++){
                // coin = 1 每一个容量的组合数都是1
                // coin = 2 容量为2、3 组合数+1 为2  
                // 当容量为4 的时候 step（coin） 为 2 直接从容量为2的状态跳过来
                dp[i] += dp[i-coin];
            }
        }
        return dp[amount];
    }
}
```

--- 

3. 摆动序列 [摆动序列 DP](https://leetcode.cn/problems/wiggle-subsequence/description/?envType=problem-list-v2&envId=greedy)

其实这道题，你的思路是对的 ，但是状态转移方程写错了。  你原先的写法是递推出所有的可能并记录，这不对 ，“维护当前结点的最长摆动序列”这个状态没有维持出来，
写着写着就变成递推找符合峰谷的结点了，这个不对。下次注意⚠️。同时注意差值为0的时候，状态要维持，这个你也没有考虑到。

```java

class Solution {

    // dp 从后向前 当前结点的最长摆动序列 正负两种都要考虑到
    public int wiggleMaxLength(int[] nums) {
        int length = nums.length;
        int res = 1;
        int[] positive = new int[length];
        int[] negitive = new int[length];
        positive[length - 1] = 1;
        negitive[length - 1] = 1;

        for (int i = length - 2; i >= 0; i--) {
            // 两种情况
            // 向后找第一个大于自己的元素 同时元素是一个负数开始的摆动序列
            // 向后找第一个小于自己的元素 同时元素是一个正数开始的摆动序列

            int j = i + 1;
            if (nums[j] > nums[i]) {
                positive[i] = Math.max(positive[j], negitive[j] + 1);
                negitive[i] = negitive[j];
            } else if (nums[j] < nums[i]) {
                negitive[i] = Math.max(positive[j] + 1, negitive[j]);
                positive[i] = positive[j];
            } else {
                // 为0 的时候 要迁移过来
                positive[i] = positive[j];
                negitive[i] = negitive[j];
            }

            res = Math.max(res, Math.max(positive[i], negitive[i]));

        }

        return res;
    }
}
```

--- 


4. 最长递增自序列 [](https://leetcode.cn/problems/longest-increasing-subsequence/)
```java

class Solution {
    // 动态规划
    public int lengthOfLIS(int[] nums) {
       // dp[i] 当前结点的最长自序列

       int length = nums.length;
       int[] dp = new int[length];
       int res = 1;
       Arrays.fill(dp,1);

       for(int i = 1;i<length;i++){
          for(int j = i-1 ;j>=0;j--){
             // 向前寻找小于当前结点的元素 并维护最长递增子序列 
             if(nums[j]<nums[i]){
                dp[i] = Math.max(dp[i],dp[j]+1);
             }
          }
          res = Math.max(res,dp[i]);
       }

       return res;
    }

    // 单调栈
    public int lengthOfLIS1(int[] nums) {
        // 注意子序列是严格递增的
        List<Integer> queue = new ArrayList<>();
        int maxLength = 0;
        int length = nums.length;
        for (int i = 0; i < length; i++) {
            if (queue.isEmpty() || queue.get(queue.size() - 1) < nums[i]) {
                queue.add(nums[i]);
            } else {
                // 否则需要替换元素. 用于维持单调递增的序列
                // 注意不是直接抛出！！
                // 替换元素不会影响元素的最大状态 但是会影响后续的状态 替换是找到queue中第一个大于nums[i]
                // 所以此时替换后 更新位置到队列头 是新的可维持最长单调自序列
                // 使用二分查找
                maxLength = Math.max(maxLength, queue.size());
                swap(nums[i], queue);
            }
        }
        return Math.max(maxLength, queue.size());
    }

    public void swap1(int num, List<Integer> queue) {
        for (int i = queue.size() - 1; i >= 0; i--) {
            if (queue.get(i) < num) {
                queue.set(i + 1, num);
                return;
            }
        }
        if (queue.get(0) > num) {
            queue.set(0, num);
        }
    }

    // 寻找queue中第一个大于等于num的元素
    public void swap(int num, List<Integer> queue) {
        int left = 0;
        int right = queue.size() - 1;
        // 最后left 和 right左右相邻 mid始终落于left
        while (left <= right) {
            int mid = (left - right) / 2 + right;
            if (num > queue.get(mid)) {
                left = mid + 1; // left 寻找大于等于num的元素
            } else if (num < queue.get(mid)) {
                right = mid - 1;
            } else {
                return;
            }
        }
        queue.set(left, num);
    }

}

```

--- 


5. 和为目标值的最长子序列长度 [01背包](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/description/)    
目标值的大小就是我们的背包容量，自序列不重复就是01背包
```java

class Solution {
    public int lengthOfLongestSubsequence(List<Integer> nums, int target) {
        // 子序列和为target
        // 子序列要最长

        // 翻译成动态规划：01的背包问题 不能重复
        // dp[i] 表示target为i时 考虑前k个数组  能维持的最长子序列 dp[j] = dp[j-num]+1;
        int[] dp = new int[target + 1];
        
        for(Integer num : nums){
            for(int i = target;i>=num;i--){
                if(i == num){
                    dp[i] = Math.max(dp[i],1);
                    break;
                }
                if(dp[i-num]!=0){
                    // 表示已经放入进去了
                    dp[i] = Math.max(dp[i],dp[i-num]+1);
                }
            }
        }
        
        return dp[target]==0?-1:dp[target];

    }
}
```

6. 分割等和子集 [分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)
分割子序列，让子序列等于目标值。因为是就子集等于目标值，那么就可以转换成01背包
```java
class Solution {
    // 求子集和为 sum/2
    // 然后就又变成01背包问题了
    public boolean canPartition(int[] nums) {
        int sum = 0;
        int target = 0;
        for (int num : nums) {
            sum += num;
        }

        if (sum % 2 != 0) {
            return false;
        }
        target = sum / 2;

        // 考虑前k个元素的和是否等于target
        boolean[] dp = new boolean[target + 1];
        dp[0] = true;
        for (int i = 0; i < nums.length; i++) {
            for (int j = target; j >= nums[i]; j--) {
                if(dp[j-nums[i]]){
                    // 表示前k个元素已经放入了
                    dp[j] = true;
                }
            }
            if(dp[target]){
                return true;
            }
        }
        return false;
    }
}
```


7. 目标和 [这题用dp做绝了](https://leetcode.cn/problems/target-sum/description/)
如何转换为dp的思路很好,要点数学逻辑 [看这里艾神](https://leetcode.cn/problems/target-sum/solutions/2119041/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-s1cx/)
```java
class Solution {
    public int findTargetSumWays(int[] nums, int target) {
        int sum = 0;
        for (int x : nums) {
            sum += x;
        }

        // 最大加和小于target 或者 不能被2 整除就返回0
        sum -= Math.abs(target);
        if (sum < 0 || sum % 2 == 1) {
            return 0;
        }

        // 正数和为 （s+target）/2 负数和为（s-target）/2 ==> 转换成01背包
        // 求子序列和为 (s+target) /2 或者 （s- target）/2 的个数
        // target>0 取后面的 target < 0 取前面的
        int cap = sum / 2;
        int length = nums.length;
        int[] dp = new int[cap + 1];
        dp[0] = 1;

        // 考虑前k个数 和为cap的数量
        for (int i : nums) {
            for (int j = cap; j >= i; j--) {
                if(dp[j - i] != 0){
                    // 表示已经放入
                    dp[j] += dp[j-i];
                }
            }
        }

        return dp[cap];

    }
}

```

8. 最长有效括号 [dp](https://leetcode.cn/problems/longest-valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked)
思路：
   1. 确认子问题()()和(()) 两种情况

```java
class Solution {
    public int longestValidParentheses(String s) {
        int maxans = 0;
        int[] dp = new int[s.length()];
        for (int i = 1; i < s.length(); i++) {
            // 不存在以（结尾的有效括号 所以直接pass
            if (s.charAt(i) == ')') {
                // 找上一个 ）结束位置()()
                // 和 (()) 结束的位置 
                // 两种可能 
                if (s.charAt(i - 1) == '(') {
                    dp[i] = (i >= 2 ? dp[i - 2] : 0) + 2;
                } else if (i - dp[i - 1] > 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
                    dp[i] = dp[i - 1] + ((i - dp[i - 1]) >= 2 ? dp[i - dp[i - 1] - 2] : 0) + 2;
                }
                // 不具有后效行 （）（这样就中断了
                maxans = Math.max(maxans, dp[i]);
            }
        }
        return maxans;
    }
}
```

9. 将一个数变成幂的和的方案数 01背包 [将一个数变成幂的和的方案数](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/description/)
思路：
   1. 每个数只能用一次
   2. j-num 每次j的遍历， 其转移的起点都不同，那么当前dp[j] 的方案数就是 dp[j-num]+ dp[j]。
   3. 对于dp[j] 从dp[j-num]转移来的方案数，其表示的意义是dp[j-num] 状态时所有可行的组合结果+当前的num值，其方案数是不变的，只不过是方案结果集中的组合数+1了
```java
class Solution {
   public int numberOfWays(int n, int x) {
      // 01背包 n是容量 x（次数）
      // n = 160是容量 1～160都可以选 1～160的数字每个方案只能用一次
      // 最终要求的是可行的方案数
      int mod = (int)Math.pow(10, 9) + 7;
      long[] dp = new long[n + 1]; // dp[i] 表示容量为i刚好装满的方案数
      dp[0] = 1;
      for (int i = 1; i <= n; i++) {
         int num = (int)Math.pow(i,x);    // 当前的数
         for (int j = n; j >= num; j--) {
            if(dp[j-num]!=0){
               dp[j] += dp[j-num];
            }
         }
      }

      return (int)(dp[n]%mod);
   }
}
```


10. 数位成本和为目标值的最大数字 [完全背包+字符串处理](https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/)    
    今天做了个hard 感觉还行，dp+字符串，不是最优写法但是也通过了   
```java
class Solution {
   public String largestNumber(int[] cost, int target) {
      // 数位是i+1 成本是cost[i]
      // dp[i] 记录符合规则的 成本为i的最大整数
      int length = cost.length;
      StringBuilder[] dp = new StringBuilder[target + 1];
      for (int i = 0; i <= target; i++) {
         dp[i] = new StringBuilder();
      }

      for (int i = 0; i < length; i++) {
         String s = String.valueOf(i + 1);
         if (s.contains("0")) {
            continue;
         }

         // 数字要最大 (还需要重排序 因为第一个插入的元素无法满足目标值为最大)
         for (int j = cost[i]; j <= target; j++) {
             // 完全背包 找到可行解
            if (dp[j - cost[i]].length() != 0 || j - cost[i] == 0) {

               // 加入贪心 位数是逐渐递加的 那么位数如果要加入就一定是放在开头 这样就能维持最大数
               dp[j - cost[i]].insert(0,i+1);
               // 找到子集的最优解 
               if (ifCurrentLargeThenOriginal(dp[j], dp[j - cost[i]])) {
                  // 如果数字更大
                  dp[j] = new StringBuilder(dp[j - cost[i]]);
               }
               dp[j - cost[i]].deleteCharAt(0);
            }

         }
      }

      return dp[target].length() == 0 ? "0": dp[target].toString();

   }
    
   // 比较大小 维持dp[i] 表示最大数
   public boolean ifCurrentLargeThenOriginal(StringBuilder original, StringBuilder current) {
      int lenOriginal = original.length();
      int lenCurrent = current.length();
      if (lenCurrent > lenOriginal) {
         return true;
      } else if (lenCurrent < lenOriginal) {
         return false;
      } else {
         // 两个长度相等
         for (int i = 0; i < lenCurrent; i++) {
            if (current.charAt(i) > original.charAt(i)) {
               return true;
            } else if (current.charAt(i) < original.charAt(i)) {
               return false;
            }
         }
      }
      return false;

   }
}
```

ok 现在来看标准解法。但是说实话，我写不出来。用贪心反推，我想不到。  
思路：
1. 用dp求满足target的数的最大长度
2. 从大到小遍历“位数”，并反推dp 从而找到具体的位数是多少。（贪心）
妙，太妙了

```java

class Solution {
    public String largestNumber(int[] cost, int t) {
        int[] f = new int[t + 1];
        Arrays.fill(f, Integer.MIN_VALUE);
        f[0] = 0;
        // 第一个dp 求最大整数的最大长度 （此时最大整数并没有求出来 只知道长度）
        for (int i = 1; i <= 9; i++) {
            int u = cost[i - 1];
            for (int j = u; j <= t; j++) {
                f[j] = Math.max(f[j], f[j - u] + 1);
            }
        }
        if (f[t] < 0) return "0";
        String ans = "";

        // 贪心求最大整数 
        // j表示剩余值 u为花费 
        for (int i = 9, j = t; i >= 1; i--) {
            int u = cost[i - 1];
            // 状态转移的反推 
            // f[j] 是 f[j-u]+1 推出来的
            // 同时从大到小遍历 保证数字一直是最大的 
            while (j >= u && f[j] == f[j - u] + 1) {
                ans += String.valueOf(i);
                j -= u;
            }
        }
        return ans;
    }
}
```


### 记忆化搜索 


1. 单词拆分 [记其实是忆化搜索](https://leetcode.cn/problems/word-break/description/?envType=problem-list-v2&envId=memoization)
没有转移方程 只有递推
```java

class Solution {
    // 这道题其实就是背包问题
    public boolean wordBreak(String s, List<String> wordDict) {
        int length = s.length();
        boolean[] dp = new boolean[length+1];
        Map<String, Integer> map = new HashMap<>();

        // 初始化
        dp[0] = true;
        for (String ss : wordDict) {
            if (!map.containsKey(ss)) {
                map.put(ss, 0);
            }
        }

        // 向右递推
        for (int right = 0; right < length; right++) {
            int rightHead = right+1;
            for (String sub : wordDict){
                int subLen = sub.length();
                int leftHead = rightHead-subLen;
                if(leftHead<0) continue;
                if(map.containsKey(s.substring(leftHead,rightHead)) && dp[leftHead]){
                    dp[rightHead] = true;   
                }
            }
        }

        return dp[length];
    }
}
```

2. 回文字符串的总数  
      提示：记录每一个窗口大小的是否是回文字符串
```java
class Solution {
    public int countSubstrings(String s) {
        int length = s.length();
        if (length==1)return 1;
        boolean[][] dp = new boolean[length][length];
        int res = length;

        // 第一行和第二行的初始化 
        Arrays.fill(dp[0], true);
        for(int i=0;i<length-1;i++){
            dp[1][i] = s.charAt(i) == s.charAt(i+1);
            res += dp[1][i] ? 1:0;
        }

        // dp[i][j] => 表示的i为起点 长度为i+1的大小的子串是否是回文的
        for (int i = 2; i < length; i++) {
            // 需要初始化窗口大小
            int left;
            int right;
            int mid;
            for (int j = 0; j < length - i; j++) {
                left = j;
                right = left + i;
                dp[i][j] = dp[i-2][left+1] && (s.charAt(left) == s.charAt(right));
                res += dp[i][j] ? 1:0;
            }
            
        }
        return res;
    }
}
```


--- 

3. 向表达式添加括号后的最小结果 (中心扩散+记忆化)[https://leetcode.cn/problems/minimize-result-by-adding-parentheses-to-expression/description/]   
注意：
   1. 注意终止条件string.substring()的用法，如果是substring(1,1)这种返回的是“” 需要注意判断   


```java

class Solution {
    /**
     * 拆成2步
     * 1. 中心扩散遍历所有可能 时间复杂度是O(num1.length()*num2.length())
     * 2. 计算结果
     * 3. 比较大小 ，更新 ，使得expression最小
     * 4. 因为使用中心扩散 有一边的记录会重复计算 使用记忆化进行记录
     * 
     * 问题：
     * 1. 如何以O(1)的时间复杂度计算表达式的结果
     * 1.1. substring 转 int ，一个四个组成部分转int ，最后变成num1 * (num2 + num3) * num4
     */
    public String minimizeResult(String expression) {
        int opeIndex = 0;
        int resAns = Integer.MAX_VALUE;
        String res = "";
        StringBuilder sb = new StringBuilder(expression);
        int[][] nums = new int[expression.length() + 1][expression.length() + 1];

        for (int i = 0; i < expression.length(); i++) {
            if (expression.charAt(i) == '+') {
                opeIndex = i;
                break;
            }
        }

        // 遍历右边所有可组成数字的结果集
        int left = opeIndex + 1;
        for (int right = left + 1; right <= expression.length(); right++) {
            int numsRightPlus = Integer.parseInt(expression.substring(left, right));
            nums[left][right] = numsRightPlus;
            if (right != expression.length()) {
                int numsRightMulti = Integer.parseInt(expression.substring(right, expression.length()));
                nums[right][expression.length()] = numsRightMulti;
            }
        }

        // 遍历左边的所有结果集合 然后最终结果
        for (int lleft = opeIndex - 1; lleft >= 0; lleft--) {
            int numsLeftPlus = Integer.parseInt(expression.substring(lleft, opeIndex));
            int numsLeftMulti = 1;

            if (lleft != 0) {
                numsLeftMulti = Integer.parseInt(expression.substring(0, lleft));
            }

            for (int right = left + 1; right <= expression.length(); right++) {
                int multiRR = right != expression.length() ? nums[right][expression.length()] : 1;
                int tempRes = numsLeftMulti * (numsLeftPlus + nums[left][right]) * multiRR;
                if (tempRes < resAns) {
                    resAns = tempRes;
                    sb.insert(lleft, '(');
                    sb.insert(right+1, ')');
                    res = sb.toString();

                    sb = new StringBuilder(expression);
                }
            }
        }

        return res;

    }
}
```


--- 


###  dfs
优先总结：
1. dfs 一般和 分治一起出现，其实这都不难，关键在于找到递归状态的那个切入点
2. dfs和分治，就是不断将父状态分割成多个子状态，然后根据子状态的返回情况进行处理，所以找好子状态的分割场景很重要。


--- 

1. 分割回文字符串(dfs+记忆化)，记忆化直接通过dp实现
```java
class Solution {
    boolean[][] allHuiWen;
    List<List<String>> res;

    public String[][] partition(String s) {
        this.allHuiWen = initDp(s);
        this.res = new ArrayList<>();
        dfs(new StringBuilder(s), new ArrayList<String>(), 0);
        String[][] resStr  = new String[res.size()][];
        for(int i =0;i<res.size();i++){
            resStr[i] = res.get(i).toArray(new String[0]);
        }
        return resStr;
    }

    // 递归向下 进行分割
    public void dfs(StringBuilder s, List<String> list, int startIndex) {
        int length = s.length();
        if(startIndex>=length){
            res.add(new ArrayList(list));
        }
        // 遍历当前结点开始的所有可能的窗口大小 
        for (int winLen = 1; winLen <= length; winLen++) {
            if (startIndex+winLen>length){
                break;
            }
            // 剪枝+回溯
            if(allHuiWen[winLen-1][startIndex]){
                // 如果是回文串 继续递归
                list.add(s.substring(startIndex,startIndex+winLen));
                dfs(s,list,startIndex+winLen);
                list.remove(list.size()-1);
            }
        }
    }

    // 记录所有回文字符
    public boolean[][] initDp(String s) {
        int length = s.length();
        boolean[][] dp = new boolean[length][length];

        // 第一行和第二行的初始化
        Arrays.fill(dp[0], true);
        for (int i = 0; i < length - 1; i++) {
            dp[1][i] = s.charAt(i) == s.charAt(i + 1);
        }

        // dp[i][j] => 表示的j为起点 长度为i+1的大小的子串是否是回文的
        for (int i = 2; i < length; i++) {
            // 需要初始化窗口大小
            int left;
            int right;
            int mid;
            for (int j = 0; j < length - i; j++) {
                left = j;
                right = left + i;
                dp[i][j] = dp[i - 2][left + 1] && (s.charAt(left) == s.charAt(right));

            }
        }
        return dp;
    }
}
```


2. 为运算表达式设计优先级 [一个运算的表达式，选择括号添加的位置](https://leetcode.cn/problems/different-ways-to-add-parentheses/description/?envType=problem-list-v2&envId=memoization)

思路：（分治+dfs）
   1. 遇到操作数就左右递归，找左右各自的解
   2. 左右各自的解是递归的，这是一个分治的过程
   3. 关键是状态的返回值是一个List，每次都对左右状态的所有解进行多对多的匹配
   4. 最后将所有各自的解再经过操作数运算之后通过一个list存起来 返回给上一个状态 

```java
class Solution {
    boolean[] set;

    public List<Integer> diffWaysToCompute(String expression) {
        this.set = new boolean[256];
        set[(int) '+'] = true;
        set[(int) '-'] = true;
        set[(int) '*'] = true;
        set[(int) '/'] = true;

        return diffWays(expression, 0, expression.length() - 1);
    }

    public List<Integer> diffWays(String expression, int left, int right) {
        // 根据分割点 左右获取全部子集 然后最后按照操作数和在一起
        List<Integer> ans = new ArrayList<Integer>();

        if (left == right) {
            return Arrays.asList(expression.charAt(left) - '0');
        }

        if(right - left == 1){
            return Arrays.asList(Integer.valueOf(expression.substring(left,right+1)));
        }

        for (int pos = left + 1; pos < right; pos++) {
            int posFlag = (int) expression.charAt(pos);
            if (this.set[posFlag]) {
                // 表示是操作符
                List<Integer> ansLeft = diffWays(expression, left, pos - 1);
                List<Integer> ansRight = diffWays(expression, pos + 1, right);

                for (Integer numsLeft : ansLeft) {
                    for (Integer numsRight : ansRight) {
                        ans.add(calculate(numsLeft, numsRight, expression.charAt(pos)));
                    }
                }

            }
        }

        return ans;
    }

    public int calculate(int left, int right, char pos) {

        return switch (pos) {
            case '*' -> left * right;
            case '+' -> left + right;
            case '-' -> left - right;
            default -> 0;
        };
    }
}
```
   
--- 

3. 不同的二叉搜索树 [和第二题一样的思路，也是分治的思想，然后dfs](https://leetcode.cn/problems/unique-binary-search-trees-ii/description/)
```java 

/**
 * Definition for a binary tree node.
    * public class TreeNode {
    * int val;
    * TreeNode left;
    * TreeNode right;
    * TreeNode() {}
    * TreeNode(int val) { this.val = val; }
    * TreeNode(int val, TreeNode left, TreeNode right) {
    * this.val = val;
    * this.left = left;
    * this.right = right;
    * }
    * }
 */
class Solution {
    // 爆搜 遍历所有可能
    // 左右所有可能 然后构造
    int n;

    public List<TreeNode> generateTrees(int n) {
        this.n = n;

        return dfs(1, n);
    }

    public List<TreeNode> dfs(int left, int right) {
        List<TreeNode> ans = new ArrayList<>();

        if(right<left){
            ans.add(null);
            return ans;
        }

        if (left == right) {
            ans.add(new TreeNode(left, null, null));
            return ans;
        }

        // 进入递推
        for (int cur = left; cur <= right; cur++) {
            // 以cur为当前的跟结点 
            List<TreeNode> ansLeft = dfs(left,cur-1);
            List<TreeNode> ansRight = dfs(cur+1,right);
            

            for(TreeNode nodeLeft : ansLeft){
                for(TreeNode nodeRight : ansRight){
                    TreeNode curNode =  new TreeNode(cur,nodeLeft,nodeRight);
                    ans.add(curNode);
                }
            }

        }

        return ans;
    }
}


```


4. 目标和 [分情况+最终判断](https://leetcode.cn/problems/target-sum/)
```java

class Solution {
    int[] nums;
    int target;
    int res;

    public int findTargetSumWays(int[] nums, int target) {
        this.res = 0;
        this.nums = nums;
        this.target = target;

        dfs(0,0);
        return this.res;
    }

    public void dfs(int index,int record){
        if(index == nums.length){
            if(record == this.target){
                this.res++;
            }
            return ;
        }

        // 每个元素都有两种可能
        dfs(index+1,record + nums[index]);
        dfs(index+1,record - nums[index]);
    }
}
```


### 递归 
1. 快速幂
   
```java
class Solution {
    // 将3^10转换为9^5
    public double myPow(double x, int n) {
        if (x == 0) return 0;
        if (x == 1 || n == 0) return 1;

        long N = n; // 将n转换为long类型 防止溢出
        if (N < 0) {
            x = 1 / x;
            N = -N;
        }

        double res = 1.0;
        while (N > 0) {
            if (N % 2 == 1) {
                res *= x;
            }
            x *= x;
            N /= 2;
        }

        return res;
    }
}
```


--- 


### 二分查找
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


--- 

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
3. 经典二分 [搜索插入位置](https://leetcode.cn/problems/search-insert-position/description/?envType=problem-list-v2&envId=binary-search)
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


### 排序
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


### 前缀和
1. 24 mod [构成整天的下标对数目](https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-i/?envType=daily-question&envId=2024-10-22)
```java
class Solution {
   public int countCompleteDayPairs(int[] hours) {
      // 前缀和问题 ( hours[i] + hours[j] ) mod 24 == （hours[i] mod 24 + hours[j] mod 24 ）
      // mod 24 == 0
      // 注意⚠️： HashMap可以转换成int[24] 因为mod的范围就这么大
      int length = hours.length;
      int res = 0;
      Map<Integer, Integer> map = new HashMap<Integer, Integer>();

      for (int i = 0; i < length; i++) {
         int mod = hours[i] % 24;
         int need = (24 - mod) % 24;
         res += map.getOrDefault(need, 0);
         map.put(mod, map.getOrDefault(mod, 0)+1);
      }
      return res;
   }
}
```

### 贪心
1. 二进制操作数    [二进制变为1的最小操作](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/submissions/574992852/?envType=daily-question&envId=2024-10-19)  
   这个解析要好好看看 [异或](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/solutions/2956434/javapython3cmei-ju-wei-yun-suan-wu-xu-pa-0n4f/?envType=daily-question&envId=2024-10-19)   
*    a^1 = ~a   
*    a^0 = a   
*    1 ^ 1 = 0  
*    0 ^ 1 = 1 
*    0 ^ 0 = 0   
*    1 ^ 0 = 1 
   
```java
class Solution {
    public int minOperations(int[] nums) {
        /**
            如果 x=0 且 k 是奇数，或者 x=1 且 k 是偶数，那么这 k 次操作执行完后 nums[i] 变成 1。所以如果 x ！= =k mod 2，则不需要操作。
            如果 x=0 且 k 是偶数（原先就是0） 或者 x =1 k是奇数（反转过 变为0了），所以需要进行操作 x == k mod 2 
        */
        int k = 0;
        for (int x : nums) {
            if (x == k % 2) { // 必须操作
                k++;
            }
        }
        return k;
    }
}
```


2. 会议室安排  [会议室安排最多] (https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/solutions/98224/chun-cui-de-tan-xin-mei-yong-you-xian-dui-lie-dai-/)

   1. 下面这个写法错误🙅，不对 
```java
class Solution {
    public int maxEvents(int[][] events) {
        Arrays.sort(events, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[1] < o2[1]) {
                    return -1;
                } else if (o1[1] > o2[1]) {
                    return 1;
                } else {
                    // 相同就看开始时间，我们希望开始时间早的在前面
                    if (o1[0] < o2[0]) {
                        return -1;
                    } else if (o1[0] > o2[0]) {
                        return 1;
                    }
                }
                return 0;
            }
        });

        // 一个萝卜一个坑
        int flag = events[0][0];
        Set<Integer> set = new HashSet();
        int res = 0;

        for(int[] arr : events){
            int startTime = arr[0];
            int endTime = arr[1];
            

            if (flag>=startTime  && flag <=  endTime) {
                flag++;
                res++;
            } else if (startTime > flag) {
                res++;
                flag = startTime;
            }else{
                for (int i = startTime; i <=endTime && i<=events[0][0]; i++) {
                if (!set.contains(i)) {
                    set.add(i);
                    break;
                }
            }
            }
        }

        return res+set.size();
    }
}



```

   2. 下面这个超时  
   ```java

class Solution {
    public int maxEvents(int[][] events) {
        Arrays.sort(events, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[1] < o2[1]) {
                    return -1;
                } else if (o1[1] > o2[1]) {
                    return 1;
                } else {
                    // 相同就看开始时间，我们希望开始时间早的在前面
                    if (o1[0] < o2[0]) {
                        return -1;
                    } else if (o1[0] > o2[0]) {
                        return 1;
                    }
                }
                return 0;
            }
        });

        // 一个萝卜一个坑
        Set<Integer> set = new HashSet<>();
        for (int[] event : events) {
            int s = event[0];
            int e = event[1];
            for (int i = s; i <=e; i++) {
                if (!set.contains(i)) {
                    set.add(i);
                    break;
                }
            }
        }
        return set.size();
    }
}
```

   3. 用优先队列 这个是对的
   ```java

class Solution {

    public int maxEvents(int[][] events) {
        // 首先排序：开始时间小的在前。这样是方便我们顺序遍历，把开始时间一样的都放进堆
        Arrays.sort(events, (o1, o2) -> o1[0] - o2[0]);
        // 小顶堆
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        // 结果、开始时间、events下标、有多少组数据
        int res = 0, last = 1, i = 0, n = events.length;
        while (i < n || !pq.isEmpty()) {
            // 将start相同的会议都放进堆里
            while (i < n && events[i][0] == last) {
                pq.offer(events[i++][1]);
            }
            // pop掉当前天数之前的
            while (!pq.isEmpty() && pq.peek() < last) {
                pq.poll();
            }
            // 顶上的就是俺们要参加的
            if (!pq.isEmpty()) {
                pq.poll();
                res++;
            }
            last++;
        }
        return res;

    }
}
```


3. 加油站 环绕一圈[加油站](https://leetcode.cn/problems/gas-station/description/?envType=problem-list-v2&envId=greedy)
```java
class Solution {
    // 前缀和
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int sumGasMCost = 0;
        int length = gas.length;
        int[] eachMinus = new int[length];
        int[] totalMinus = new int[length+1];

        for(int i = 0;i<length;i++){
            eachMinus[i] = gas[i] - cost[i];
            sumGasMCost += eachMinus[i];
            totalMinus[i+1] = totalMinus[i] + eachMinus[i];

        }

        // 如果消耗量大于加油量 那么一定无法回到起点
        if (sumGasMCost<0){
            return -1;
        }

        // 因为是一个环 看当前结点后面的加油增量 是否大于等于 当前结点之前的加油增量
        // 但同时需要保证 环绕一圈内每次前往下一个站点都是可行的 
        for(int i = 0;i<length;i++){
            if (eachMinus[i]>=0 &&  totalMinus[length]-totalMinus[i]+totalMinus[i]>=0){
                // 找到一个可行解 需要二次循环判断
                // 通过 cntBack 记录可到达的最远一步，如果break了 表示i～j的区间无法满足增量大于等于0 剪枝 跳过这一段区域 
                boolean judgeBack = true;
                int cntBack = 0;
                for(int j = i+1 ;j<length;j++){
                    if(totalMinus[j+1] - totalMinus[i] <0){
                        judgeBack = false;
                        cntBack = j;
                        break;
                    }
                }
                int sum = totalMinus[length] - totalMinus[i];

                boolean judgeFront = true;
                for(int j = 0;j<=i;j++){
                    if(totalMinus[j+1] - totalMinus[0] + sum <0){
                        judgeFront =false;
                        break;
                    }
                }
                
                if(judgeBack && judgeFront){
                    return i;
                }
                
                i = cntBack;
            }
        }

        return -1;

    }
}
```

### 单调栈
1. [去除重复字母，保证 返回结果的字典序最小同时维持字典序（要求不能打乱其他字符的相对位置）](https://leetcode.cn/problems/remove-duplicate-letters/description/?envType=problem-list-v2&envId=greedy)   
要求是最终子序列元素不重复，字典序最小，不能打乱原字符在字符串中的序列。咋一看，难。一步一步来。    
`cbacdcbc` 的答案是 `acdb` 不是 `abcd`。虽然 abcd 的字典序最小 但是不符合 （不能打乱原字符在字符串中的序列），因为c出现在d和b之前。    
最终字典序最小的前提是不能打乱原字符的顺序。只有 当前元素在后面还会重复出现，才有重新排序的可能。


```java

class Solution {
    /**
    1. 通过单调栈来保证字符的字典序 如果栈顶元素小于当前元素需要pop 否则add
        1.1. pop的时候注意：如果栈顶元素后面不会再出现了 那么就不能再pop 了 bcac  ==> bac  ，不然 acb是最小的
    2. 如果当前遍历的元素已经存在在栈空间中，那么就不需要进行任何操作
    3. 抛出元素就表示需要重新排列字典序了
    4. 栈顶元素是否在之后还会出现 需要优先遍历一遍才能知道
    5. 维护一个元素不重复的单调栈 （插入元素时，当前元素可能不会再出现）*/
    public String removeDuplicateLetters(String s) {
        // 记录字符数量
        int[] alpha = new int[26];
        boolean[] inStack = new boolean[26];
        StringBuilder sb = new StringBuilder();
        Deque<Character> queue = new LinkedList<>();

        // 统计 每个字符的数量有多少
        for (int i = 0 ; i < s.length(); i++) {
            alpha[s.charAt(i) - 'a']++;
        }

        for (char c : s.toCharArray()) {
            int index = c - 'a';
            alpha[index]--;

            // 当前字符在之前已经被确认了（符合字典序）
            if (inStack[index]) {
                continue;
            }

            while(!queue.isEmpty( )&& c < queue.peekLast()){
                if(alpha[queue.peekLast()-'a']==0){
                    // 表示当前元素后面不会重复了
                    break;
                }

                // 表示 后面还会出现与当前栈顶元素相同的字符
                inStack[queue.pollLast()-'a'] = false; 
            }

            queue.addLast(c);
            inStack[index] = true;
        }

        queue.forEach(e -> sb.append(e));
        return sb.toString();

    }
}
```



2. 最长有效括号 [最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked)  
思路：
   1. 栈底元素为当前已经遍历过的元素中 「最后一个没有被匹配的右括号的下标」
   2. 这个作用是划定子串的区间
   3. **其实也是变相的只考虑前k个元素**


```java

class Solution {
    public int longestValidParentheses(String s) {
        int length = s.length();
        int maxans = 0;
        Deque<Integer> queue = new LinkedList<>();
        queue.push(-1);

        // 栈底元素为当前已经遍历过的元素中 「最后一个没有被匹配的右括号的下标」
        // 其他元素维护左括号
        for(int i = 0;i<length;i++){
            if(s.charAt(i) == '('){
                queue.push(i);
            }else{
                queue.pop();

                if(queue.isEmpty()){
                    // 没有与之匹配的左括号
                    queue.push(i);
                }else{
                    maxans = Math.max(maxans,i-queue.peek());
                }
            }
        }

        return maxans;
    }
}
```




### 链表  [LRU](https://leetcode.cn/problems/lru-cache/)
1. LRU
```java
class LRUCache {
    private static class Node {
        int key, value;
        Node prev, next;

        Node(int k, int v) {
            key = k;
            value = v;
            
        }
    }

    private final int capacity;
    private final Node dummy = new Node(0, 0); // 哨兵节点 头指针
    private final Map<Integer, Node> keyToNode = new HashMap<>();

    // 关键问题：如何逐出最久未使用的关键字 
    // 以及注意：keyToNode只负责维护key和node的关系 但是 访问次数的顺序是node自身的双向链表决定的
    public LRUCache(int capacity) {
        this.capacity = capacity;
        dummy.prev = dummy;
        dummy.next = dummy;     // !!!! 这两步很重要 形成环形链表
    }
    
    // 获取数据 同时更新元素的头
    public int get(int key) {
        Node res =this.getNode(key);
        return res == null ? -1 : res.value;
    }
    

    public void put(int key, int value) {
        Node node = getNode(key);
        if(node == null){
            // 更新新的
            Node n = new Node(key,value);
            keyToNode.put(key,n);
            this.pushFront(n);  // 推到最顶

            if(keyToNode.size() > this.capacity){
                // 淘汰旧的
                Node backNode = dummy.prev;    // 尾巴
                keyToNode.remove(backNode.key);
                remove(backNode);
            }
        }else{
            node.value = value;
            return ;
        }
    }

    private Node getNode(int key){
        if(!keyToNode.containsKey(key)){
            return null;
        }

        Node node = keyToNode.get(key);  // 结点存在
        remove(node);                    // 将结点删除掉
        pushFront(node);                 // 结点插入头
        
        return  node ;
    }

    // 删除结点 双向指针移动
    private void remove(Node x) {
        x.prev.next = x.next;
        x.next.prev = x.prev;
    }

    // 头结点插入并更新
    private void pushFront(Node x){
        x.next = this.dummy.next;
        x.prev = this.dummy;

        x.next.prev = x;
        x.prev.next = x;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */


```


