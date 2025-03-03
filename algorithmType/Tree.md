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