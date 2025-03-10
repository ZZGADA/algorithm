# 记忆化搜索

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


5. 最长公共子序列 [记忆化搜索递归 ](https://leetcode.cn/problems/longest-common-subsequence/description/)
   这题可以翻译成dp

```java
class Solution {
    private char[]s,t;
    private int[][] memo;


    public int longestCommonSubsequence(String text1, String text2) {
        s=text1.toCharArray();
        t=text2.toCharArray();
        int n = s.length;
        int m = t.length;

        // nm结尾的最长公共子序列
        memo = new int[n][m];

        for(int[] row : memo){
            Arrays.fill(row,-1);    // -1 表示没有计算过
        }

        return dfs(n-1,m-1);
    }

    private int dfs(int i ,int j){
        if(i<0 || j<0){
            return 0;
        }

        if (memo[i][j] != -1) {
            return memo[i][j]; // 之前计算过
        }
        
        // 找到重复元素的结尾
        if (s[i] == t[j]) {
            return memo[i][j] = dfs(i - 1, j - 1) + 1;
        }

        // s 和 t 各自移动
        // 找到自己最长的公共子序列长度
        return memo[i][j] = Math.max(dfs(i-1,j),dfs(i,j-1));
    }
}
```

6. 编辑距离 [编辑距离](https://leetcode.cn/problems/edit-distance/)
```java
class Solution {
    int[][] memory;
    public int minDistance(String word1, String word2) {
        int m = word1.length();
        int n = word2.length();
        memory = new int[m + 1][n + 1];
        return dfs(m,n,word1,word2);
    }

    // 从后往前 考虑前m 和 n个元素
    // dfs的定义：让word1前m个元素 和 word2 前n个元素相同要做的最小操作数
    public int dfs(int m,int n,String word1,String word2){
        // 结束状态 剩余没有处理的元素 
        if(m == 0){
            return n;
        }
        if(n == 0){
            return m;
        }
        
        // 考虑前word1前m 个元素 和 word2 前n个元素 使其相等需要的最小操作数
        if(memory[m][n] != 0){
            return memory[m][n];
        }


        if(word1.charAt(m - 1) == word2.charAt(n - 1)){
            // 向前找 进入下一个递归状态 
            memory[m][n] =  dfs(m - 1,n - 1,word1,word2);

        }else {
            // 插入 修改 删除 然后操作数+1
            memory[m][n] = Math.min(
                Math.min(
                    // 删除 考虑前m-1元素
                    dfs(m - 1,n,word1,word2),
                    // 插入 插入元素m不变 依然考虑考虑前m个元素 但是新增元素和word2的结尾元素相等 所以考虑前n-1个元素
                    dfs(m,n - 1,word1,word2)),
                
                // 修改
                dfs(m - 1,n - 1,word1,word2)) + 1;
        }
        return memory[m][n];
    }
}

```

7. [最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/solutions/2203001/shi-pin-jiao-ni-yi-bu-bu-si-kao-dong-tai-kgkg/)
```java
class Solution {
    public int longestPalindromeSubseq(String S) {
        // 题目要求是子序列：寻找最长回文子序列 然后返回长度
        char[] s = S.toCharArray();
        int n = s.length;
        int[][] memo = new int[n][n];   // i~j 的最长回文子序列长度
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示还没有计算过
        }
        return dfs(0, n - 1, s, memo);
    }

    private int dfs(int i, int j, char[] s, int[][] memo) {
        if (i > j) {
            return 0; // 空串
        }
        if (i == j) {
            return 1; // 只有一个字母
        }
        if (memo[i][j] != -1) { // 之前计算过
            return memo[i][j];
        }
        if (s[i] == s[j]) {
            // 外侧边界找到了 只看里面的子序列
            return memo[i][j] = dfs(i + 1, j - 1, s, memo) + 2; // 都选
        }
        return memo[i][j] = Math.max(dfs(i + 1, j, s, memo), dfs(i, j - 1, s, memo)); // 枚举哪个不选
    }
}
```
--- 
9. [交错字符串](https://leetcode.cn/problems/interleaving-string/description/?envType=problem-list-v2&envId=dynamic-programming)
```java
class Solution {
    public boolean isInterleave(String s1, String s2, String s3) {
        int n = s1.length();
        int m = s2.length();
        if (n + m != s3.length()) {
            return false;
        }

        int[][] memo = new int[n + 1][m + 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(s1.toCharArray(),s2.toCharArray(),s3.toCharArray(),n-1,m-1,memo);
    }

    // 表示 s3[:i+j+2] 能否由 s1[:i+1] 和 s2[:j+1] 交错组成
    public boolean dfs(char[] s1, char[] s2, char[] s3, int i, int j, int[][] memo) {
        if (i < 0 && j < 0) {
            return true;
        }

        if(memo[i+1][j+1] != -1){
            return memo[i+1][j+1] == 1 ;
        }

        boolean res = (i>=0 && s1[i] == s3[i+j+1] && dfs(s1,s2,s3,i-1,j,memo)) ||
                      (j>=0 && s2[j] == s3[i+j+1] && dfs(s1,s2,s3,i,j-1,memo));
        memo[i+1][j+1] = res ? 1:0;
        return res;
    }
}
```