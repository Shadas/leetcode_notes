Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.



Example 1:

```
Input: nums = [1,1,1], k = 2
Output: 2
```

Example 2:

```
Input: nums = [1,2,3], k = 3
Output: 2
```


Constraints:

 - 1 <= nums.length <= 2 * 104
 - -1000 <= nums[i] <= 1000
 - -107 <= k <= 107

----

leetcode 560

暴力解法尝试不予考虑

因为包含负数，所以滑动窗口思路不适用

连续子串和为目标值，则说明从初始位置到子串头尾位置的总和之差正好为目标值