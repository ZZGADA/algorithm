# 前缀和

1. 24 mod [构成整天的下标对数目](https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-i/?envType=daily-question&envId=2024-10-22)
```java
class Solution {
   public int countCompleteDayPairs(int[] hours) {
      // 前缀和问题 ( hours[i] + hours[j] ) mod 24 == （hours[i] mod 24 + hours[j] mod 24 ）
      // mod 24 == 0
      // 注意⚠️： HashMap可以转换成int[24] 因为mod的范围就这么大
      int length = hours.length;
      int res = 0;
      Map<Integer, Integer> map = new HashMap<Integer, Integer>();

      for (int i = 0; i < length; i++) {
         int mod = hours[i] % 24;
         int need = (24 - mod) % 24;
         res += map.getOrDefault(need, 0);
         map.put(mod, map.getOrDefault(mod, 0)+1);
      }
      return res;
   }
}
```