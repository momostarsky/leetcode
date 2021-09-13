package leetcode

import (
	"fmt"
	"sort"
)

/*
颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
进阶：
   你可以不使用代码库中的排序函数来解决这道题吗？
   你能想出一个仅使用常数空间的一趟扫描算法吗？

*/
func SortColors(nums []int) {
	var ln = len(nums)
	if ln <= 1 {
		return
	}

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}

func SortColorsWithScan(nums []int) {
	var ln = len(nums)
	if ln <= 1 {
		return
	}
	var sscaner = make(map[int]int)
	for _, v := range nums {
		if vv, mok := sscaner[v]; !mok {
			sscaner[v] = 1
		} else {
			sscaner[v] = vv + 1
		}
	}
	fmt.Printf("%v\n", sscaner)
	var keys []int
	for k := range sscaner {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var pbose = 0
	for i := 0; i < len(keys); i++ {

		skey := keys[i]
		var svalue, _ = sscaner[skey]
		for j := pbose; j < pbose+svalue; j++ {

			nums[j] = skey
		}

		pbose = pbose + svalue
	}
}

/*
前 K 个高频元素
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
*/

type cpair struct {
	K int
	F int
}

func TopKFrequent(nums []int, k int) []int {

	var ln = len(nums)
	if ln <= 1 {
		return nums
	}
	var sscaner = make(map[int]int)
	for _, v := range nums {
		if vv, mok := sscaner[v]; !mok {
			sscaner[v] = 1
		} else {
			sscaner[v] = vv + 1
		}
	}
	var topx []cpair = nil
	for k, v := range sscaner {
		topx = append(topx, cpair{
			K: k,
			F: v,
		})
	}

	fmt.Printf("Befor : %d,%v\n", len(topx), topx)
	sort.Slice(topx, func(i, j int) bool {
		return topx[i].F > topx[j].F
	})
	fmt.Printf("Afert : %d,%v\n", len(topx), topx)
	var res []int = nil
	for i := 0; i < k && i < len(topx); i++ {

		res = append(res, topx[i].K)
	}
	return res

}
func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	// fmt.Printf("Afert : %d,%v\n", len(nums), nums)
	var res []int = nil
	var ln = len(nums)

	// res = append(res, nums[ln-1])
	// for i := ln - 1; i >= 0; i-- {
	// 	if nums[i] != res[len(res)-1] {
	// 		res = append( res, nums[i])
	// 	}
	// }
	// fmt.Printf("res : %d,%v\n", len(res), res)

	//  for j := len(res) - 1; j >= 0; j-- {
	// 	if k == 1 {
	// 		return res[j]
	// 	}
	// 	k = k - 1
	// }

	res = append(res, nums[0])

	for i := 0; i < ln; i++ {

		if nums[i] != res[len(res)-1] {
			res = append(res, nums[i])
		}

	}

	return res[k-1]
}

func findPeakElement(nums []int) int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1

}

// nums 中的每个值都 独一无二
func SearchRotate(nums []int, target int) int {
	var rotateIndex = -1
	var nl = len(nums)
	for i := 1; i < nl; i++ {
		if nums[i-1] > nums[i] {
			rotateIndex = i - 1
			break
		}
	}

	return rotateIndex

}

