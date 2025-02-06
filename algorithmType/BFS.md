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

2. [标准bfs搜索](https://leetcode.cn/problems/number-of-islands/?envType=problem-list-v2&envId=breadth-first-search)
```java
class Solution {
    class Node {
        public int i;
        public int j;

        public Node(int i, int j) {
            this.i = i;
            this.j = j;
        }
    }

    // 上下左右
    public int numIslands(char[][] grid) {
        int m = grid.length, n = grid[0].length;
        int res = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    // 为陆地 进行bfs搜索
                    bfs(i, j, grid, m, n);
                    res++;
                }
            }
        }

        return res;
    }

    public void bfs(int i, int j, char[][] grid, int m, int n) {
        Queue<Node> queue = new LinkedList<>();
        queue.add(new Node(i, j));
        while (!queue.isEmpty()) {
            int size = queue.size();
            while (size-- > 0) {
                Node node = queue.poll();
                // 边界情况
                if (node.i - 1 >= 0 && grid[node.i - 1][node.j] == '1') {
                    grid[node.i - 1][node.j] = '0';
                    queue.add(new Node(node.i - 1, node.j));
                }

                if (node.j + 1 < n && grid[node.i][node.j + 1] == '1') {
                    grid[node.i][node.j + 1] = '0';
                    queue.add(new Node(node.i, node.j + 1));
                }

                if (node.i + 1 < m && grid[node.i + 1][node.j] == '1') {
                    grid[node.i + 1][node.j] = '0';
                    queue.add(new Node(node.i + 1, node.j));
                }

                if (node.j - 1 >= 0 && grid[node.i][node.j - 1] == '1') {
                    grid[node.i][node.j - 1] = '0';
                    queue.add(new Node(node.i, node.j - 1));
                }
            }
        }
    }

}
```

3. [省份数量](https://leetcode.cn/problems/number-of-provinces/description/?envType=problem-list-v2&envId=breadth-first-search)
```java
class Solution {
    class Node {
        public int i;
        public int j;

        public Node(int i, int j) {
            this.i = i;
            this.j = j;
        }
    }

    public int findCircleNum(int[][] grid) {
        // bfs (注意有环 ， 和最大岛屿面积 没有什么差别)
        int m = grid.length, n = grid[0].length;
        int res = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    // 为陆地 进行bfs搜索
                    bfs(i, j, grid, m, n);
                    res++;
                }
            }
        }

        return res;
    }

    public void bfs(int i, int j, int[][] grid, int m, int n) {
        Queue<Node> queue = new LinkedList<>();
        queue.add(new Node(i, j));
        grid[i][j] = 0;

        while (!queue.isEmpty()) {
            int size = queue.size();
            while (size-- > 0) {
                Node nodeOrigin = queue.poll();
                Node node = new Node(nodeOrigin.j, nodeOrigin.i);
                // 边界情况
                for (int k = 0; k < n; k++) {
                    if(grid[node.i][k] == 1){
                        queue.add(new Node(node.i,k));
                        grid[node.i][k] = 0;
                    }
                }
            }
        }
    }
}
```