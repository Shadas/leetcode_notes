Given the root of a binary tree, return the leftmost value in the last row of the tree.



Example 1:

```
Input: root = [2,1,3]
Output: 1
```

Example 2:

```
Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7
```

Constraints:

 - The number of nodes in the tree is in the range [1, 104].
 - -231 <= Node.val <= 231 - 1

----

思路1：<br>
深度优先，对比左右子树最大深度<br>
思路2：<br>
广度优先层序遍历，记录每一层左侧第一个<br>

----

leetcode 513