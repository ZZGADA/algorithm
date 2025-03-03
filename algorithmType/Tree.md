# 树题目

--- 
1.[ 二叉搜索树第K小元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)
```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class Solution {
    int res;
    public int kthSmallest(TreeNode root, int k) {
        // 二叉搜索树 第k小的树
        findMinK(root,k,0);
        return res;
    }

    public int findMinK(TreeNode root, int k, int cnt) {
        if(root == null){
            return cnt;
        }

        cnt = findMinK(root.left,k,cnt)+1;
        if (cnt == k){
            res = root.val;
        }
        return findMinK(root.right,k,cnt);
    }

}
```

--- 
2. [二叉树的坡度](https://leetcode.cn/problems/binary-tree-tilt/?envType=problem-list-v2&envId=depth-first-search)
```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class Solution {
    int res;

    public int findTilt(TreeNode root) {
        res = 0;

        sumTilt(root);
        return res;
    }

    // 后续遍历
    public int sumTilt(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int sum = 0;
        int leftSum = sumTilt(root.left);
        int rightSum = sumTilt(root.right);

        sum = sum + leftSum + rightSum + root.val;
        res += countTilt(leftSum, rightSum);
        return sum;
    }

    // 计算坡度
    public int countTilt(int left, int right) {
        return Math.abs(left - right);
    }
}
```

--- 
3. [寻找重复子树](https://leetcode.cn/problems/find-duplicate-subtrees/?envType=problem-list-v2&envId=depth-first-search)
```java
class Solution {
    Map<String, Integer> map = new HashMap<>();
    List<TreeNode> ans = new ArrayList<>();
    
    public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
        dfs(root);
        return ans;
    }

    // 将当前节点的结构记录下来 
    // 然后通过map判断当前结构是否存在 
    
    String dfs(TreeNode root) {
        if (root == null) return " ";
        StringBuilder sb = new StringBuilder();
        sb.append(root.val).append("_");
        sb.append(dfs(root.left)).append(dfs(root.right));
        
        String key = sb.toString();
        map.put(key, map.getOrDefault(key, 0) + 1);

        if (map.get(key) == 2) ans.add(root);

        return key;
    }
}
```