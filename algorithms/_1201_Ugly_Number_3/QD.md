Write a program to find the n-th ugly number.

Ugly numbers are positive integers which are divisible by a or b or c.

 

Example 1:

```
Input: n = 3, a = 2, b = 3, c = 5
Output: 4
Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.
```

Example 2:

```
Input: n = 4, a = 2, b = 3, c = 4
Output: 6
Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.
```

Example 3:

```
Input: n = 5, a = 2, b = 11, c = 13
Output: 10
Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.
```

Example 4:

```
Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
Output: 1999999984
```

Constraints:

 - 1 <= n, a, b, c <= 10^9
 - 1 <= a * b * c <= 10^18
 - It's guaranteed that the result will be in range [1, 2 * 10^9]

----

leetcode 264 丑数2 的提高版本，数据量会有所提高。不能用2的依次计算累加，可以用二分查找定位的方式。

也即，计算对于任意数k，在[0,k]内，满足条件的丑数有多少个。

计算方案使用最小公倍数、最大公约数及容斥原理实现。
至于具体实现，这是一道数学题了。。