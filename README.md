我的算法练习册
-----------------------

- [我的算法练习册](#我的算法练习册)
- [题目分类](#题目分类)
  - [二叉树](#二叉树)
  - [回溯](#回溯)
  - [动态规划](#动态规划)
  - [贪心算法](#贪心算法)
  - [数组、链表、哈希表、字符串](#数组链表哈希表字符串)
  - [栈和队列](#栈和队列)
  - [双指针](#双指针)
  - [图](#图)
  - [其他](#其他)
- [参考图书](#参考图书)
- [其它参考](#其它参考)

## 题目分类

### 二叉树
[二叉树leetcode题单](https://leetcode-cn.com/problem-list/xzLo320B/) <br/>
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

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|组合:选k个数|[77. 组合](https://leetcode-cn.com/problems/combinations/)|✅||1.startIndex<br/> 2.剪枝优化：已选择+剩余元素<k时|
|组合:从1-9中选和为n的k个数|[216. 组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii/)|✅||1.startIndex<br/> 2.剪枝优化：已选择+剩余元素<k时|
|组合:从候选数组中选和为target的所有组合|[39. 组合总和](https://leetcode-cn.com/problems/combination-sum/)|✅||1.startIndex<br/> 2.对target做减操作，则不用累加 <br/> 3.剪枝：排序+如果target减操作后小于0，则跳出循环|
|组合:从重复元素的候选数组中选和为target的所有组合|[40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/)|✅||1.startIndex<br/> 2.对target做减操作，则不用累加<br/> 3.排序+同层去重 <br/> 4.剪枝:如果target减操作后小于0，则跳出循环|
||[17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)|✅|||
||[131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)|✅||1.startIndex<br/> 2.判断回文|
||[93. 复原 IP 地址](https://leetcode-cn.com/problems/restore-ip-addresses/)|✅||1.选择3个有效数字后，直接判断剩余的1个是否合法，可作为递归终止条件|
|子集|[78. 子集](https://leetcode-cn.com/problems/subsets/)|✅||1.startIndex|
|子集:输入数组有重复元素|[90. 子集 II](https://leetcode-cn.com/problems/subsets-ii/)|✅||1.startIndex <br/> 2.排序+同层去重|
||[491. 递增子序列](https://leetcode-cn.com/problems/increasing-subsequences/)|✅||1.startIndex <br/> 2.同层去重使用局部变量|
|全排列|[46. 全排列](https://leetcode-cn.com/problems/permutations/)|✅||1.标记数组|
|全排列|[47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)|✅||1.标记数组 <br/> 2.同层去重|
||[51. N 皇后](https://leetcode-cn.com/problems/n-queens/)|✅|||
||[37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/)||||
||[332. 重新安排行程](https://leetcode-cn.com/problems/reconstruct-itinerary/)||||

### 动态规划
[动态规划leetcode题单](https://leetcode-cn.com/problem-list/WgkqSsgU/)
更多题目：[动态规划leetcode题单-花花酱](https://leetcode-cn.com/problem-list/PgGHdyoW)

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
||[509. 斐波那契数](https://leetcode-cn.com/problems/fibonacci-number/)|✅|||
||[70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/submissions/)|✅|||
||[746. 使用最小花费爬楼梯](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)|✅||第一步不需要花费体力值，最后一步需要花费|
||[62. 不同路径](https://leetcode-cn.com/problems/unique-paths/submissions/)|✅||二维dp，可压测成2行；还有数论方法|
||[63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)|✅||必须使用二维dp|
|☆|[343. 整数拆分](https://leetcode-cn.com/problems/integer-break/)|✅||最值是 dp[i]=max(dp[i], dp[i-j]*j, (i-j)*j), 不要遗忘(i-j)*j<br/>此题有贪心解法|
||[96. 不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/)|✅||选择一个根节点，左子树的数量*右子树的数量，注意是乘法|
|0-1背包|[416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/)|✅||转化成挑选元素等于total/2的背包问题，物品重量和价值都是nums[i],如果最终重量等于价值，则存在|
|☆0-1背包|[494. 目标和](https://leetcode-cn.com/problems/target-sum/)|✅||解法思路： left - right = target  =>  left - (sum-left) = target  =>  left = (sum+target)/2 ，即从数组中选出和为left的有多少种方法。<br/>dp数组的定义和普通背包问题略有不同|
|☆0-1背包|[474. 一和零](https://leetcode-cn.com/problems/ones-and-zeroes/)|✅||背包重量是二维的|
|完全背包，求组合数|[518. 零钱兑换 II](https://leetcode-cn.com/problems/coin-change-2/)|✅||如果单纯问“能否”凑成总和，遍历物品、遍历背包容量是没有顺序要求的。但是本题求组合数<br/>dp[j]：兑换金额为 amount 的时候，有多少种组合，等于所有 dp[j-coin] 之和|
|完全背包，求排列数|[377. 组合总和 Ⅳ](https://leetcode-cn.com/problems/combination-sum-iv/)|✅||本题求排列数<br/>如果求组合数，外层for循环遍历物品、内层for循环遍历背包<br/>如果求排列数，外层for循环遍历背包、内层for循环遍历物品|
||~~爬楼梯进阶版~~||||
|完全背包，最少硬币数|[322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)|✅||求最小，组合和排列都可以，物品在外层效率高一些|
|完全背包，最少数量|[279. 完全平方数](https://leetcode-cn.com/problems/perfect-squares/)|✅||和求最少硬币数一样|
|完全背包，求排列|[139. 单词拆分](https://leetcode-cn.com/problems/word-break/)|✅||完全背包问题求排列，所以背包重量在外层循环|
|股票买卖|[121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|✅||dp[i][0] 表示第i天持有股票的最大收益（当天买入或者前一天就持有）， dp[i][1]表示不持有股票的最大收益（当天卖出，或者前一天就没持有)<br/>可用贪心算法：找曲线最低点和最高点|
|股票买卖,可多次交易|[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|✅||dp[i][0] 表示第i天持有股票的最大收益， dp[i][1]表示不持有股票的最大收益<br/>可用贪心算法：多次找曲线最低点和最高点|
|股票买卖,最多2次交易|[123. 买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|✅||hard，有点绕|
|股票买卖,最多k次交易|[188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|✅||hard,可以从2次交易的推导过来|
|股票买卖,有冷冻期|[309. 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)|✅||4种状态|
|股票买卖,有手续费|[714. 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)|✅|||从多次交易中推导过来|
|最长问题，输入单个数组|[300.最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)|✅||||
|最长问题，输入单个数组|[674. 最长连续递增序列](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/)|✅|||只需要判断nums[i]和nums[i-1]<br/>有贪心解法|
|编辑距离，最长问题，输入2个字符串(数组)|[718. 最长重复子数组](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/)|✅|||二维dp，滚动数组注意清0|
|编辑距离(只删)，最长问题|[1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)|✅|||二维dp，text1[i-1]==text2[j-1]时，dp来源是dp[i-1][j-1]+1，否则来源是左or上的最大值|
|编辑距离(只删)，最长问题|[1035.不相交的线](https://leetcode-cn.com/problems/uncrossed-lines/)|✅|||转换成求最长公共子序列|
||[53. 最大子数组和](https://leetcode-cn.com/problems/maximum-subarray/)|✅||||
|编辑距离(只删)|[392.判断子序列](https://leetcode-cn.com/problems/is-subsequence/)|✅|||转换成求最长公共子序列，如果长度等于s，则是他的子序列<br/>编辑距离中删除的思想<br/>有双指针解法|
|编辑距离(只删)|[115.不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/)|✅|||hard，好难理解<br/> s[i-1]==t[j-1]时，dp[i][j]=dp[i-1][j-1]+dp[i-1][j],s[i-1] != t[j-1]时，dp[i][j]=dp[i-1][j]|
|编辑距离(只删)|[583. 两个字符串的删除操作](https://leetcode-cn.com/problems/delete-operation-for-two-strings/)|✅|||求最小的删除次数，word1[i-1]==word2[j-1]时，dp[i][j]=dp[i-1][j-1],word1[i-1]!=word2[j-1]时, dp[i][j]=min(dp[i-1][j]+1, dp[i][j-1])+1, dp[i-1][j-1]+2)|
|编辑距离(增、删、改)|[72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/)|✅|||word1[i-1]==word2[j-1]时，dp[i][j]=dp[i-1][j-1],word1[i-1]!=word2[j-1]时, dp[i][j]=min(dp[i-1][j]+1, dp[i][j-1])+1, dp[i-1][j-1]+1)|
||[647. 回文子串](https://leetcode-cn.com/problems/palindromic-substrings/)|||<br/>还有双指针的解法|


### 贪心算法
[贪心算法leetcode题单](https://leetcode-cn.com/problem-list/Gi5g2iZo)

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
||[455.分发饼干](https://leetcode-cn.com/problems/assign-cookies)|✅||easy，排序|
||[376.摆动序列](https://leetcode-cn.com/problems/wiggle-subsequence)|✅||找峰、谷|
||[53.最大子序和](https://leetcode-cn.com/problems/maximum-subarray)|✅||easy|
||[122.买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii)|✅||1.可以使用贪心（找低谷和峰值）<br> 2.使用动态规划|
||[55.跳跃游戏](https://leetcode-cn.com/problems/jump-game)|✅||使用cover解，注意终止条件|
||[45.跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii)|✅||使用cover、nextCover，转化的时候+1操作|
||[1005.K 次取反后最大化的数组和](https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations)|✅||按照绝对值排序，从后往前遍历，将负数取反；如果k还大于0且k%2==1，则取最小的正数取反|
||[134.加油站](https://leetcode-cn.com/problems/gas-station)|✅||大前提：油量和>使用和；i从0累加rest[i](gas[i]-cost[i]),和为curSum，一旦curSum小于0说明[0,i]区间的位置不能作为起始位置，起始维中从i+1开始，再重新计算curSum。为什么从i+1开始？i之前有多少负数，i之后就有对应的正数（因为大前提）|
||[135.分发糖果](https://leetcode-cn.com/problems/candy)|||hard，初始化结果数组为1，先从左往右遍历，如果右边孩子评分搞，则candyVec[i]=candyVec[i-1]+1,然后从右往左遍历，如果左边孩子评分高，则candyVec[i]=max(candyVec[i+1]+1, candyVec[i])，最后糖果数量相加|
||[860.柠檬水找零](https://leetcode-cn.com/problems/lemonade-change)|✅||easy|
||[406.根据身高重建队列](https://leetcode-cn.com/problems/queue-reconstruction-by-height)|✅||有点难度|
||[452.用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons)|✅||无交集的数量 => 放箭的数量，先排序|
|☆|[435.无重叠区间](https://leetcode-cn.com/problems/non-overlapping-intervals)|✅||有点难度，按照右边界排序，前一个的end和下一个的start没有重合，说明非交叉，每次取非交叉区间的时候，都是可右边界最小的来做分割点。区间总数-非交叉区间的个数 => 要移除的区间个数|
||[763.划分字母区间](https://leetcode-cn.com/problems/partition-labels)|✅||1.先遍历一遍找到每个字母的最远边界；2.再次遍历，如果下标和最远边界相等，则是一个分界点|
||[56.合并区间](https://leetcode-cn.com/problems/merge-intervals)|✅||排序+合并|
||[738.单调递增的数字](https://leetcode-cn.com/problems/monotone-increasing-digits)|✅||1.如果strNum[i-1]>strNum[i],则strNum[i-1]=strNum[i-1]-1, strNum[i]=9;2.要从后往前遍历;3.一旦出现i变9的情况，i后边的数字都要直接变9|
||[714.买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)|✅||参考：[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/) 的动归解法<br>1.可以使用贪心（难理解）<br>2.使用动态规划|
||[968.监控二叉树](https://leetcode-cn.com/problems/binary-tree-cameras)|||hard|

### 数组、链表、哈希表、字符串

[数组、链表、哈希表、字符串leetcode题单]()
|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|数组|[704.二分查找](https://leetcode-cn.com/problems/binary-search)|✅||easy|
|数组|[34.在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array)|✅||二分查找变种题，不用判断边界，使用pre变量|
|数组|[27.移除元素](https://leetcode-cn.com/problems/remove-element)|✅||easy|
|数组|[977.有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array)|✅||easy|
|数组|[209.长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum)|✅||双指针|
|数组|[59.螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii)|✅||坚持所有循环都是左开右闭|
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
|哈希表|[15. 三数之和](https://leetcode-cn.com/problems/3sum/)|✅||通用解法，参考：https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q|
|哈希表|[18. 四数之和](https://leetcode-cn.com/problems/4sum/)|✅||通用解法，参考：https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q|
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

### 栈和队列

[栈和队列leetcode题单](https://leetcode-cn.com/problem-list/eDUdhG7b)
|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
||[232. 用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/)|✅|||
||[225. 用队列实现栈](https://leetcode-cn.com/problems/implement-stack-using-queues/)|✅|||
|栈|[20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)|✅|||
|栈|[150. 逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)|✅|||
|单调队列|[239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)|✅||hard,push x的时候，将x与队尾元素比较，如果队尾元素小于x，则不断删除队尾元素，使队列保持单调递减；pop的时候将nums[i-k+1]和队头元素比较，相等时才删除|
|单调栈|[496. 下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)|✅|||
|单调栈|[739. 每日温度](https://leetcode-cn.com/problems/daily-temperatures/)|✅|||
|单调栈|[503. 下一个更大元素 II](https://leetcode-cn.com/problems/next-greater-element-ii/)|✅|||
|优先队列|[347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)|✅||前k个高频元素（不是大小的top k），使用partition的思想，堆的方式代码不好记|


### 双指针
[双指针leetcode题单](https://leetcode-cn.com/problem-list/hRlpbXao)
|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|双指针|[42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)|✅||方法一：提前计算好lmax、rmax数组，空间复杂度O(n)；方法二：同步计算lmax和rmax，空间复杂度O(1)|
|双指针|[11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)|✅||计算面积，求最大|

### 图

[图leetcode题单](https://leetcode-cn.com/problem-list/AOJgRUlL)

|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
||[133. 克隆图](https://leetcode-cn.com/problems/clone-graph/)|||队列+hashtable|
||[138. 复制带随机指针的链表](https://leetcode-cn.com/problems/copy-list-with-random-pointer/)|||队列+hashtable|
|岛屿数量|[200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)|✅||DFS，遇到岛屿则全淹没|
|岛屿数量|~~ [694. 不同岛屿的数量](https://leetcode-cn.com/problems/number-of-distinct-islands/) ~~||||
|岛屿数量|[695. 岛屿的最大面积](https://leetcode-cn.com/problems/max-area-of-island/)|✅||DFS|
|岛屿数量|[1254. 统计封闭岛屿的数目](https://leetcode-cn.com/problems/number-of-closed-islands/)|✅||先把4个边界淹掉，在计算岛屿数量，则是封闭的岛屿数量|
|岛屿数量|[1905. 统计子岛屿](https://leetcode-cn.com/problems/count-sub-islands/)|✅||先把grid2中肯定不是子岛屿（grid2[i][j] == 1 && grid1[i][j] == 0）的岛屿淹掉|
|并查集|[547. 省份数量](https://leetcode-cn.com/problems/number-of-provinces/)|||图+连通分量|
||[733. 图像渲染](https://leetcode-cn.com/problems/flood-fill/)|||图+连通分量|
||[827. 最大人工岛](https://leetcode-cn.com/problems/making-a-large-island/)|||图+连通分量|
||[1020.飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/)||||
||[1162. 地图分析](https://leetcode-cn.com/problems/as-far-from-land-as-possible/)|||图+连通分量|
||[841. 钥匙和房间](https://leetcode-cn.com/problems/keys-and-rooms/)|||DFS+连通分量|
||[1202. 交换字符串中的元素](https://leetcode-cn.com/problems/smallest-string-with-swaps/)|||DFS+连通分量|
||[207. 课程表](https://leetcode-cn.com/problems/course-schedule/)|||拓扑排序|
||[210. 课程表 II](https://leetcode-cn.com/problems/course-schedule-ii/)|||拓扑排序|
||[802. 找到最终的安全状态](https://leetcode-cn.com/problems/find-eventual-safe-states/)|||拓扑排序|
||[399. 除法求值](https://leetcode-cn.com/problems/evaluate-division/)|||并查集|
||[839. 相似字符串组](https://leetcode-cn.com/problems/similar-string-groups/)|||并查集|
||[952. 按公因数计算最大组件大小](https://leetcode-cn.com/problems/largest-component-size-by-common-factor/)|||并查集|
||[990. 等式方程的可满足性](https://leetcode-cn.com/problems/satisfiability-of-equality-equations/)|||并查集|
||[721. 账户合并](https://leetcode-cn.com/problems/accounts-merge/)|||并查集|
||[785. 判断二分图](https://leetcode-cn.com/problems/is-graph-bipartite/)|||二分图+图着色|
||[886. 可能的二分法](https://leetcode-cn.com/problems/possible-bipartition/)||||
||[1042. 不邻接植花](https://leetcode-cn.com/problems/flower-planting-with-no-adjacent/)||||
||[997. 找到小镇的法官](https://leetcode-cn.com/problems/find-the-town-judge/)||||
||[433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)|||未加权的最短路径+BFS|
||[815. 公交路线](https://leetcode-cn.com/problems/bus-routes/)|||未加权的最短路径+BFS|
||[863. 二叉树中所有距离为 K 的结点](https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/)|||未加权的最短路径+BFS|
||[1129. 颜色交替的最短路径](https://leetcode-cn.com/problems/shortest-path-with-alternating-colors/)|||未加权的最短路径+BFS|
||[1263. 推箱子](https://leetcode-cn.com/problems/minimum-moves-to-move-a-box-to-their-target-location/)|||未加权的最短路径+BFS|
||[684. 冗余连接](https://leetcode-cn.com/problems/redundant-connection/)|||圈,并查集|
||[685. 冗余连接 II](https://leetcode-cn.com/problems/redundant-connection-ii/)|||圈,并查集|
||[1319. 连通网络的操作次数](https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/)|||圈,并查集|
||[743. 网络延迟时间](https://leetcode-cn.com/problems/network-delay-time/)|||加权最短路径|
||[787. K 站中转内最便宜的航班](https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/)|||加权最短路径|
||[882. 细分图中的可到达结点](https://leetcode-cn.com/problems/reachable-nodes-in-subdivided-graph/)|||加权最短路径|
||[924. 尽量减少恶意软件的传播](https://leetcode-cn.com/problems/minimize-malware-spread/)|||加权最短路径|
||[1334. 阈值距离内邻居最少的城市](https://leetcode-cn.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/)|||加权最短路径|
||[847. 访问所有节点的最短路径](https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/)|||BFS|
||[864. 获取所有钥匙的最短路径](https://leetcode-cn.com/problems/shortest-path-to-get-all-keys/)|||BFS|
||[1298. 你能从盒子里获得的最大糖果数](https://leetcode-cn.com/problems/maximum-candies-you-can-get-from-boxes/)|||BFS|
||[332. 重新安排行程](https://leetcode-cn.com/problems/reconstruct-itinerary/)|||欧拉路径|
||[1192. 查找集群内的「关键连接」](https://leetcode-cn.com/problems/critical-connections-in-a-network/)|||强连通分量|
||[943. 最短超级串](https://leetcode-cn.com/problems/find-the-shortest-superstring/)|||哈密顿路径|
||[980. 不同路径 III](https://leetcode-cn.com/problems/unique-paths-iii/)|||哈密顿路径|
||[996. 正方形数组的数目](https://leetcode-cn.com/problems/number-of-squareful-arrays/)|||哈密顿路径|
||[959. 由斜杠划分区域](https://leetcode-cn.com/problems/regions-cut-by-slashes/)|||并查集/图+CCs|

### 其他

[其他类型leetcode题单](https://leetcode-cn.com/problem-list/ezM46lve)
|分类|相关题号|done|代码|备注|
|----|----|----|----|----|
|前缀树|[208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)|✅|||
|前缀树|[648. 单词替换](https://leetcode-cn.com/problems/replace-words/)||||
|前缀树|[211. 添加与搜索单词 - 数据结构设计](https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/)||||
|前缀树|[677. 键值映射](https://leetcode-cn.com/problems/map-sum-pairs/)||||
||[]()||||


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
