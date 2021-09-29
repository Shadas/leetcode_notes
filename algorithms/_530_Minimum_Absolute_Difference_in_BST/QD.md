Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.



Example 1:

```
Input: root = [4,2,6,1,3]
Output: 1
```

Example 2:

```
Input: root = [1,0,48,null,null,12,49]
Output: 1
```

Constraints:

- The number of nodes in the tree is in the range [2, 104].
- 0 <= Node.val <= 105


Note: 
- This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/

----

思路：<br>
BST其实是有序的，任意两个节点的最小差一定来自于相邻有序两个节点。<br>
因此，按照搜索序遍历找最小即可。

也可以遍历后放到有序数组，在数组中遍历，但是数组空间实际上没有必要。在遍历过程中即可计算出来。
