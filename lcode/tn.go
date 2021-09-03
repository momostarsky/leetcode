package leetcode

import (
	"fmt"
	"math"
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		v1 := 1 + MaxDepth(root.Left)
		v2 := 1 + MaxDepth(root.Right)

		if v1 >= v2 {
			return v1
		} else {
			return v2
		}
	}

}

func isValidBSTSub(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}
	//每个节点如果超过这个范围，直接返回false
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	//  左右子树也必须满足二叉搜索树
	// 左子树的最大值是   root.Val
	//  右子树的最小值 是  root.Val
	//  Left.Val < root.Val < Right.Val

	return isValidBSTSub(root.Left, minVal, root.Val) && isValidBSTSub(root.Right, root.Val, maxVal)
}

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return true
	} else {
		return isValidBSTSub(root, math.MinInt64, math.MaxInt64)
	}
}

func isSymmetricExt(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		// 没有字节点的
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		//---只有一个分支,瘸腿
		return false
	}
	//  [1,2,2,3,4,4,3]
	//   镜像的判断标准以第三层3,4,4,3 为例:
	//   root.Left.Val == root.Right.Val
	//   root.Left.Left == root.Right.Right
	//   root.Left.Right == root.Right.Left

	return isSymmetricExt(left.Left, right.Right) && isSymmetricExt(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricExt(root.Left, root.Right)
}

func SortedArrayToBST(nums []int) *TreeNode {

	lx := len(nums)
	if lx == 0 {
		return nil
	}

	if lx == 1 {
		return &TreeNode{Val: nums[0]}
	}

	if lx == 2 {
		var l = &TreeNode{Val: nums[0]}
		return &TreeNode{Val: nums[1], Left: l}
	}

	if lx == 3 {
		var l = &TreeNode{Val: nums[0]}
		var r = &TreeNode{Val: nums[2]}
		return &TreeNode{Val: nums[1], Left: l, Right: r}

	}

	var midPos = lx / 2
	var root = &TreeNode{Val: nums[midPos]}

	var rNums = nums[:midPos]

	fmt.Printf("%v\n", rNums)
	var lNums = nums[midPos+1:]
	fmt.Printf("%v\n", lNums)
	root.Left = SortedArrayToBST(rNums)
	root.Right = SortedArrayToBST(lNums)

	return root

}

func Merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		nums1 = nums1[0:m]
	} else if m == 0 {

		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}

	} else {

		// PS:  nums1 的长度为 m+n +1
		//  nums1 从小往大,
		var p1 = m - 1
		var p2 = n - 1
		var pRes = m + n - 1
		for p2 >= 0 {

			if p1 >= 0 && nums1[p1] > nums2[p2] {
				nums1[pRes] = nums1[p1]
				p1 = p1 - 1
			} else {
				nums1[pRes] = nums2[p2]
				p2 = p2 - 1
			}
			pRes = pRes - 1
		}

	}

}

func Fibonacci2_(n int) (res int64) {
	if n <= 1 {
		res = int64(n)
		return
	}

	a1, a2 := int64(0), int64(1)

	for i := 2; i <= n; i++ {
		a1, a2 = a2, a1+a2
	}
	res = a2
	return
}

// def fibonacciTail(main: Int, current: Int, next: Int): Long = {
// 	main match {
// 	  case 1 | 2 => next
// 	  case _ => fibonacciByTail(main - 1, next, current+next)
// 	}
// 	fibonacciTail(n, 1, 1)
//   }

func fibTail(main int, currrent int, next int) int {

	if main <= 1 {
		return next
	} else {
		return fibTail(main-1, next, currrent+next)
	}

}

func climbStairs(n int) int {

	//===这种递归会超时
	//return  climbStairs(n-1) + climbStairs(n-2)
	//-- 采用为尾递归或是所谓的克里话
	return fibTail(n, 1, 1)

}

