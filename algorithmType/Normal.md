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


--- 
2.  [整数转罗马数字](https://leetcode.cn/problems/integer-to-roman/?envType=study-plan-v2&envId=top-interview-150)
```java
class Solution {
    public String intToRoman(int num) {
        int[] arrNum = new int[] { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 };
        String[] arrChar = new String[] { "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" };
        int index = 0;
        StringBuilder sb = new StringBuilder();
        while (num != 0 && index < arrNum.length) {
            int chu = num / arrNum[index];
            int mod = num % arrNum[index];
            while (chu-- > 0) {
                sb.append(arrChar[index]);
            }
            index++;
            num = mod;
        }
        return sb.toString();
    }
}
```
---- 
3. [罗马数字转整数](https://leetcode.cn/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150)
```java
class Solution {
    public int romanToInt(String s) {
        char[] arr = s.toCharArray();
        Map<Character, Integer> map = new HashMap<>();
        map.put('I', 1);
        map.put('V', 5);
        map.put('X', 10);
        map.put('L', 50);
        map.put('C', 100);
        map.put('D', 500);
        map.put('M', 1000);

        int res = map.get(arr[arr.length - 1]);
        for (int i = arr.length - 2; i >= 0; i--) {
            if (map.get(arr[i]) < map.get(arr[i + 1])) {
                res -= map.get(arr[i]);
            } else {
                res += map.get(arr[i]);
            }
        }

        return res;
    }
}
```

--- 
4. [反转数组](https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150)
```java
class Solution {
    public void rotate(int[] nums, int k) {
        int length = nums.length;
        k %= length;
        reverse(nums, 0, length - 1);
        reverse(nums, 0, k - 1);
        reverse(nums, k, length - 1);
    }

    public void reverse(int[] nums, int i, int j) {
        while (i < j) {
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            i++;
            j--;
        }
    }
}
```