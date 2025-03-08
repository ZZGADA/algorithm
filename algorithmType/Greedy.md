# 贪心


1. 二进制操作数    [二进制变为1的最小操作](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/submissions/574992852/?envType=daily-question&envId=2024-10-19)  
   这个解析要好好看看 [异或](https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/solutions/2956434/javapython3cmei-ju-wei-yun-suan-wu-xu-pa-0n4f/?envType=daily-question&envId=2024-10-19)
*    a^1 = ~a
*    a^0 = a
*    1 ^ 1 = 0
*    0 ^ 1 = 1
*    0 ^ 0 = 0
*    1 ^ 0 = 1

```java
class Solution {
    public int minOperations(int[] nums) {
        /**
            如果 x=0 且 k 是奇数，或者 x=1 且 k 是偶数，那么这 k 次操作执行完后 nums[i] 变成 1。所以如果 x ！= =k mod 2，则不需要操作。
            如果 x=0 且 k 是偶数（原先就是0） 或者 x =1 k是奇数（反转过 变为0了），所以需要进行操作 x == k mod 2 
        */
        int k = 0;
        for (int x : nums) {
            if (x == k % 2) { // 必须操作
                k++;
            }
        }
        return k;
    }
}
```


2. 会议室安排  [会议室安排最多] (https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/solutions/98224/chun-cui-de-tan-xin-mei-yong-you-xian-dui-lie-dai-/)

    1. 下面这个写法错误🙅，不对
```java
class Solution {
    public int maxEvents(int[][] events) {
        Arrays.sort(events, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[1] < o2[1]) {
                    return -1;
                } else if (o1[1] > o2[1]) {
                    return 1;
                } else {
                    // 相同就看开始时间，我们希望开始时间早的在前面
                    if (o1[0] < o2[0]) {
                        return -1;
                    } else if (o1[0] > o2[0]) {
                        return 1;
                    }
                }
                return 0;
            }
        });

        // 一个萝卜一个坑
        int flag = events[0][0];
        Set<Integer> set = new HashSet();
        int res = 0;

        for(int[] arr : events){
            int startTime = arr[0];
            int endTime = arr[1];
            

            if (flag>=startTime  && flag <=  endTime) {
                flag++;
                res++;
            } else if (startTime > flag) {
                res++;
                flag = startTime;
            }else{
                for (int i = startTime; i <=endTime && i<=events[0][0]; i++) {
                if (!set.contains(i)) {
                    set.add(i);
                    break;
                }
            }
            }
        }

        return res+set.size();
    }
}



```

2. 下面这个超时
```java

class Solution {
    public int maxEvents(int[][] events) {
        Arrays.sort(events, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[1] < o2[1]) {
                    return -1;
                } else if (o1[1] > o2[1]) {
                    return 1;
                } else {
                    // 相同就看开始时间，我们希望开始时间早的在前面
                    if (o1[0] < o2[0]) {
                        return -1;
                    } else if (o1[0] > o2[0]) {
                        return 1;
                    }
                }
                return 0;
            }
        });

        // 一个萝卜一个坑
        Set<Integer> set = new HashSet<>();
        for (int[] event : events) {
            int s = event[0];
            int e = event[1];
            for (int i = s; i <=e; i++) {
                if (!set.contains(i)) {
                    set.add(i);
                    break;
                }
            }
        }
        return set.size();
    }
}
```

--- 
3. 用优先队列 这个是对的  

```java 
class Solution {
    public int maxEvents(int[][] events) {
        // 首先排序：开始时间小的在前。这样是方便我们顺序遍历，把开始时间一样的都放进堆
        Arrays.sort(events, (o1, o2) -> o1[0] - o2[0]);
        // 小顶堆
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        // 结果、开始时间、events下标、有多少组数据
        int res = 0, last = 1, i = 0, n = events.length;
        while (i < n || !pq.isEmpty()) {
            // 将start相同的会议都放进堆里
            while (i < n && events[i][0] == last) {
                pq.offer(events[i++][1]);
            }
            // pop掉当前天数之前的
            while (!pq.isEmpty() && pq.peek() < last) {
                pq.poll();
            }
            // 顶上的就是俺们要参加的
            if (!pq.isEmpty()) {
                pq.poll();
                res++;
            }
            last++;
        }
        return res;
    }
}
```


