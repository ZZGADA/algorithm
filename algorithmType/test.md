```java

package org.example.interview_zoom;

import java.util.*;

class TreeNode {
    int id;
    char letter;
    List<TreeNode> children;

    public TreeNode(int id, char letter) {
        this.id = id;
        this.letter = letter;
        this.children = new ArrayList<>();
    }
}

public class TreeQuery {
    private static TreeNode[] nodes;
    private static final String TARGET = "BUG";

    // 构建树
    private static void buildTree(int n, int[] parents, char[] letters) {
        nodes = new TreeNode[n + 1];
        for (int i = 1; i <= n; i++) {
            nodes[i] = new TreeNode(i, letters[i - 1]);
        }
        for (int i = 0; i < n - 1; i++) {
            int parentId = parents[i] + 1;
            int childId = i + 2;
            nodes[parentId].children.add(nodes[childId]);
        }
    }

    // 查找两个节点之间的路径
    private static List<TreeNode> findPath(int start, int end) {
        List<TreeNode> path = new ArrayList<>();
        boolean[] visited = new boolean[nodes.length];
        dfs(nodes[start], end, path, visited);
        return path;
    }

    // 深度优先搜索找路径
    private static boolean dfs(TreeNode node, int end, List<TreeNode> path, boolean[] visited) {
        visited[node.id] = true;
        path.add(node);
        if (node.id == end) {
            return true;
        }
        for (TreeNode child : node.children) {
            if (!visited[child.id]) {
                if (dfs(child, end, path, visited)) {
                    return true;
                }
            }
        }
        path.remove(path.size() - 1);
        return false;
    }

    // 检查路径上的字符序列是否包含 "BUG" 子序列
    private static boolean containsSubsequence(List<TreeNode> path) {
        int targetIndex = 0;
        for (TreeNode node : path) {
            if (node.letter == TARGET.charAt(targetIndex)) {
                targetIndex++;
                if (targetIndex == TARGET.length()) {
                    return true;
                }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int q = scanner.nextInt();
        scanner.nextLine();

        int[] parents = new int[n - 1];
        String parentLine = scanner.nextLine();
        for (int i = 0; i < n - 1; i++) {
            parents[i] = Integer.parseInt(String.valueOf(parentLine.charAt(i)));
        }

        char[] letters = scanner.nextLine().toCharArray();
        buildTree(n, parents, letters);

        for (int i = 0; i < q; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            List<TreeNode> path = findPath(u, v);
            if (containsSubsequence(path)) {
                System.out.println("NO");
            } else {
                System.out.println("YES");
            }
        }
        scanner.close();
    }
}
```

```java
package org.example.XM;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class MountaineeringPlan {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        // 读取适合登山的天数 n 和最多能在大本营生活的总天数 k
        int n = scanner.nextInt();
        int k = scanner.nextInt();
        int[] suitableDays = new int[n];
        // 读取每个适合登山的日期
        for (int i = 0; i < n; i++) {
            suitableDays[i] = scanner.nextInt();
        }
        // 调用计算最少移动次数的方法
        int result = calculateMinMoves(suitableDays, k);
        System.out.println(result-2);
    }

    public static int calculateMinMoves(int[] suitableDays, int k) {
        // 合并适合登山的日期为连续区间
        List<int[]> intervals = mergeIntervals(suitableDays);
        int totalDays = 0;
        int moves = 0;
        for (int[] interval : intervals) {
            int intervalDays = interval[1] - interval[0] + 1;
            if (totalDays + intervalDays <= k) {
                // 如果加入当前区间后总天数不超过 k，直接加入
                totalDays += intervalDays;
                moves += 2;
            } else {
                // 若超过 k，尝试从当前区间拆分出一部分
                int remainingDays = k - totalDays;
                if (remainingDays > 0) {
                    moves += 2;
                    totalDays = k;
                }
                break;
            }
        }
        return moves;
    }

    public static List<int[]> mergeIntervals(int[] suitableDays) {
        List<int[]> intervals = new ArrayList<>();
        int start = suitableDays[0];
        int end = suitableDays[0];
        for (int i = 1; i < suitableDays.length; i++) {
            if (suitableDays[i] == end + 1) {
                // 如果当前日期和前一个日期连续，更新区间结束日期
                end = suitableDays[i];
            } else {
                // 不连续则添加当前区间并开始新的区间
                intervals.add(new int[]{start, end});
                start = suitableDays[i];
                end = suitableDays[i];
            }
        }
        // 添加最后一个区间
        intervals.add(new int[]{start, end});
        return intervals;
    }
}
/**
 *
5 8
2 3 5 6 10
 */



```