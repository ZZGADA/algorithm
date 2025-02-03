# 递归



1. 快速幂

```java
class Solution {
    // 将3^10转换为9^5
    public double myPow(double x, int n) {
        if (x == 0) return 0;
        if (x == 1 || n == 0) return 1;

        long N = n; // 将n转换为long类型 防止溢出
        if (N < 0) {
            x = 1 / x;
            N = -N;
        }

        double res = 1.0;
        while (N > 0) {
            if (N % 2 == 1) {
                res *= x;
            }
            x *= x;
            N /= 2;
        }

        return res;
    }
}
```

2. [括号生成](https://leetcode.cn/problems/generate-parentheses/description/)
```java
class Solution {
    List<String> res;
    int end;

    public List<String> generateParenthesis(int n) {
        this.res = new ArrayList<String>();
        this.end = n * 2;
        StringBuilder sb =new StringBuilder();
        Deque<Character> stack = new LinkedList<Character>();
        stack.push('(');
        sb.append('(');
        dfs(stack,sb, 1);

        return res;
    }

    // 回溯
    public void dfs(Deque<Character> stack, StringBuilder sb ,int depth) {
        // 边界条件
        if (depth >= end) {
            // 为空 表示符合
            if (stack.isEmpty()){
                this.res.add(sb.toString());
                return;
            }
            return;
        }

        depth++;
        stack.push('(');
        sb.append('(');

        dfs(stack,sb,depth);

        stack.pop();
        sb.deleteCharAt(sb.length()-1);

        // 递归右括号
        Deque<Character> stackCopy = new LinkedList(stack);
        if(!stackCopy.isEmpty() && stackCopy.peekLast().equals('(')){
            stackCopy.pollLast();
            sb.append(')');
            dfs(stackCopy,sb,depth);
            sb.deleteCharAt(sb.length()-1);
        }
    }
}
```