/*
在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var pl = 0
	var pr = len(nums) - 1
	var mid = -1

	var tpos = -1
	for pl <= pr {
		var mid = (pl + pr) / 2
		if nums[mid] == target {
			//找着了
			fmt.Printf("Yes:%d\n", mid)
			tpos = mid
			break
		}
		if nums[mid] < target {
			pl = mid + 1
		} else {
			pr = mid - 1
		}
	}

	fmt.Printf("Yes2:%d\n", tpos)
	mid = tpos
	if mid == -1 {
		return []int{-1, -1}
	}
	var ps = mid
	var pe = mid
	for ; ps > 0 && nums[ps-1] == target; ps-- {

	}
	for ; pe < len(nums)-1 && nums[pe+1] == target; pe++ {

	}

	return []int{ps, pe}

}

func MMerge(intervals [][]int) [][]int {

	if len(intervals) == 1 {
		return intervals
	}
	fmt.Printf("Unsort :%v\n", intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Printf("Sorted :%v\n", intervals)

	var rows = len(intervals)
	var res [][]int = nil

	for j := 0; j < rows; {

		var pjStart = intervals[j][0]
		var pjEnd = intervals[j][1]

		fmt.Printf("%d,%d,%d\n", j, pjEnd, pjStart)
		var k = j + 1
		for ; k < rows; k++ {

			if pjEnd >= intervals[k][1] {
				//  Nothing to do
			} else {
				if pjEnd >= intervals[k][0] {
					pjEnd = intervals[k][1]
				} else {
					break
				}
			}

		}
		res = append(res, []int{pjStart, pjEnd})
		fmt.Printf("Crow :%d\n", k)
		j = k
	}

	return res

}

// 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
/*
提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 109

	grid = [][]int{
		{1,  4,  7,  11, 15},
		{2,  5,  8,  12, 19},
		{3,  6,  9,  16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
*/
func SsearchMatrix(matrix [][]int, target int) bool {

	var rows = len(matrix)
	var cols = len(matrix[0])
	if target < matrix[0][0] || target > matrix[rows-1][cols-1] {
		return false
	}
	binaryRowSearh := func(rowIndex, colStart, colEnd int) bool {
		if target > matrix[rowIndex][colEnd] || target < matrix[rowIndex][colStart] {
			return false
		}
		var l = colStart
		var r = colEnd
		for l >= 0 && r <= colEnd && l <= r {
			var mid = l + (r-l)/2
			if matrix[rowIndex][mid] == target {
				return true
			}
			if target < matrix[rowIndex][mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}
	binaryColSearh := func(colIndex, rowStart, rowEnd int) bool {
		if target < matrix[rowStart][colIndex] || target > matrix[rowEnd][colIndex] {
			return false
		}
		var l = rowStart
		var r = rowEnd
		for l >= 0 && r <= rowEnd && l <= r {
			var mid = l + (r-l)/2
			if matrix[mid][colIndex] == target {
				return true
			}
			if target < matrix[mid][colIndex] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}

	if rows == 1 {
		return binaryRowSearh(0, 0, cols-1)
	}
	if cols == 1 {
		return binaryColSearh(0, 0, rows-1)
	}

	var colStart = 0
	var colEnd = cols

	var rowStart = 0
	var rowEnd = rows

	//--- 沿着对角线查找

	for colStart < colEnd && rowStart < rowEnd {

		fmt.Printf("CP:%d,%d=%d\n", rowStart, colStart, matrix[rowStart][colStart])
		if target == matrix[rowStart][colStart] {
			return true
		}
		// fmt.Printf("rowData:%v\n", matrix[rowStart][0:colStart])
		if binaryRowSearh(rowStart, 0, colStart) {
			return true
		}
		var colData []int = nil
		for i := 0; i < rowStart; i++ {
			colData = append(colData, matrix[i][colStart])
		}
		fmt.Printf("colData:%v\n", colData)

		if binaryColSearh(colStart, 0, rowStart) {
			return true
		}

		rowStart = colStart + 1
		colStart = colStart + 1

	}
	return false

}

func SSesssearchMatrix(matrix [][]int, target int) bool {

	var rows = len(matrix)
	var cols = len(matrix[0])
	if target < matrix[0][0] || target > matrix[rows-1][cols-1] {
		return false
	}
	binaryRowSearh := func(rowIndex, colStart, colEnd int) bool {
		if target > matrix[rowIndex][colEnd] || target < matrix[rowIndex][colStart] {
			return false
		}
		var l = colStart
		var r = colEnd
		for l >= 0 && r <= colEnd && l <= r {
			var mid = l + (r-l)/2
			if matrix[rowIndex][mid] == target {
				return true
			}
			if target < matrix[rowIndex][mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}
	binaryColSearh := func(colIndex, rowStart, rowEnd int) bool {
		if target < matrix[rowStart][colIndex] || target > matrix[rowEnd][colIndex] {
			return false
		}
		var l = rowStart
		var r = rowEnd
		for l >= 0 && r <= rowEnd && l <= r {
			var mid = l + (r-l)/2
			if matrix[mid][colIndex] == target {
				return true
			}
			if target < matrix[mid][colIndex] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return false
	}

	if rows != cols {
		for ci := 0; ci < rows; ci++ {
			if binaryRowSearh(ci, 0, cols-1) {
				return true
			}
		}
		return false
	} else {
		fmt.Printf("对焦先查找")

		if matrix[0][0] == target {
			return true
		}
		if binaryRowSearh(0, 0, cols-1) {
			return true
		}
		if binaryColSearh(0, 0, rows-1) {
			return true
		}
		//--- 沿着对角线查找当M==N 是可以
		var lastRC = -1
		for cp := 1; cp < rows; cp++ {
			if matrix[cp][cp] == target {
				return true
			}
			if target < matrix[cp][cp] && target > matrix[cp-1][cp-1] {
				lastRC = cp
				break
			}
		}
		if lastRC == -1 {
			return false
		}

		if binaryRowSearh(lastRC-1, lastRC, cols-1) {
			return true
		}

		if binaryRowSearh(lastRC, 0, lastRC) {
			return true
		}

		if binaryColSearh(lastRC, 0, lastRC) {
			return true
		}

	}
	return false

}

func canJump(nums []int) bool {

	// https://zhuanlan.zhihu.com/p/355951310
	// 思路一：贪心算法（尽可能去最远的位置）
	// 如果能够到达最远的地方，那一定能到到最远前面的任意地方。
	// 1 遍历数组，把每个位置都作为起始点，设置最远能够到达位置为0
	// 2 更新最远能够到达的位置，能够到达的最远的位置为当前位置索引+当前元素的值
	// 3 如果每个位置都能够到达，则可以成功，否则失败

	var k = 0
	var i = 0
	for i = 0; i < len(nums); i++ {
		// 判断能否到达i位置，不能则也不能够到达终点，终止，直接返回结果false
		if i > k {
			return false
		}
		// 更新能够到达的最远位置
		if (i + nums[i]) > k {
			k = i + nums[i]
		}

		// 如果能够到达的最远位置超过终点，则提前跳出循环
		if k > len(nums) {
			return true
		}
	}
	return true

}

/*

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

*/

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	//dp[i][j] 表示  网格为 i 行 j 列 的布局下， 从[0,0]到[i-1][j-1]有多上种走法
	var dp [][]int = nil

	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				//相当与1行或是1lie 的i情况
				dp[i][j] = 1
			} else {

				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}

	return dp[m-1][n-1]

}

func CoinChange(coins []int, amount int) int {
	/*
		给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
		计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
		你可以认为每种硬币的数量是无限的。
	*/
	sort.Ints(coins)
	fmt.Printf("Coins:%v\n", coins)

	N := len(coins)
	// dp[i] = dp[i-Cj]+1, 其中Cj为数组coins中的一个硬币面额
	var dp = make([]int, amount+1)
	// 0元需要0个硬币
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1 //默认没有找到，因为  amount 个金额 需要 amount +1 个硬币的话，肯有至少一个硬币是小于0 的
		//求  dp[i]
		for j := 0; j < N; j++ {
			// 如果当前硬币面额coins[j]小于当前金额i
			if coins[j] <= i {
				//  cons[j]表示的是当前的面额

				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] == (amount + 1) {
		return -1
	}
	return dp[amount]

}

func CoinChangeExt(coins []int, amount int) int {
	// 采用递归算法

	/*


			https://mingshan.fun/2021/02/19/dp-coin-change/


		public static int solution3(int[] coins, int amount) {
		  // f(x) 代表x面额最少用多少枚硬币拼出
		  int[] f = new int[amount + 1];
		  // 初始化
		  f[0] = 0;

		  // 数值从小到大进行计算
		  // f[x] = min{ f[x - coin1] + 1 , f[x - coin2] + 1, ....}
		  // 1,2..27
		  for (int i = 1; i <= amount; i++) {
		    // 先设置f(i) 为最大值，好进行比较更新
		    f[i] = Integer.MAX_VALUE;

		    for (int j = 0; j < coins.length; j++) {
		      int coin = coins[j];
		      // 当前待拼的面额必须大于等于硬币的面额
		      // 并且f[x - coin1]不为Integer.MAX_VALUE，注意不能直接用 f[i] > (f[i - coin] + 1) 来判断，
		      // 因为 f[i - coin]为Integer.MAX_VALUE， + 1 就是负数了，此时的f[i]计算结果为负数
		      if (i >= coin && f[i - coin] != Integer.MAX_VALUE) {
		        f[i] = Math.min(f[i - coin] + 1, f[i]);
		      }
		    }
		  }

		  // 不能拼出，返回-1
		  if (f[amount] == Integer.MAX_VALUE) {
		    return -1;
		  }

		  return f[amount];
		}
	*/

	/*
		给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

		计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

		你可以认为每种硬币的数量是无限的。
	*/

	/*


			现有2元，5元，7元三种硬币，假设硬币都足够多，现求解：最少用多少枚上述硬币拼出27块钱？

		递归解法
		看到这个问题，我们一下子就可以想到一个树形结构，27块钱可以分别减去上面三种硬币的面额，剩下的值可以继续减去上面三种硬币的面额，直至无法再减，计算出刚好能减完的路线（从顶点到叶子节点）的节点数，最小的即是本题答案的解，图如下所示：

		image

		从上面的分析来看，很明显这道题可以用递归来做，而且比较简单，代码如下：

		public static int solution1(int x) {
			if (x == 0) {
			  return 0;
			}

			int res = x + 1;

			if (x >= 2) {
		    	res = Math.min(solution1(x - 2) + 1, res);
			}

			if (x >= 5) {
				res = Math.min(solution1(x - 5) + 1, res);
			}

			if (x >= 7) {
				res = Math.min(solution1(x - 7) + 1, res);
			}

			return res;
		}
	*/

	// 	这是一个挺经典的问题，解决方法有很多。这里介绍几个方法。
	// 首先要有思路，这类问题的解决方法都类似。
	// 总数为amount的现金换成n种不同方式的数目为：

	// A :将现金换成除第一种硬币之外的所有其他硬币方式的数目，加上
	// B: 将现金amount-coin换成所有硬币方式的数目，其中coin为第一种硬币的面值
	// 这样就包括了所有的方法。
	// 再定义最简单的情况：
	// amount==0，算作有一种方法，1
	// amount<0，没有方法，0
	// coin==0，算作没有方法，0
	// 递归

	var helper func(coins []int, start int, amout int) int

	helper = func(coins []int, start int, amout int) int {
		if amout == 0 {
			return 1
		}
		if amout < 0 || start == len(coins) {
			return 0
		}
		return helper(coins, start+1, amout) + helper(coins, start, amount-coins[start])

	}
	// private int helper(int[] coins, int start, int amount) {
	// 	if (amount == 0) return 1;
	// 	if (amount < 0 || start == coins.length) return 0;
	// 	return helper(coins, start + 1, amount) + helper(coins, start, amount - coins[start]);
	// }
	return helper(coins, 0, amount)

}

func LlengthOfLIS(nums []int) int {

	N := len(nums)

	if N == 1 {
		return 1
	}
	if N == 2 {
		if nums[0] < nums[1] {
			return 2
		} else {
			return 1
		}
	}

	// 算法思路

	// 因为所求为子序列，很容易想到一种线性动态规划。
	// 对于求最长上升子序列，上升我们肯定需要知道当前阶段最后一个元素为多少，最长我们还要知道当前我们的序列有多长。
	// 那么我们就可以这样定义状态：
	//    设 dp[i] 表示以 nums[i] 为结尾的最长上升子序列的长度
	//   ，为了保证元素单调递增肯定只能从 i 前面且末尾元素比 nums[i] 小的状态转移过来
	// 代码思路

	// 状态转移方程为

	// 每个位置初始值为 dp[i]=1（将自己作为一个序列）
	// 答案可以是任何阶段中只要长度最长的那一个，所以我们边转移边统计答案

	dp := make([]int, N)

	var ans = 0
	for i := 0; i < N; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {

			if nums[j] < nums[i] {
				//
				if (dp[j] + 1) > dp[i] {
					dp[i] = dp[j] + 1
				}
			}

		}
		//---取最长的
		if dp[i] > ans {
			ans = dp[i]
		}

	}
	return ans

}

// First Greater or Equal
func firstGE(nums []int, target int) int {

	if nums[0] >= target {
		return 0
	}
	var start = 0
	var end = len(nums)
	for start+1 < end {
		var mid = start + (end-start)/2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}
	return end
}

func LlengthOfLISExt(nums []int) int {

	//10, 9, 2, 5, 3, 7, 101, 18

	N := len(nums)
	if N <= 1 {
		return 1
	}

	var lis = make([]int, N)
	lis[0] = 1 << 31
	for i := 1; i < N; i++ {
		lis[i] = 1 << 31
	}

	var ans = 0

	for i := 0; i < N; i++ {
		index := firstGE(lis, nums[i])

		lis[index] = nums[i]
		if index > ans {
			ans = index
		}
	}
	fmt.Printf("%v\n", lis[:ans+1])
	return ans + 1 //  ans 的基数为0

}
