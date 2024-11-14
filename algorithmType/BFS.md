# BFS 


1. 拓扑排序+BFS  
   将有向无环图转变为线性排序的一个过程。重点：结点和入度和出度的关系  
   [LeetCode 207 课程表](https://leetcode.cn/problems/course-schedule/?envType=problem-list-v2&envId=breadth-first-search)
```java
class Solution {
    // 将一个有向无环图转换为线性的排序称之为拓扑排序
    // 如果存在一条有向边 A --> B，则这条边给 A 增加了 1 个出度，给 B 增加了 1 个入度。

    // 每次只能选入度为 0 的课，因为它不依赖别的课，是当下你能上的课。
    // 假设选了 0，课 3 的先修课少了一门，入度由 2 变 1。

    // 让入度为 0 的课入列，它们是能直接选的课。
    // 然后逐个出列，出列代表着课被选，需要减小相关课的入度。
    // 如果相关课的入度新变为 0，安排它入列、再出列……直到没有入度为 0 的课可入列。

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        int[] ans = new int[numCourses];// 存每个结点的入度
        List<List<Integer>> res = new ArrayList<>();// 存结点之间依赖关系
        Queue<Integer> queue = new LinkedList<>();

        // 初始化二维List集合
        for (int i = 0; i < numCourses; i++)
            res.add(new ArrayList<>());

        // 遍历每一个结点
        for (int[] temp : prerequisites) {
            ans[temp[0]]++;// 给需要依赖的结点入度
            res.get(temp[1]).add(temp[0]);  // 追加出度和入度的关系
        }

        //先把所有入度为0的结点加入队列
        for (int i = 0; i < numCourses; i++)
            if (ans[i] == 0)
                queue.add(i);

        while (!queue.isEmpty()) {
            int pre = queue.poll();
            numCourses--;   // 记录所有结点
            for (int relateNode : res.get(pre)) {
                if (--ans[relateNode] == 0) {
                    // 入度-1
                    queue.add(relateNode);
                }
            }
        }
        return numCourses == 0;

    }
}
```
