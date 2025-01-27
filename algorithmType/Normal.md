# 常规题目
--- 
1. 最长回文子串   [最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/?envType=problem-list-v2&envId=dynamic-programming)   
* 中心扩算法
```java
class Solution {
    public String longestPalindrome(String s) {
        char[] arr = s.toCharArray();
        int length = arr.length;
        String res = "";
        for (int i = 0; i < length; i++) {
            int left = i, right = i;
            while (right < length - 1 && arr[left] == arr[right + 1]) {
                right++;
            }
            while (left > 0 && right < length && arr[right] == arr[left - 1]) {
                left--;
            }

            while (left > 0 && right < length - 1 && arr[left - 1] == arr[right + 1]) {
                left--;
                right++;
            }
            if (res.length() < right - left + 1) {
                res = s.substring(left, right + 1);
            }
        }
        return res;
    }
}
```

* 动态规划
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