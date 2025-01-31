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

--- 
5. [指定区间的反转链表](https://leetcode.cn/problems/reverse-linked-list-ii/description/)
```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode ppre = new ListNode();
        ppre.next = head;
        int flag = 1;
        ListNode reverseHead = ppre;
        while (flag < left ) {
            reverseHead = reverseHead.next;
            flag++;
        }

        reverse(reverseHead, flag, right);
        return ppre.next;
    }

    public void reverse(ListNode ls, int flag, int right) {
        ListNode lr = null;
        ListNode node = ls.next, pre = ls.next, next = null;

        while (flag < right) {
            next = node.next; // 下一个节点
            node.next = next.next; // 指向下下个节点
            next.next = pre; // 指向前置节点
            ls.next = next;
            pre = ls.next;
            flag++;
        }
    }
}
```
--- 
6. [回文子串](https://leetcode.cn/problems/palindromic-substrings/)
 返回回文子串的个数
```java
class Solution {
    // 返回回文子串的个数
    public int countSubstrings(String s) {
        char[] arr = s.toCharArray();
        int res = arr.length;   
        for (int i = 0; i < arr.length; i++) {
            // 区分两种情况
            res += centerExpansion(arr, i - 1, i + 1);
            res += centerExpansion(arr, i, i + 1);
        }

        return res;
    }

    public int centerExpansion(char[] arr, int left, int right) {
        int res = 0;
        // 中心扩散
        while (left >= 0 && right < arr.length && arr[left] == arr[right]) {
            res++;
            left--;
            right++;
        }
        return res;
    }
}
```

```java
class Solution {
    // 返回回文子串的个数
    public int countSubstrings(String s) {
        char[] arr = s.toCharArray();
        int res = 0, left = 0, right = 0;
        boolean[][] dp = new boolean[arr.length][arr.length];
        for (int i = 0; i < arr.length; i++) {
            // 纯种中心扩散
            left = i;
            right = i;
            while (right < arr.length && arr[left] == arr[right]) {
                if (!dp[left][right]) {
                    dp[left][right] = true;
                    res++;
                }
                right++;
            }
            right--;
            left--;
            while (left >= 0 && arr[left] == arr[right]) {
                if (!dp[left][right]) {
                    dp[left][right] = true;
                    res++;
                }
                left--;
            }
            right++;
            while (left >= 0 && right < arr.length && arr[left] == arr[right]) {
                if (!dp[left][right]) {
                    res++;
                    dp[left][right] = true;
                }
                left--;
                right++;
            }
        }

        return res;
    }
}
```
--- 
7. [最长公共前缀和](https://leetcode.cn/problems/find-the-length-of-the-longest-common-prefix/solutions/2644176/liang-chong-xie-fa-yong-zi-fu-chuan-bu-y-qwh8/)
* 这个方法超时了 但是我觉得我写的没有错
```java
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        int prefixRes = 0;
        for (int num1 : arr1) {
            int prefix = 0;
            if(prefixRes != 0 && num1 < prefixRes){
                continue;
            }
            for (int num2 : arr2) {
                int prefixTemp;
                if (prefixRes != 0) {
                    prefixTemp = getPrefixSecond(num1, num2, prefixRes);
                } else {
                    prefixTemp = prefix == 0 ? getPrefixFirst(num1, num2) : getPrefixSecond(num1, num2, prefix);
                }
                prefix = Math.max(prefix, prefixTemp);
            }
            prefixRes = Math.max(prefix, prefixRes);
        }
        return prefixRes == 0 ? 0 : String.valueOf(prefixRes).length();
    }

    // 一组中已经知道最长的公共前缀了
    public int getPrefixSecond(int num1, int num2, int prefix) {
        // num1 不变 然后看公共前缀
        int len = String.valueOf(num1).length() - String.valueOf(prefix).length();
        int mod1 = (int) Math.pow(10, len);
        int mod2 = mod1;

        int prefix1, prefix2;
        while (mod1 > 0 && mod2 > 0) {
            prefix1 = num1 / mod1;
            prefix2 = num2 / mod2;
            if (prefix1 > prefix2) {
                mod2 /= 10;
            } else if (prefix1 < prefix2) {
                mod1 /= 10;
            } else {
                prefix = prefix1;
                mod1 /= 10;
                mod2 /= 10;
            }
        }
        return prefix;
    }

    // 一组中第一次求最长公共前缀
    public int getPrefixFirst(int num1, int num2) {
        int prefix = 0;
        int mod1 = 1, mod2 = 1;
        int prefix1 = num1 / mod1;
        int prefix2 = num2 / mod2;

        while (prefix1 > 0 && prefix2 > 0) {
            if (prefix1 > prefix2) {
                mod1 *= 10;
            } else if (prefix1 < prefix2) {
                mod2 *= 10;
            } else {
                prefix = prefix1;
                break;
            }
            prefix1 = num1 / mod1;
            prefix2 = num2 / mod2;
        }
        return prefix;
    }
}
```
* 这个是正确解法
```java
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        Set<Integer> set = new HashSet<Integer>();
        for (int num1 : arr1) {
            int mod = 1;
            int prefix = num1 / mod;
            while (prefix > 0) {
                set.add(prefix);
                mod *= 10;
                prefix = num1 / mod;
            }
        }

        int ans = 0;
        for (int num2 : arr2) {
            int mod = 1;
            int prefix = num2 / mod;
            while (prefix > 0) {
                if (set.contains(prefix)) {
                    int len = Integer.toString(prefix).length();
                    ans = ans > len ? ans : len;
                    break;
                }
                mod *= 10;
                prefix = num2 / mod;
            }
        }
        Integer.toString()
        return ans;
    }
}
```