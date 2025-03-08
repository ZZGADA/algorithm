# 字符串题目

---

1. [比较版本号](https://leetcode.cn/problems/compare-version-numbers/description/)
```java
class Solution {
    public int compareVersion(String version1, String version2) {
        // 所有修订号都可以存储在32位整形中
        // 一个个修订号依次比较
        // 缺失的修订号按 "0" 处理。
        String[] v1 = version1.split("\\.");
        String[] v2 = version2.split("\\.");
        int lenV1 = v1.length, lenV2 = v2.length;
        int i = 0, j = 0;
        while (i < lenV1 || j < lenV2) {
            int num1 = 0, num2 = 0;
            if (i < lenV1) {
                num1 = Integer.parseInt(v1[i++]);
            }

            if (j < lenV2) {
                num2 = Integer.parseInt(v2[j++]);
            }

            if (num1 < num2) {
                return -1;
            } else if (num1 > num2) {
                return 1;
            }
        }

        return 0;
    }
}
```

---  
2. [实现Trie（前缀树）](https://leetcode.cn/problems/implement-trie-prefix-tree/description/) 
```java

class Alpha {
    public Alpha[] table;
    public boolean isEndOfAlpha;

    Alpha() {
        this.table = new Alpha[26];
        isEndOfAlpha = false;
    }
}

public class Trie {
    public Alpha alpha;

    public Trie() {
        this.alpha = new Alpha();
    }

    public void insert(String word) {
        insertIntoAlpha(word, 0, this.alpha);
    }

    public void insertIntoAlpha(String word, int index, Alpha alpha) {
        if (index >= word.length()) {
            alpha.isEndOfAlpha = true;
            return;
        }
        int charIndex = word.charAt(index) - 'a';
        if (alpha.table[charIndex] != null) {
            insertIntoAlpha(word, index + 1, alpha.table[charIndex]);
        } else {
            // alpha不存在
            alpha.table[charIndex] = new Alpha();
            insertIntoAlpha(word, index + 1, alpha.table[charIndex]);
        }
    }

    public boolean search(String word) {
        return searchWithIndex(word, 0, this.alpha);
    }

    public boolean searchWithIndex(String word, int index, Alpha alpha) {
        if (index >= word.length()) {
            return alpha.isEndOfAlpha;
        }
        int charIndex = word.charAt(index) - 'a';
        if (alpha.table[charIndex] != null) {
            return searchWithIndex(word, index + 1, alpha.table[charIndex]);
        }else{
            return false;
        }
    }

    public boolean startsWith(String prefix) {
        return searchWithPrefix(prefix, 0, this.alpha);
    }

    public boolean searchWithPrefix(String word, int index, Alpha alpha) {
        if (index >= word.length()) {
            return true;
        }
        int charIndex = word.charAt(index) - 'a';
        if (alpha.table[charIndex] != null) {
            return searchWithPrefix(word, index + 1, alpha.table[charIndex]);
        }else{
            return false;
        }
    }
}


/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
``` 

--- 
--- 

3. [文件中的最长路径](https://leetcode.cn/problems/longest-absolute-file-path/?envType=problem-list-v2&envId=depth-first-search)
```java
class Solution {
    public int lengthLongestPath(String s) {
        // 指向 文件 的 最长绝对路径 的长度
        Map<Integer, String> map = new HashMap<>();
        int n = s.length();
        String ans = "";

        for (int i = 0; i < n;) {
            int level = 0;
            while (i < n && s.charAt(i) == '\t') {
                // 判断当前 文件或者文件夹是第几级的
                i++;
                level++;
            }

            int j = i;
            boolean isDir = true;
            while (j < n && s.charAt(j) != '\n') {
                if (s.charAt(j) == '.') {
                    isDir = false;
                }
                j++;
            }

            String cur = s.substring(i, j);
            String prev = map.getOrDefault(level - 1, null);
            String path = prev == null ? cur : prev + "/" + cur;

            if (isDir) {
                // if is dir
                // 将当前路径推入到 map中 
                // map永远记录只最新一级的路径
                map.put(level, path);
            } else {
                // is file
                ans = ans.length() > path.length() ? ans : path;
            }
            i = j + 1;
        }
        return ans.length();
    }
}
```
---- 