3. 加油站 环绕一圈[加油站](https://leetcode.cn/problems/gas-station/description/?envType=problem-list-v2&envId=greedy)
```java
class Solution {
    // 前缀和
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int sumGasMCost = 0;
        int length = gas.length;
        int[] eachMinus = new int[length];
        int[] totalMinus = new int[length+1];

        for(int i = 0;i<length;i++){
            eachMinus[i] = gas[i] - cost[i];
            sumGasMCost += eachMinus[i];
            totalMinus[i+1] = totalMinus[i] + eachMinus[i];

        }

        // 如果消耗量大于加油量 那么一定无法回到起点
        if (sumGasMCost<0){
            return -1;
        }

        // 因为是一个环 看当前结点后面的加油增量 是否大于等于 当前结点之前的加油增量
        // 但同时需要保证 环绕一圈内每次前往下一个站点都是可行的 
        for(int i = 0;i<length;i++){
            if (eachMinus[i]>=0 &&  totalMinus[length]-totalMinus[i]+totalMinus[i]>=0){
                // 找到一个可行解 需要二次循环判断
                // 通过 cntBack 记录可到达的最远一步，如果break了 表示i～j的区间无法满足增量大于等于0 剪枝 跳过这一段区域 
                boolean judgeBack = true;
                int cntBack = 0;
                for(int j = i+1 ;j<length;j++){
                    if(totalMinus[j+1] - totalMinus[i] <0){
                        judgeBack = false;
                        cntBack = j;
                        break;
                    }
                }
                int sum = totalMinus[length] - totalMinus[i];

                boolean judgeFront = true;
                for(int j = 0;j<=i;j++){
                    if(totalMinus[j+1] - totalMinus[0] + sum <0){
                        judgeFront =false;
                        break;
                    }
                }
                
                if(judgeBack && judgeFront){
                    return i;
                }
                
                i = cntBack;
            }
        }

        return -1;

    }
}
```

4. [最接近的三数和](https://leetcode.cn/problems/3sum-closest/description/)
```java
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int length = nums.length;
        int gap = Integer.MAX_VALUE;
        int res = 0;
        for (int i = 0; i < length - 2; i++) {
            int left = i + 1, right = length - 1;
            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                int originGap = target - sum;
                int ggap = Math.abs(originGap);
                if (ggap < gap) {
                    res = sum;
                    gap = ggap;
                }

                if (originGap < 0) {
                    right--; // 大了
                } else if (originGap > 0) {
                    left++;
                } else {
                    return target;
                }
            }
        }

        return res;
    }
}
```

5. [跳跃游戏](https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150)
```java
class Solution {
    public boolean canJump(int[] nums) {
        int length = nums.length;
        int end = length - 1;
        int maxJumpIndex = nums[0];
        int index = 0;
        while (index <= maxJumpIndex) {
            maxJumpIndex = Math.max(maxJumpIndex, index + nums[index]);
            index++;
            if (maxJumpIndex >= end) {
                break;
            }
        }

        return maxJumpIndex >= end;
    }
}
```

6. [跳跃游戏2](https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150)
* 贪心做法（有动态规划做法 可见动态规划）
```java
class Solution {
    public int jump(int[] nums) {
        int end = nums.length - 1;
        int ans = 0, index = 0, maxJumpIndex = nums[0];

        if(nums.length == 1){
            return 0 ;
        }


        while (index <= maxJumpIndex && index <= end) {
            int temp = maxJumpIndex;
            while (index <= temp && index <= end) {
                maxJumpIndex = Math.max(maxJumpIndex, nums[index] + index); // 更新最远跳跃距离
                index++;    
            }
            ans++;
        }

        return ans;
    }
}
```

7. [最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/description/)
```java
class Solution {
     public int longestValidParentheses(String s) {
        int length = s.length();
        int maxans = 0;
        Deque<Integer> queue = new LinkedList<>();
        queue.push(-1);

        // 栈底元素为当前已经遍历过的元素中 「最后一个没有被匹配的右括号的下标」
        // 其他元素维护左括号(栈只会存左括号的下标)
        for(int i = 0;i<length;i++){
            if(s.charAt(i) == '('){
                queue.push(i);
            }else{
                queue.pop();
                if(queue.isEmpty()){
                    // 没有与之匹配的左括号
                    queue.push(i);
                }else{
                    // 最大值是当前元素 与上一个不可行解的距离
                    maxans = Math.max(maxans,i-queue.peek());
                }
            }
        }

        return maxans;
    }
}
```
--- 
8. [划分字母区间](https://leetcode.cn/problems/partition-labels/description/?envType=problem-list-v2&envId=greedy)
本质就是合并区间
```java
class Solution {
    public List<Integer> partitionLabels(String S) {
        // 本质是合并区间
        char[] s = S.toCharArray();
        int n = s.length;
        int[] last = new int[26];
        List<Integer> res = new ArrayList<>();

        // 记录字符出现的最后一个位置
        for (int i = 0; i < n; i++) {
            last[s[i] - 'a'] = i;
        }

        // 记录最远合并区间的长度
        int i, j, maxEdge = last[s[0] - 'a'];
        for (i = 0, j = 0; j < n; j++) {
            int index = s[j] - 'a';
            if (j <= maxEdge) {
                // 当前字符位置小于最大合并区间
                maxEdge = Math.max(maxEdge, last[index]);
            } else {
                // 当前字符位置大于最大合并区间
                res.add(j - i);
                i = j;
                maxEdge = last[index];
            }
        }
        if (i != j) {
            res.add(j - i);
        }
        return res;
    }
}
```

---
9. 多多的字符变换
多多君最近在研究字符串之间的变换，可以对字符串进行若干次变换操作:

交换任意两个相邻的字符，代价为0。
将任意一个字符a修改成字符b，代价为 |a - b|（绝对值）。
现在有两个长度相同的字符串X和Y，多多君想知道，如果要将X和Y变成两个一样的字符串，需要的最少的代价之和是多少。    

**示例1** 

```text
输入例子：
4
abca
abcd
输出例子：
3
例子说明：
其中一种代价最小的变换方案：
都修改为abcd，那么将第一个字符串X最后一个字符a修改为d，代价为|a - d| = 3。
```   
 
**示例2**   
```text
输入例子：
        4
baaa
        aabb
输出例子：
        1
例子说明：
其中一种代价最小的变换方案：
首先将第一个字符串通过交换相邻的字符：baaa -> abaa -> aaba，代价为0。
然后将第二个字符串修改最后一个字符b：|b - a| = 1。
两个字符都修改为aaba，所以最小的总代价为1。
```


```java
import java.util.Scanner;
import java.util.Arrays;

// 注意类名必须为 Main, 不要有任何 package xxx 信息
public class Main {
    int minCost;
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = Integer.parseInt(in.nextLine().trim());
        char[] up = in.nextLine().trim().toCharArray();
        char[] down = in.nextLine().trim().toCharArray();
        Arrays.sort(up);
        Arrays.sort(down);

        Main main = new Main();
        main.minCost = Integer.MAX_VALUE;
        main.judge(up, down, 0, n, 0);
        System.out.println(main.minCost);
    }

    public void judge(char[] up, char[] down, int index, int n, int cost) {
        if (index >= n) {
            minCost = Math.min(minCost, cost);
            return ;
        }

        if (up[index] == down[index]) {
            judge(up, down, index + 1, n, cost);
            return ;
        }

        // 直接替换的情况
        int replaceCost = cost + Math.abs(down[index] - up[index]);
        judge(up, down, index + 1, n, replaceCost);

        // 前后交换的情况
        // if (index < n - 1) {
        //     // 交换 up[index] 和 up[index + 1]
        //     if (up[index + 1] == down[index]) {
        //         swap(up, index, index + 1);
        //         judge(up, down, index + 1, n, cost);
        //         swap(up, index, index + 1); // 恢复状态
        //     }
        // 
        //     if (down[index + 1] == up[index]) {
        //         // 交换 down[index] 和 down[index + 1]
        //         swap(down, index, index + 1);
        //         judge(up, down, index + 1, n, cost);
        //         swap(down, index, index + 1); // 恢复状态
        //     }
        // }

        return ;
    }
    // 交换字符数组中两个位置的字符
    private void swap(char[] arr, int i, int j) {
        char temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

}
```