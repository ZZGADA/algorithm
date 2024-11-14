# DFS


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