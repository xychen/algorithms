我的算法练习册
-----------------------

- [我的算法练习册](#我的算法练习册)
- [题目分类](#题目分类)
  - [二叉树](#二叉树)
  - [回溯](#回溯)
  - [动态规划](#动态规划)
  - [贪心算法](#贪心算法)
  - [数组、链表、哈希表、字符串](#数组链表哈希表字符串)
- [参考图书](#参考图书)
- [其它参考](#其它参考)

## 题目分类

### 二叉树
[二叉树leetcode题单](https://leetcode-cn.com/problem-list/xzLo320B/)
更多的题目：[二叉树leetcode题单-labuladong专项训练](https://leetcode-cn.com/problem-list/jvxH8e0C/)

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|二叉树遍历|[144. 二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)|✅||递归、迭代，注意迭代的写法：右孩子先进栈|
|☆二叉树遍历|[94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)|✅||递归、迭代，注意迭代的写法|
|二叉树遍历|[145. 二叉树的后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)|✅||递归、迭代，注意迭代的写法：左孩子先进栈，最后要翻转|
|二叉树遍历|[102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|✅||队列|
|二叉树遍历|[226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)|✅||前序or后序遍历，不能用中序|
|二叉树遍历|[101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/)|✅||compare:左外vs右外、左内vs右内|
||[104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|✅||1+max(左子树深度,右子树深度)，也可以用层次遍历|
|☆|[111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)|✅||此题递归解法有陷阱，1+min(minDepth(左),minDepth(右)) <br/>如果左、右孩子都为nil，返回1，如果左为nil，返回右的最小深度，如果右为nil，返回左的最小深度，如果左右都不为nil，返回 1+min(minDepth(左),minDepth(右))；<br/>迭代解法：层次遍历，遇到第一个叶子节点即是最小深度|
|☆|[110. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)|✅||获取深度，如果不平衡返回-1，递归中判断如果有-1的情况直接返回，否则返回1+max(左子树深度,右子树深度)，也可以用层次遍历|
|DFS|[257. 二叉树的所有路径](https://leetcode-cn.com/problems/binary-tree-paths/)|✅||回溯|
|DFS|[112. 路径总和](https://leetcode-cn.com/problems/path-sum/)|✅||回溯|
|DFS|[113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)|✅||回溯|
|构造二叉树|[106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)|✅||注意边界即可|
|构造二叉树|[105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)|✅||注意边界即可|
|☆遍历|[617. 合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/)|✅||同时遍历2棵树，孩子nil节点处理，root.Val相加|
|二分查找|[700. 二叉搜索树中的搜索](https://leetcode-cn.com/problems/search-in-a-binary-search-tree/)|✅|||
|☆中序遍历+pre节点|[98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)|✅||迭代、递归都可，注意如何保存pre节点|
|☆中序遍历+pre节点|[530. 二叉搜索树的最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)|✅||迭代、递归都可，注意如何保存pre节点|
|☆中序遍历+pre节点|[501. 二叉搜索树中的众数](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)|✅||迭代、递归都可，注意如何保存pre节点，注意计数的变化、何时清空结果集|
|☆|[236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)|✅||注意递归终止条件包含root = p or q;左边找一下返回n1,右边找一下返回n2，处理n1、n2同时为nil和有一个为nil的场景|
||[235. 二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)|✅|||
||[701. 二叉搜索树中的插入操作](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)|✅||插入到叶子节点上|
|☆|[450. 删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/)|✅||递归函数有返回值，利用返回值；分为5中情况：1.未找到 2.是叶子节点，直接删 3.左孩子为空，右孩子补到当前节点上 4.右孩子为空，左孩子补到当前节点上 5.左右孩子都不为空，左孩子插到右孩子的最左边|
|☆|[669. 修剪二叉搜索树](https://leetcode-cn.com/problems/trim-a-binary-search-tree/)|✅||递归函数有返回值，注意利用返回值；root.Val < low时，左孩子都不要了；root.Val > right时，右孩子都不要了|
||[108. 将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)|✅||类似二分查找|


### 回溯
[回溯leetcode题单](https://leetcode-cn.com/problem-list/v7gtwrSv)

### 动态规划
[动态规划leetcode题单](https://leetcode-cn.com/problem-list/PgGHdyoW)
### 贪心算法
[贪心算法leetcode题单](https://leetcode-cn.com/problem-list/Gi5g2iZo)

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
||[455.分发饼干](https://leetcode-cn.com/problems/assign-cookies)|✅||easy|
||[376.摆动序列](https://leetcode-cn.com/problems/wiggle-subsequence)|✅|||
||[53.最大子序和](https://leetcode-cn.com/problems/maximum-subarray)|✅||easy|
||[122.买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii)|✅||1.可以使用贪心（找低谷和峰值）<br> 2.使用动态规划|
||[55.跳跃游戏](https://leetcode-cn.com/problems/jump-game)|✅||1.可以使用动态规划|
||[45.跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii)|✅||代码随想录的方法一好理解一点|
||[1005.K 次取反后最大化的数组和](https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations)|✅||easy|
||[134.加油站](https://leetcode-cn.com/problems/gas-station)|✅||有点小技巧|
||[135.分发糖果](https://leetcode-cn.com/problems/candy)|||hard|
||[860.柠檬水找零](https://leetcode-cn.com/problems/lemonade-change)|✅||easy|
||[406.根据身高重建队列](https://leetcode-cn.com/problems/queue-reconstruction-by-height)|✅||有点难度|
||[452.用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons)|✅||无交集的数量 => 放箭的数量|
||[435.无重叠区间](https://leetcode-cn.com/problems/non-overlapping-intervals)|✅||总集合数-无交集的数量 =>移除的数量|
||[763.划分字母区间](https://leetcode-cn.com/problems/partition-labels)|✅|||
||[56.合并区间](https://leetcode-cn.com/problems/merge-intervals)|✅|||
||[738.单调递增的数字](https://leetcode-cn.com/problems/monotone-increasing-digits)|✅|||
||[714.买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)|✅||参考：[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/) 的动归解法<br>1.可以使用贪心（难理解）<br>2.使用动态规划|
||[968.监控二叉树](https://leetcode-cn.com/problems/binary-tree-cameras)|✅||hard|

### 数组、链表、哈希表、字符串

[数组、链表、哈希表、字符串leetcode题单]()
|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|数组|[704.二分查找](https://leetcode-cn.com/problems/binary-search)|✅||easy|
|数组|[34.在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array)|✅||二分查找变种题，不用判断边界，使用pre变量|
|数组|[27.移除元素](https://leetcode-cn.com/problems/remove-element)|✅||easy|
|数组|[977.有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array)|✅||easy|
|数组|[209.长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum)|✅||双指针|
|数组|[59.螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii)|✅|||
|数组|[31.下一个排列](https://leetcode-cn.com/problems/next-permutation/)|✅||跟排列有点关系|
|数组|[169.多数元素](https://leetcode-cn.com/problems/majority-element)|✅|||
|----|||||
|链表|[203.移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements)|✅||easy|
|链表|[707.设计链表](https://leetcode-cn.com/problems/design-linked-list)|✅|||
|链表|[206.反转链表](https://leetcode-cn.com/problems/reverse-linked-list)|✅||easy|
|链表|[24.两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs)|✅|||
|链表|[19.删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list)|✅||fast走n+1步|
|链表|[面试题 02.07.链表相交](https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci)|✅||easy|
|链表|[142.环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii)|✅|||
|----|||||
|哈希表|[242.有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram)|✅||easy|
|哈希表|[349.两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays)|✅||easy|
|哈希表|[202.快乐数](https://leetcode-cn.com/problems/happy-number)|✅||easy|
|哈希表|[1.两数之和](https://leetcode-cn.com/problems/two-sum)|✅||easy|
|哈希表|[454.四数相加 II](https://leetcode-cn.com/problems/4sum-ii)|✅|||
|哈希表|[383.赎金信](https://leetcode-cn.com/problems/ransom-note)|✅||easy|
|----|||||
|字符串|[344.反转字符串](https://leetcode-cn.com/problems/reverse-string)|✅||easy<br/>使用双指针|
|字符串|[541.反转字符串 II](https://leetcode-cn.com/problems/reverse-string-ii)|✅|||
|字符串|[剑指 Offer 05.替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof)|✅||easy|
|字符串|[151.翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string)|✅|||
|字符串|[剑指 Offer 58 - II.左旋转字符串](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof)|✅||easy|
|字符串|[28.实现 strStr()](https://leetcode-cn.com/problems/implement-strstr)|||KMP|
|字符串|[459.重复的子字符串](https://leetcode-cn.com/problems/repeated-substring-pattern)|||KMP|

## 参考图书
- [《代码随想录》](https://book.douban.com/subject/35680544/)
- [《labuladong的算法小抄》](https://book.douban.com/subject/35252621/)
- [《剑指Offer：名企面试官精讲典型编程题（第2版）》](https://book.douban.com/subject/27008702/)
- [《编程之美》](https://book.douban.com/subject/3004255/)

## 其它参考
- [代码随想录](https://programmercarl.com/)
- [labuladong的算法小抄](https://labuladong.github.io/algo/)
- [花花酱题单](https://docs.google.com/spreadsheets/d/1SbpY-04Cz8EWw3A_LBUmDEXKUMO31DBjfeMoA0dlfIA/htmlview#)
- [codetop](https://codetop.cc/home)
- [LeetCode中文版：https://leetcode-cn.com/](https://leetcode-cn.com/)
- [https://github.com/aQuaYi/LeetCode-in-Go](https://github.com/aQuaYi/LeetCode-in-Go)
