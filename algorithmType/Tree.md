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
--- 
4. [根据二叉树创建字符串](https://leetcode.cn/problems/construct-string-from-binary-tree/)
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
    // 输入 输出关系要一一对应
    // 如果 root 左右子节点都不存在，则返回 root
    // 如果 root 左右子节点都存在，则返回 root(left)(right)
    // 如果 root 只有左节点存在，则返回 root(left)
    // 如果 root 只有右节点存在，则返回 root()(right)
    public String tree2str(TreeNode root) {
        return transalte(root, new StringBuilder()).toString();
    }

    public StringBuilder transalte(TreeNode root, StringBuilder sb) {
        if (root == null) {
            return sb;
        }

        sb.append(root.val);
        if (root.left == null && root.right == null) {
            // 如果是叶子节点 doing nothing
            return sb;
        }

        sb.append('(');
        transalte(root.left, sb).append(')');
        if (root.right != null) {
            sb.append('(');
            transalte(root.right, sb).append(')');
        }
        return sb;
    }
}
```
--- 

5. [从前序遍历和中序遍历构造2叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)   
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
    Map<Integer, Integer> indexMap;
    int[] preorder;
    int[] inorder;

    public TreeNode buildTree(int[] pre, int[] in) {
        indexMap = new HashMap<>();
        preorder = pre;
        inorder = in;
        int n = pre.length;
        for (int i = 0; i < inorder.length; i++) {
            indexMap.put(inorder[i], i);
        }

        return myBuildTree(0, n - 1, 0, n - 1);
    }

    public TreeNode myBuildTree(int pl, int pr, int il, int ir) {
        if(pl > pr){
            return null;
        }

        int inRootIndex = indexMap.get(preorder[pl]);
        int leftTreeNum = inRootIndex - il; // 计算左子树的节点数量

        TreeNode root = new TreeNode(preorder[pl]);

        TreeNode left = myBuildTree(pl + 1, pl + leftTreeNum, il, inRootIndex - 1);
        TreeNode right = myBuildTree(pl + 1 + leftTreeNum, pr, inRootIndex + 1, ir);
        root.left = left;
        root.right = right;

        return root;
    }
}
```

--- 
6. [序列化与反序列化二叉树](https://leetcode.cn/problems/serialize-and-deserialize-bst/description/?envType=problem-list-v2&envId=depth-first-search)
```java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode(int x) { val = x; }
 * }
 */
public class Codec {
    Map<String, Integer> indexMap;
    String[] preS;
    String[] inS;

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sbPre = preOrder(root, new StringBuilder());
        StringBuilder sbIn = inOrder(root, new StringBuilder());

        return sbPre + ";" + sbIn;
    }

    public StringBuilder preOrder(TreeNode root, StringBuilder sb) {
        if (root == null) {
            return sb;
        }

        sb.append(root.val).append(",");
        preOrder(root.left, sb);
        preOrder(root.right, sb);

        return sb;
    }

    public StringBuilder inOrder(TreeNode root, StringBuilder sb) {
        if (root == null) {
            return sb;
        }

        inOrder(root.left, sb);
        sb.append(root.val).append(",");
        inOrder(root.right, sb);

        return sb;
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        indexMap = new HashMap<>();
        String[] two = data.split(";");
        if (two.length == 0) {
            return null;
        }

        preS = two[0].split(",");
        inS = two[1].split(",");

        for (int i = 0; i < inS.length; i++) {
            indexMap.put(inS[i], i);
        }
        return decode(0, preS.length-1, 0, inS.length-1);
    }

    public TreeNode decode(int preLeft, int preRight, int inLeft, int inRight) {
        if (preLeft > preRight) {
            return null;
        }

        int inSIndex = indexMap.get(preS[preLeft]);
        int leftSize = inSIndex - inLeft;
        TreeNode root = new TreeNode(Integer.parseInt(inS[inSIndex]));
        root.left = decode(preLeft + 1, preLeft + leftSize , inLeft, inSIndex - 1);
        root.right = decode(preLeft + 1 + leftSize, preRight, inSIndex + 1, inRight);

        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec ser = new Codec();
// Codec deser = new Codec();
// String tree = ser.serialize(root);
// TreeNode ans = deser.deserialize(tree);
// return ans;
```