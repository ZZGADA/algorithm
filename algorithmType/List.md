# 链表  
1. LRU [LRU](https://leetcode.cn/problems/lru-cache/)
```java
class LRUCache {
    private static class Node {
        int key, value;
        Node prev, next;

        Node(int k, int v) {
            key = k;
            value = v;
            
        }
    }

    private final int capacity;
    private final Node dummy = new Node(0, 0); // 哨兵节点 头指针
    private final Map<Integer, Node> keyToNode = new HashMap<>();

    // 关键问题：如何逐出最久未使用的关键字 
    // 以及注意：keyToNode只负责维护key和node的关系 但是 访问次数的顺序是node自身的双向链表决定的
    public LRUCache(int capacity) {
        this.capacity = capacity;
        dummy.prev = dummy;
        dummy.next = dummy;     // !!!! 这两步很重要 形成环形链表
    }
    
    // 获取数据 同时更新元素的头
    public int get(int key) {
        Node res =this.getNode(key);
        return res == null ? -1 : res.value;
    }
    

    public void put(int key, int value) {
        Node node = getNode(key);
        if(node == null){
            // 更新新的
            Node n = new Node(key,value);
            keyToNode.put(key,n);
            this.pushFront(n);  // 推到最顶

            if(keyToNode.size() > this.capacity){
                // 淘汰旧的
                Node backNode = dummy.prev;    // 尾巴
                keyToNode.remove(backNode.key);
                remove(backNode);
            }
        }else{
            node.value = value;
            return ;
        }
    }

    private Node getNode(int key){
        if(!keyToNode.containsKey(key)){
            return null;
        }

        Node node = keyToNode.get(key);  // 结点存在
        remove(node);                    // 将结点删除掉
        pushFront(node);                 // 结点插入头
        
        return  node ;
    }

    // 删除结点 双向指针移动
    private void remove(Node x) {
        x.prev.next = x.next;
        x.next.prev = x.prev;
    }

    // 头结点插入并更新
    private void pushFront(Node x){
        x.next = this.dummy.next;
        x.prev = this.dummy;

        x.next.prev = x;
        x.prev.next = x;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */


```


--- 
2. 旋转链表 [旋转链表](https://leetcode.cn/problems/rotate-list/)
```java

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        int len = 0;
        ListNode node = head;
        ListNode pre = head;
        ListNode next = head;
        ListNode res = head;
        // 统计总长度
        while(node != null){
            len++;
            node = node.next;
        }

        if (len == 0){
            return head;
        }

        // head 需要移动几步 
        int step = k % len;
        int count = 0;
        if(step == 0){
            return head;
        }
        
        node = head;
        while(node!=null){
            count++;
            if(len - count == step){
                // 边界条件 
                res = node.next;
                node.next = null;
                node = res;
                pre = node;
                continue;
            }
            pre = node;
            node = node.next;
        }
        pre.next = head;
        return res;
    }
}

```

