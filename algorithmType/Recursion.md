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

--- 


2. [组合数字](https://leetcode.cn/problems/combination-sum-ii/description/)
```go
func combinationSum3(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	size := len(candidates)
	var dfs func(flag int, cnt int, tmp []int)

	// 本质上是一个递归树
	// 然后进行选和不选
	// 选的时候要过滤掉和当前元素一样的情况，因为同一个开头的树可能会有同样的组合
	dfs = func(flag int, cnt int, tmp []int) {
		if cnt == target {
			t := append([]int(nil), tmp...)
			res = append(res, t)
			return
		}

		if cnt > target || flag >= size {
			return
		}

		// 选
		tmp = append(tmp, candidates[flag])
		cnt += candidates[flag]
		dfs(flag+1, cnt, tmp)

		// 回溯
		cnt -= candidates[flag]
		tmp = tmp[:len(tmp)-1]

		// 不选
		flag++
		for flag < size && candidates[flag] == candidates[flag-1] {
			flag++
		}
		dfs(flag, cnt, tmp)
	}

	dfs(0, 0, []int{})
	return res
}
```

3. [全排列](https://leetcode.cn/problems/permutations/?envType=study-plan-v2&envId=top-interview-150)
```go
func permute(nums []int) [][]int {
	size := len(nums)
	res := make([][]int, 0)

	isUsed := make([]int, size) // 0 没有用 1 用了

	var dfs func(arr []int,depth int)
	dfs = func(arr []int, depth int) {
		if depth >= size {
			res = append(res, append([]int(nil), arr...))
			return
		}

		for i, v := range isUsed {
			if v == 0 {
				// 没有使用
				arr = append(arr, nums[i])
				isUsed[i] = 1
				dfs(arr, depth+1)

                // 回溯
				isUsed[i] = 0
				arr = arr[:len(arr)-1]
			}
		}
	}
	dfs([]int{}, 0)

	return res
}


```



