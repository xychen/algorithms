## 三种解题思路

### 1. 直接判断进行二分查找

<https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode/167103>

### 2. 把mid下标还原到原数组

- <https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0033.search-in-rotated-sorted-array/search-in-rotated-sorted-array.go>
- 需要先寻找旋转点（既然已经找到旋转点，其实按方法3效率更高效）


### 3. 找到旋转点，缩小范围，在2个递增数组中进行二分查找

- 参考《剑指offer》中寻找旋转点的讲解