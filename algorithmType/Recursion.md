# 递归



1. 快速幂

```java
class Solution {
    // 将3^10转换为9^5
    public double myPow(double x, int n) {
        if (x == 0) return 0;
        if (x == 1 || n == 0) return 1;

        long N = n; // 将n转换为long类型 防止溢出
        if (N < 0) {
            x = 1 / x;
            N = -N;
        }

        double res = 1.0;
        while (N > 0) {
            if (N % 2 == 1) {
                res *= x;
            }
            x *= x;
            N /= 2;
        }

        return res;
    }
}
```