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

--- 
4. [求和路径Ï](https://leetcode.cn/problems/paths-with-sum-lcci/description/)
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
    Map<Long, Integer> map;
    int res;

    public int pathSum(TreeNode root, int sum) {
        // 本质是一个前缀和问题
        // sum[i] = sum[j+1] - target
        map = new HashMap<>();
        res = 0;
        map.put(0L,1);
        dfs(root,0,sum);
        return res;
    }

    public void dfs(TreeNode root,long pathSum,int target) {
        if(root == null){
            return ;
        }

        pathSum+= root.val;
        long find = pathSum - target;
        res += map.getOrDefault(find,0);
        map.put(pathSum,map.getOrDefault(pathSum,0)+1);

        dfs(root.left,pathSum,target);
        dfs(root.right,pathSum,target);

        map.put(pathSum,map.get(pathSum)-1);
    }
}
```

--- 
5. 多多的求和计算
```text 
多多路上从左到右有N棵树（编号1～N），其中第i个颗树有和谐值Ai。
多多鸡认为，如果一段连续的树，它们的和谐值之和可以被M整除，那么这个区间整体看起来就是和谐的。
现在多多鸡想请你帮忙计算一下，满足和谐条件的区间的数量。
```
示例1
```text
输入例子：
5 2
1 2 3 4 5
输出例子：
6
例子说明：
长度为1: [2], [4]
长度为2: 无
长度为3: [1,2,3], [3,4,5]
长度为4: [1,2,3,4], [2,3,4,5]
长度为5: 无
共6个区间的和谐值之和可以被2整除。
```

```java
import java.util.*;

// 注意类名必须为 Main, 不要有任何 package xxx 信息
public class Main {
    public static void main(String[] args) {
        // 前缀和 问题
        Scanner in = new Scanner(System.in);
        String[] nm = in.nextLine().trim().split(" ");
        int n = Integer.parseInt(nm[0]), m = Integer.parseInt(nm[1]);
        String[] numsStrs = in.nextLine().trim().split(" ");
        int[] nums = new int[numsStrs.length];
        long sums = 0L;
        Map<Long, Integer> map = new HashMap<>();
        map.put(0L, 1);
        long res = 0;
        for (int i = 0; i < numsStrs.length; i++) {
            nums[i] = Integer.parseInt(numsStrs[i]);
            sums += nums[i];
            long mod = sums % m;
            res += map.getOrDefault(mod, 0);
            map.put(mod, map.getOrDefault(mod, 0) + 1);
        }
        System.out.println(res);
    }
}
```

--- 
6. [和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/?envType=problem-list-v2&envId=prefix-sum)
**关键是同余定理** 
```java
class Solution {
    public int subarraysDivByK(int[] nums, int k) {
        Map<Integer, Integer> record = new HashMap<Integer, Integer>();
        record.put(0, 1);
        int sum = 0, ans = 0;
        for (int elem : nums) {
            sum += elem;
            // 注意 Java 取模的特殊性，当被除数为负数时取模结果为负数，需要纠正
            int modulus = (sum % k + k) % k;
            int same = record.getOrDefault(modulus, 0);
            ans += same;
            record.put(modulus, same + 1);
        }
        return ans;
    }
}
```

--- 
7.  [简化路径](https://leetcode.cn/problems/simplify-path/)

使用栈来解决

```go
func simplifyPath(path string) string {
	// path 绝对符合unix风格
	left, right := 0, 0
	size := len(path)
	stack := make([]string, 0)
	for right < size {
		// 抵达分隔符号
		if path[right:right+1] == "/" {
			middle := path[left:right]
			switch middle {
			case "..":
				if len(stack) > 0 {
					// 抛出元素
					stack = stack[:len(stack)-1]
				}
			case "/", ".", "":
				// 直接跳过
			default:
				// 入栈 && 指针滑动
				stack = append(stack, path[left:right])
			}
			left = right + 1
		} 
		right++
	}
    // right  抵达最后一位进行单独处理
	middle := path[left:right]
	switch middle {
	case "..":
		if len(stack) > 0 {
			// 抛出元素
			stack = stack[:len(stack)-1]
		}
	case "/", ".", "":
		// 直接跳过
	default:
		// 入栈 && 指针滑动
		stack = append(stack, path[left:right])
	}

	res := strings.Join(stack, "/")

	return "/" + res
}


// 简化写法
func simplifyPath(path string) string {
    stk := []string{}
    for _, s := range strings.Split(path, "/") {
        if s == "" || s == "." {
            continue
        }
        if s != ".." {
            stk = append(stk, s)
        } else if len(stk) > 0 {
            stk = stk[:len(stk)-1]
        }
    }
    return "/" + strings.Join(stk, "/")
}
```


8. [连续数组](https://leetcode.cn/problems/A1NYOS/description/) 
```go
func findMaxLength(nums []int) int {
	// 问题转换 求最长的连续子数组，其元素和为 0
	// 前缀和问题
	m := make(map[int]int)
	sum := 0
	res := 0
	m[0] = -1

	for i, v := range nums {
		if v == 0 {
			v = -1
			nums[i] = v
		}

		// 计算当前前缀和
		sum += v
		if pFlag, ok := m[sum]; ok {
			// 如果存在
			res = max(res, i-pFlag)
		} else {
			// 将前缀和元素 放入map
			// 如果当前前缀和元素重复，那么就记录早出现的

			m[sum] = i
		}
	}

	return res
}


```