func MaxProfit(prices []int) int {

	// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

	// 作者：力扣 (LeetCode)
	// 链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn8fsh/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	var lx = len(prices)
	if lx <= 1 {
		//--- 不产生交易的情况下,最大利润为0
		return 0
	}
	//---记录第i天交易产生的利润
	var profits []int = make([]int, lx)
	profits[0] = 0
	// 记录当天以前股价最低的一天的价格
	var minPrice = prices[0]

	for i := 1; i < lx; i++ {
		//-- 第i天的最大利润
		var profitDay = prices[i] - minPrice
		if profits[i-1] > profitDay {
			profits[i] = profits[i-1]
		} else {
			profits[i] = profitDay
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profits[lx-1]
	// 则转移方程为：
	// dp[i] = max( dp[i-1], prices[i] - min(prices[0 -> i]))。

	// 也就是说，每日的最大利润为以下两项的大者：

	// 昨天的做大利润
	// 在当天以前的某个股价最低的一天买入股票，并在当天卖出股票，以此赚取的利润

}

func MaxSubArray(nums []int) int {
	//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	// 	1 <= nums.length <= 3 * 104
	// -10^5 <= nums[i] <= 10^5
	var lx = len(nums)
	if lx <= 1 {
		return nums[0]
	}

	// var ms = make([]int, lx)
	// ms[0] = nums[0]
	// var max = ms[0]
	// for i := 1; i < lx; i++ {

	// 	if ms[i-1] < 0 {
	// 		// 前一天为复数,越加越小,干脆从今天算起
	// 		ms[i] = nums[i]
	// 	} else {
	// 		// 前一天为正数,从X-1 天算起
	// 		ms[i] = ms[i-1] + nums[i]
	// 	}
	// 	if ms[i] > max {
	// 		max = ms[i]
	// 	}

	// }
	// fmt.Printf("%v\n", ms)

	var ms = [2]int{0, 0}
	ms[0] = nums[0]
	ms[1] = nums[0]
	var max = ms[0]
	for i := 1; i < lx; i++ {
		if ms[0] < 0 {
			// 前一天为复数,越加越小,干脆从今天算起
			ms[0] = nums[i]
			//	ms[1] = ms[0]
		} else {
			// 前一天为正数,从X-1 天算起
			ms[0] = ms[0] + nums[i]
			//	ms[1] = ms[0]
		}
		if ms[0] > max {
			max = ms[0]
		}

	}
	fmt.Printf("%v\n", ms)

	return max

	/*
		说明:
		    nums[0..N] 的最大连续子数组

			// dp[i]表示某个以下标i元素为右端点（包含）的连续子数组的和，这个连续子数组的和最大。
			ms[i] =   表示    nums[0]+ ...+ nums[i]

		    则 ms[i+1] 是最大的连续数组
			if nums[i+1] > 0 {
				//-- 继续累加
				ms[i+1]= ms[i]+ nums[i+1]
			}  else {
				//  小于0 ,越加越小,不要加了
				 ms[i+1] =nums[i+1]
			}
	*/

}

func Rob(nums []int) int {

	// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

	lx := len(nums)
	if lx == 1 {
		return nums[0]
	} else if lx == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	} else {
		/*
			说明:
			   假设 X 个房间能抢劫到的最大值为:rob[X], 那么对第N 个房间有且仅有两种选择
			   1)   抢   :   rob[N] = rob[N-2] + nums[N]
			   2)   不抢 :   rob[N] = rob[N-1]

			rob[0] = nums[0]

			rob[1] = Max( nums[0], nums[1])

		*/
		// var rob = make([]int, lx)
		// //----第一个房间
		// rob[0] = nums[0]
		// if nums[1] > nums[0] {
		// 	rob[1] = nums[1]
		// } else {
		// 	rob[1] = nums[0]
		// }
		// for N := 2; N < lx; N++ {
		// 	var mj = rob[N-2] + nums[N]
		// 	if rob[N-1] > mj {
		// 		rob[N] = rob[N-1]
		// 	} else {
		// 		rob[N] = mj
		// 	}
		// }
		// return rob[lx-1]

		// PS:  每次运算的时候只用到最后2个

		var rob = [3]int{0, 0, 0}
		//----第一个房间
		rob[0] = nums[0]
		if nums[1] > nums[0] {
			rob[1] = nums[1]
		} else {
			rob[1] = nums[0]
		}
		for N := 2; N < lx; N++ {
			var mj = rob[0] + nums[N]
			if rob[1] > mj {
				rob[2] = rob[1]
			} else {
				rob[2] = mj
			}

			rob[0], rob[1] = rob[1], rob[2]
		}

		return rob[2]

	}

}

func toInt(l1 *ListNode) int {

	var aStep = 10
	var arr1 = l1.Val
	for l1.Next != nil {
		arr1 = arr1 + aStep*l1.Next.Val
		l1 = l1.Next
		aStep = aStep * 10

	}
	return arr1
}

func xprintListNode(node *ListNode) {

	fmt.Print("\n[")
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println("]")
}

func AddTwoNumbersX2(l1 *ListNode, l2 *ListNode) *ListNode {

	var res *ListNode = l1
	for l2 != nil && l1 != nil {

		l1.Val = l1.Val + l2.Val

		if l1.Next == nil && l2.Next == nil {
			break
		}
		if l1.Next != nil && l2.Next != nil {
			l1 = l1.Next
			l2 = l2.Next
		} else {

			if l1.Next == nil {
				l1.Next = l2.Next
				break
			}

		}

	}

	xprintListNode(res)

	var Carray = 0

	pf := res
	for pf != nil {
		Carray = pf.Val / 10
		pf.Val = pf.Val % 10

		if Carray > 0 && pf.Next == nil {
			pf.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			break
		} else {
			if Carray > 0 {
				pf.Next.Val = pf.Next.Val + 1
			}
			pf = pf.Next
		}

	}

	return res

}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 	每个链表中的节点数在范围 [1, 100] 内
	// 0 <= Node.val <= 9
	// 题目数据保证列表表示的数字不含前导零

	if l1.Next == nil && l2.Next == nil {
		var cb = l1.Val + l2.Val
		if cb >= 10 {
			l1.Val = cb % 10
			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		} else {
			l1.Val = cb
		}
		return l1

	}

	var arrx = [101]byte{0}

	var n1 = 0
	for i := 0; l1 != nil; n1, i = n1+1, i+1 {
		arrx[i] = byte(l1.Val)
		l1 = l1.Next
	}
	var lp = 0
	var px byte = 0
	var n2 = 0
	for i := 0; i < 101; i++ {

		var l2v = 0
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
			n2++
			if n2 > n1 {
				n1 = n2
			}
		}
		var cb = arrx[i] + byte(l2v) + px
		arrx[i] = cb % 10
		px = cb / 10

		if l2 == nil && i > n1 {
			lp = i
			break
		}

	}
	//fmt.Printf("%v\n", arrx[0:lp])
	var res *ListNode = &ListNode{

		Val:  int(arrx[0]),
		Next: nil,
	}
	var p *ListNode = res

	for i := 1; i < lp; i++ {

		if i == lp-1 && arrx[i] == 0 {
			//最后一个ie为0
			break
		}
		p.Next = &ListNode{
			Val:  int(arrx[i]),
			Next: nil,
		}
		p = p.Next
	}

	return res

}
