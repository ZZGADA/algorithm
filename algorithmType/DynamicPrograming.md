# 动态规划


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

此外：dp[i] 表示兑换金额为i的时候 可行的方案数
递推公式为：dp[i] = dp[i] + dp[i-coin] ; 表示为 硬币coin的加入 将从兑换i-coin的金额跳转过来。所以直接将dp[i-coin]加入到dp[i] 的方案数中。同时dp[i] 为上一轮递推记录的结果。
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

--- 

11. 最长公共子序列 [最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/solutions/696763/zui-chang-gong-gong-zi-xu-lie-by-leetcod-y7u0/)
```java
class Solution {
    public int longestCommonSubsequence(String text1, String text2) {
        char[] s = text1.toCharArray();
        char[] t = text2.toCharArray();
        int n = s.length;
        int m = t.length;
        int[][] f = new int[n + 1][m + 1];

        // 遍历所有text1[i]，text2[j]结尾的可能
        // 考虑前i+1，j+1的元素 i+1 只是为了移位
        // dp [i][j] 表示考虑前i，j个元素所能维护的最长自序列
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                f[i + 1][j + 1] = s[i] == t[j] ? f[i][j] + 1 :
                                  Math.max(f[i][j + 1], f[i + 1][j]);
            }
        }
        return f[n][m];
    }
}

```

--- 

