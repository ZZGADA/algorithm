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
* 中心扩算法 go 版本
```go 
import "fmt"

func longestPalindrome(s string) string {
	// 中心扩展
	// 每一个字符都可能是奇数开头的 也可以是偶数开头的
	// 分类讨论

	res := ""
	if len(s) == 0 {
		return res
	}
	res = string(s[0])

	for i := range s {
		// 奇数回文
		left, right := i-1, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
		// 偶数回文
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	return res

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
        Integer.toString();
        return ans;
    }
}
```

--- 
8. [LRU](https://leetcode.cn/problems/lru-cache/submissions/601965683/?envType=study-plan-v2&envId=top-100-liked)
关键步骤
1. list 是一个环形链表
2. pushToFront 中 有两个步骤。
   1. 先将节点从链表中remove（remove）
   2. 再将节点插入到第一个index的位置（addToFront）
3. put的时候，先判断是否到capacity
   1. 如果到了 先移除尾元素 （remove）
   2. 在将当前需要put的节点插入到第一个index的位置
```java
package org.example.algorithm_14;

import java.util.HashMap;
import java.util.Map;

class Node {
    public int key;
    public int value;

    public Node prev, next;

    public Node() {
    }

    public Node(int key, int value) {
        this.key = key;
        this.value = value;
        this.next = null;
        this.prev = null;
    }
}

class LRUCache {
    public int capacity;
    public Node dummy; // 哨兵节点 也就是链表的头节点 环形链表
    public Map<Integer, Node> keyToNode;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.dummy = new Node();
        this.dummy.next = dummy;
        this.dummy.prev = dummy;
        this.keyToNode = new HashMap<>();
    }

    public int get(int key) {
        if (!keyToNode.containsKey(key)) {
            return -1;
        }

        // 将 key的节点 移动到头节点
        this.pushToFront(key);
        return keyToNode.get(key).value;
    }


    // 将节点移动到链表头
    public void pushToFront(int key) {
        remove(key);
        addToFront(key);
    }

    public void put(int key, int value) {
        // 判断是否存在
        // 1. 存在更新value
        // 2. 不存在就创建
        // 3. 将插入节点孤立
        // 4. 将新对象移动到列表头部

        if (!keyToNode.containsKey(key)) {
            // 不存在
            if (keyToNode.size() >= capacity){
                int keyRemove = this.dummy.prev.key;
                remove(keyRemove);
                keyToNode.remove(keyRemove);
            }
            Node node = new Node(key, value);
            this.keyToNode.put(key, node);
            addToFront(key);
        } else {
            // 存在 更新节点value
            keyToNode.get(key).value = value;
            pushToFront(key);
        }
    }

    public void addToFront(int key){
        Node node = this.keyToNode.get(key);
        node.next = dummy.next;
        dummy.next.prev = node;
        node.prev = dummy;
        dummy.next = node;
    }

    public void remove(int key) {
        Node cur = this.keyToNode.get(key);
        cur.prev.next = cur.next;
        cur.next.prev = cur.prev;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

```