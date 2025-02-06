# 前缀和

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


--- 


2.  路径总和3 [路径总和](https://leetcode.cn/problems/path-sum-iii/description/)    
每一个路径都视为前缀和的数组 ，每次遍历节点的时候就根据前缀和判断路径和是否与目标值匹配
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
    // 一个变向的前缀和问题
    // arr 1 2 3
    // sum 0 1 3 6
    // arr[i~j] = sum[j+1] -sum[i]

    // now - pre = targetSum
    // pre = now - targetSum
    // pre 的值 放入map 中 k: sum , v: 次数

    public Map<Long, Integer> preSum;
    public int result;

    public int pathSum(TreeNode root, int targetSum) {
        preSum = new HashMap<Long, Integer>();
        preSum.put(0L, 1);
        result = 0;

        find(root, 0L, targetSum);
        return this.result;
    }

    public void find(TreeNode root, long sum, int targetSum) {
        if (root == null) {
            return;
        }

        sum += root.val;

        // 先计算当前路径和中有多少个满足条件的前缀和，这样可以确保当前节点的值不会影响到路径和的计算
        this.result += preSum.getOrDefault(sum - targetSum, 0);
        preSum.put(sum, preSum.getOrDefault(sum, 0) + 1);
    
        find(root.left, sum, targetSum);
        find(root.right, sum, targetSum);
        preSum.put(sum, preSum.get(sum) - 1);
    }
}
```

3. [最长连续数组](https://leetcode.cn/problems/contiguous-array/)
* 超时版本
```java
class Solution {
    public int findMaxLength(int[] nums) {
        // 子数组中0和1的数量相同
        // 要求子数组长度最长

        // 使用前缀和
        int[] preSum1 = new int[nums.length + 1];
        int[] preSum0 = new int[nums.length + 1];
        int res = 0;

        // 前缀和初始化
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                preSum1[i + 1] = preSum1[i] + 1;
                preSum0[i + 1] = preSum0[i];
            } else {
                preSum0[i + 1] = preSum0[i] + 1;
                preSum1[i + 1] = preSum1[i];
            }
        }

        
        for (int j = 2; j <= nums.length; j++) {
            for (int i = j - 1; i >= 0 && i>=k; i--) {
                if(preSum1[j] - preSum1[i] == preSum0[j] - preSum0[i]){
                    res = Math.max(res,j-i);
                }
            }
        }
        return res;
    }
}
```

* **正确版本** 前缀和 + hash map
```java
class Solution {
    public int findMaxLength(int[] nums) {
        int n = nums.length, ans = 0;
        int[] sum = new int[n + 1];
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] == 1 ? 1 : -1);
        }

        map.put(0, 0);
        for (int i = 1; i <= n; i++) {
            int s = sum[i];
            // 先判断后插入 保证顺序
            if (map.containsKey(s)) {
                ans = Math.max(ans, i - map.get(s));
            } else {
                map.put(s, i);
            }
        }
        return ans;
    }
}
```