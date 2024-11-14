# 单调栈

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