12. 两个字符串的删除操作 [两个字符串的删除操作](https://leetcode.cn/problems/delete-operation-for-two-strings/)
    一定注意状态转移 和 dp[i][j] 表示的具体含义
```java

class Solution {
    public int minDistance(String word1, String word2) {
        char[] s = word1.toCharArray();
        char[] t = word2.toCharArray();
        int lens = s.length;
        int lent = t.length;

        // dp[j][j] 考虑前i 和 j个元素 的最长子序列长度
        int[][] dp = new int[lens+1][lent+1];
        for (int i = 1; i <= lens; i++) {
            for (int j = 1; j <= lent; j++) {
                if(s[i-1] == t[j-1]){
                    // 如果两个元素相等 
                    // 从各自的前i-1 和 j-1 的考虑范围转移 
                    dp[i][j] = dp[i-1][j-1]+1;
                }else{
                    dp[i][j] = Math.max(dp[i-1][j],dp[i][j-1]);
                }
            }
        }
        return lens+lent - 2*dp[lens][lent];
    }
}
```

--- 

13. 最大子数组和 [最大子数组和](https://leetcode.cn/problems/maximum-subarray/)
```java
class Solution {
    public int maxSubArray(int[] nums) {
        // dp[i] 表示 以 nums[i] 结尾的连续子数组的最大和
        int len = nums.length;
        int[] dp = new int[len];
        dp[0] = nums[0];
        int res = dp[0];

        for(int i = 1;i<len;i++){
            if(dp[i-1]+nums[i] > 0 ){
                // 上一个状态是负数 用max 过滤掉负数的起点 直接从正数开始
                dp[i] = Math.max(nums[i] , dp[i-1] + nums[i]);
            }else{
                // 重新开始
                // 子数组 不是子序列 所以一定有负数
                dp[i] = nums[i];
            }
            res = Math.max(dp[i],res);
        }
        

        return res;
    }
}
```

--- 
14. [最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/?envType=problem-list-v2&envId=dynamic-programming)
```java
class Solution {
    public String longestPalindrome(String s) {
        // 动态规划
        // dp[i][j] 表示 s[i~j]中是否为回文字符串
        // 递推公式：dp[i][j] == s[i] == s[j] && dp[i+1][j-1];
        // 注意substring 长度为2的特殊情况

        boolean[][] dp = new boolean[s.length()][s.length()];
        String res = "";
        for (int i = 0; i < s.length(); i++) {
            dp[i][i] = true;
            res = s.substring(i,i+1);
        }

        for (int step = 1; step < s.length(); step++) {
            for (int i = 0; i < s.length() - step; i++) {
                int j = i + step;
                dp[i][j] = s.charAt(i) == s.charAt(j) && (dp[i + 1][j - 1]|| step == 1);
                if (dp[i][j] && j-i +1 >res.length()) {
                    res = s.substring(i,j+1);
                }
            }
        }
        return res;
    }
}
```

--- 
15. 打家劫舍1 [打家劫舍1](https://leetcode.cn/problems/house-robber/?envType=problem-list-v2&envId=dynamic-programming)   
```java
class Solution {
    public int rob(int[] nums) {
        // 可以空两间房子 只要不连续就好了
        // 区分两种状态 当前房屋偷还是不透

        int[] dpSteal = new int[nums.length];
        int[] dpNotSteal = new int[nums.length];
        int res = 0;

        dpSteal[0] = nums[0];
        dpNotSteal[0] = 0;
        res = res = Math.max(dpNotSteal[0],dpSteal[0]);

        for (int i = 1; i < nums.length; i++) {
            dpSteal[i] = dpNotSteal[i - 1] + nums[i];
            dpNotSteal[i] = Math.max(dpNotSteal[i - 1], dpSteal[i - 1]); // 连续两次不偷
            res = Math.max(dpNotSteal[i],dpSteal[i]);
        }

        return res;
    }
}
```


16. [打家劫舍2](https://leetcode.cn/problems/house-robber-ii/)   
区分两种情况 直接递推就好了
```java
class Solution {
    public int rob(int[] nums) {
        // 可以空两间房子 只要不连续就好了
        // 区分两种状态 当前房屋偷还是不透
        return Math.max(stealFirst(nums, true), stealFirst(nums, false));
    }

    public int stealFirst(int[] nums, boolean ifStealFirst) {
        int[] dpSteal = new int[nums.length];
        int[] dpNotSteal = new int[nums.length];
        int res = 0;

        dpSteal[0] = ifStealFirst ? nums[0] : 0;
        dpNotSteal[0] = 0;
        res = res = Math.max(dpNotSteal[0], dpSteal[0]);

        for (int i = 1; i < nums.length; i++) {
            if (ifStealFirst && i == nums.length - 1) {
                // 偷了第一个 那么最后一个不能偷
                dpSteal[i] = Math.max(dpNotSteal[i - 1], dpSteal[i - 1]);
            }else {
                // 没偷第一个 那么最后一个可以偷
                dpSteal[i] = dpNotSteal[i - 1] + nums[i];
            }
            dpNotSteal[i] = Math.max(dpNotSteal[i - 1], dpSteal[i - 1]); // 连续两次不偷
            int temp = Math.max(dpNotSteal[i], dpSteal[i]);
            res = Math.max(res,temp);
        }
        return res;
    }
}
```

17. [跳跃游戏2](https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150)
* 动态规划做法 （还有贪心做法 见贪心）
```java
class Solution {
    public int jump(int[] nums) {
        // 返回最小跳跃次数
        // 两个条件
        // 1. 可到达边界
        // 2. 跳跃次数最少

        // dp[i] 跳到i 需要的最小跳跃数
        // 递推：dp[i] = dp[i-k] + 1
        int end = nums.length - 1;
        int[] dp = new int[nums.length];
        Arrays.fill(dp, nums.length);
        dp[0] = 0;
        int index = 0, maxJumpIndex = nums[0];
        while (index <= maxJumpIndex && index <= end) {
            int i = index + 1;
            while (i <= index + nums[index] && i <= end) {
                dp[i] = Math.min(dp[index] + 1, dp[i]);
                maxJumpIndex = Math.max(i + nums[i], maxJumpIndex);
                i++;
            }
            index++;
        }

        return dp[end];
    }
}
```
18.  [最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/solutions/2203001/shi-pin-jiao-ni-yi-bu-bu-si-kao-dong-tai-kgkg/)   
```java
class Solution {
   // 区间dp
   public int longestPalindromeSubseq(String S) {
      char[] s = S.toCharArray();
      int[][] dp = new int[s.length][s.length];
      // 根据递推公式 确定left为倒叙 right为正序
      for (int left = s.length - 1; left >= 0; left--) {
          for (int right = left; right < s.length; right++) {
              if (left == right) {
                  dp[left][right] = 1;
                  continue;
              }
              if (s[left] == s[right]) {
                  dp[left][right] = dp[left + 1][right - 1] + 2;
              } else {
                  dp[left][right] = Math.max(dp[left + 1][right], dp[left][right - 1]);
              }
          }
      }
      return dp[0][s.length-1];
   }
}
```

```java
class Solution {
    public int longestCommonSubsequence(String text1, String text2) {
        char[] s1 = text1.toCharArray();
        char[] s2 = text2.toCharArray();
        int[][] dp = new int[s1.length+1][s2.length+1]; // 边界情况 
        int end1 = s1.length-1,end2 = s2.length-1;

        for(int i = end1;i>=0;i--){
            for(int j = end2 ;j>=0;j--){
                if(s1[i] == s2[j]){
                    dp[i][j] = dp[i+1][j+1] + 1;
                }else{
                    dp[i][j] = Math.max(dp[i+1][j],dp[i][j+1]);
                }
            }
        }
        
        return dp[0][0];
    }
}
```

---- 
19. [买卖股票的最佳时机IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/)
```java
class Solution {
    public int maxProfit(int k, int[] prices) {
        int n = prices.length;
        int[][][] dp = new int[n][k + 1][2]; // 0 卖 1 买
        
         // 初始化
        for (int j = 0; j <= k; j++) {
            if (j == 1) {
                dp[0][j][1] = -prices[0];
            } else {
                dp[0][j][1] = Integer.MIN_VALUE;
            }
        }

        for (int i = 1; i < n; i++) {
            for (int j = 1; j <= k ; j++) {
                dp[i][j][0] = Math.max(dp[i-1][j][0],dp[i-1][j][1] + prices[i]);
                dp[i][j][1] = Math.max(dp[i-1][j][1],dp[i-1][j-1][0] - prices[i]);
            }
        } 

        int res = 0;
        for(int i = 1;i<=k;i++){
            res = Math.max(res,dp[n-1][i][0]);
        }
        return res;
    }
}
```
--- 
20. [交错字符串](https://leetcode.cn/problems/interleaving-string/description/?envType=problem-list-v2&envId=dynamic-programming)
```java
class Solution {
    public boolean isInterleave(String S1, String S2, String S3) {
        int n = S1.length();
        int m = S2.length();
        int g = S3.length();
        if (n + m != g) {
            return false;
        }

        char[] s1 = S1.toCharArray();
        char[] s2 = S2.toCharArray();
        char[] s3 = S3.toCharArray();

        // f[i][j] 来表示 S1 的前 i 个字符和 S2 的前 j 个字符是否能够交错组成 S3 的前 i + j 个字符。
        boolean[][] f = new boolean[n + 1][m + 1];
        f[0][0] = true;

        // 假设 s1 为空
        for (int j = 0; j < m; j++) {
            f[0][j + 1] = s2[j] == s3[j] && f[0][j];
        }

        for (int i = 0; i < n; i++) {
            // 要假设 s2为空
            f[i + 1][0] = s1[i] == s3[i] && f[i][0];
            for (int j = 0; j < m; j++) {
                f[i + 1][j + 1] = 
                s1[i] == s3[i + j + 1] && f[i][j + 1] ||
                        s2[j] == s3[i + j + 1] && f[i + 1][j];
            }
        }

        return f[n][m];
    }
}
```
21. [1和0](https://leetcode.cn/problems/ones-and-zeroes/description/?envType=problem-list-v2&envId=dynamic-programming)
```java
class Node {
    int zeorNum;
    int oneNum;

    Node(int zeorNum, int oneNum) {
        this.zeorNum = zeorNum;
        this.oneNum = oneNum;
    }
}

class Solution {

    public int findMaxForm(String[] strs, int m, int n) {
        List<Node> list = new ArrayList<>();
        int N = strs.length;
        for (String s : strs) {
            int count = 0;
            for (int i = 0; i < s.length(); i++) {
                if (s.charAt(i) == '1') {
                    count++;
                }
            }
            Node node = new Node(s.length() - count, count);
            list.add(node);
        }
        // 一个子序列问题
        // dp[i][j] 表示 考虑使用 i个0 和 j个1 能组成的最大子集
        int[][] dp = new int[m + 1][n + 1];
        for (Node node : list) {
            int nz = node.zeorNum;
            int no = node.oneNum;
            for (int i = m; i >= nz; i--) {
                for (int j = n; j >= no; j--) {
                    dp[i][j] = Math.max(dp[i][j], dp[i - nz][j - no] + 1);
                }
            }
        }

        return dp[m][n];
    }
}
```