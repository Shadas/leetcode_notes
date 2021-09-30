Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.


Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

----

具有有序性，用二分解。
其实重点考察二分写法。

思路1：<br>
可以先找行，再找列，但是效率不如直接二分<br>
思路2：<br>
可以把二维数组读到一个一维数组中，但是空间占用不是很好<br>
思路3：<br>
将idx直接转化为i，j的访问，多一步计算即可，还是直接二分，不用额外空间。