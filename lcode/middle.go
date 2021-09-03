package leetcode

import (
	"fmt"
	"strings"
)

//简单、暴力遍历所有元素
func firstSetZeroes(matrix [][]int) {

	lx := len(matrix)
	ly := len(matrix[0])

	var bl [][]bool = nil

	for i := 0; i < lx; i++ {

		var br = make([]bool, ly)

		for j := 0; j < ly; j++ {
			br[j] = matrix[i][j] == 0
		}
		bl = append(bl, br)
	}

	for i := 0; i < lx; i++ {

		for j := 0; j < ly; j++ {
			if bl[i][j] {

				for ci := 0; ci < lx; ci++ {
					matrix[ci][j] = 0
				}

				for cj := 0; cj < ly; cj++ {
					matrix[i][cj] = 0
				}

			}
		}

	}
}

//--- M+N 个空间
func secondSetZeroes(matrix [][]int) {

	lx := len(matrix)
	ly := len(matrix[0])

	// 行记录
	var rows []bool = make([]bool, lx)
	// 列
	var cols []bool = make([]bool, ly)

	for i := 0; i < lx; i++ {

		for j := 0; j < ly; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	fmt.Printf("%v", rows)
	fmt.Printf("%v", cols)

	for i := 0; i < lx; i++ {

		if rows[i] {
			for j := 0; j < ly; j++ {
				matrix[i][j] = 0
			}
		}

	}
	for j := 0; j < ly; j++ {

		if cols[j] {
			for i := 0; i < lx; i++ {
				matrix[i][j] = 0
			}
		}
	}

}

func SetZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	firstRowFillZero := false
	firstColFillZero := false

	for c0 := 0; c0 < cols; c0++ {
		if matrix[0][c0] == 0 {
			firstRowFillZero = true
			break
		}
	}

	for r0 := 0; r0 < rows; r0++ {

		if matrix[r0][0] == 0 {
			firstColFillZero = true
			break
		}
	}

	fmt.Printf("FirstRow :%v, FirstCol:%v\n", firstRowFillZero, firstColFillZero)

	for i := 1; i < rows; i++ {

		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
				fmt.Printf("Find\n")
			}
		}
	}

	for jx := 1; jx < cols; jx++ {

		if matrix[0][jx] == 0 {

			for jr := 1; jr < rows; jr++ {

				matrix[jr][jx] = 0
			}
		}

	}

	for jy := 1; jy < rows; jy++ {

		if matrix[jy][0] == 0 {

			for jc := 1; jc < cols; jc++ {

				matrix[jy][jc] = 0
			}
		}

	}

	fmt.Printf("%v\n", matrix)
	if firstColFillZero {

		for j := 0; j < rows; j++ {
			matrix[j][0] = 0
		}
	}
	if firstRowFillZero {

		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
}
func lengthOfLongestSubstring(s string) int {
	//最长字符串
	var lx = len(s)
	if lx == 0 || lx == 1 {
		return lx
	}
	if lx == 2 && s[0] != s[1] {
		return lx
	}
	if lx == 3 && s[0] != s[1] && s[0] != s[2] && s[1] != s[2] {
		return lx
	}

	var maxSize = 1

	var startPos = 0
	var endPos = 0

	for i := 1; i < lx; i++ {

		fmt.Printf("CStr:%s\n", s[startPos:endPos+1])

		var cIndex = strings.Index(s[startPos:endPos+1], string(s[i:i+1]))
		if -1 != cIndex {
			//
			//startPos =  cIndex + 1
			//  cIndex  只是  s[i]  在 子串 s[startPos:endPos +1]中的位置
			//
			startPos = startPos + cIndex + 1
		}
		endPos = i
		if (endPos - startPos + 1) > maxSize {
			maxSize = endPos + 1 - startPos
		}
		fmt.Printf("CStr2:%s\n", s[startPos:endPos+1])
	}
	return maxSize

}
func printListNode(node *ListNode) {

	fmt.Print("\n[")
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println("]")
}

func LongestPalindrome(s string) string {
	// 给你一个字符串 s，找到 s 中最长的回文子串。
	lx := len(s)
	if lx <= 1 {
		return s
	}
	if lx == 2 && s[0] == s[1] {
		return s[0:1]
	}

	fmt.Printf("Args:%s\n", s)

	var maxSize = 0
	var str string = ""

	//如果字符串长度超过4 则第一个字符和最后一个字符肯定不是最长回文子串
	for i := 1; i < lx-1; i++ {

		var pl = i

		var pr = i

		for pr < lx-1 && s[i] == s[pr+1] {
			pr = pr + 1
		}

		for pl > 0 && s[i] == s[pl-1] {
			pl = pl - 1
		}

		for pl > 0 && pr < lx-1 {

			//fmt.Printf("PL:%d,PR=%d\n", pl, pr)
			if s[pl-1] == s[pr+1] {
				pl, pr = pl-1, pr+1
			} else {
				break
			}

		}

		var cl = pr - pl + 1

		//fmt.Printf("CStrnig：i:%d  %d,%d-->%s\n", i, pl, pr, s[pl:pr+1])
		if cl > maxSize {
			maxSize = cl
			str = s[pl : pr+1]

		}

	}

	return str

}

func LongestPalindromeDynamic(s string) string {
	// 给你一个字符串 s，找到 s 中最长的回文子串。
	lx := len(s)
	if lx <= 1 {
		return s
	}
	if lx == 2 && s[0] == s[1] {
		return s
	}

	return ""

}

func increasingTriplet(nums []int) bool {

	nl := len(nums)

	if nl < 3 {
		return false
	}
	if nl == 3 {
		return nums[0] < nums[1] && nums[1] < nums[2]
	}

	// 使用两个int变量m1与m2，初始化为整型最大值。
	// 开始遍历数组,此时分三种情况
	// m1大于等于数组当前值nums[i]，此时，将数组当前值nums[i]赋值给m1

	// m1小于数组当前值nums[i]、m2大于数组当前值nums[i]，此时，将数组当前值赋值给m2。这里有个隐含的状态是，m2更新代表着我们已经找到了一个递增的二元子序列。
	// 接下来的查找中只需要找到一个值大于m2就说明存在递增的三元子序列，直接返回true即可。

	// m2也小于数组当前值nums[i]，因为前两种情况进入时已经找到了递增的二元子序列，此时直接返回true即可。
	// 当在遍历中没有返回时，代表数组中最多只有递增的二元子序列，直接返回false即可
	// 注意在循环中，实际上每次判断都会尽量减小m1与m2的值，毕竟m2的值越小，在接下来的数据中大于m2的可能性就更高

	// 作者：王可尊
	// 链接：https://www.jianshu.com/p/ed5e72bc4c89
	// 来源：简书
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	var m1 = 1<<31 - 1
	var m2 = 1<<31 - 1

	for i := 0; i < nl; i++ {
		if nums[i] <= m1 {
			m1 = nums[i]
		} else {
			if nums[i] <= m2 {
				m2 = nums[i]
			} else {
				return true
			}
		}

	}
	return false

}

func OddEvenSwitchList(head *ListNode) *ListNode {

	// 这个是交错数组

	// 示例 1:

	// 输入: 1->2->3->4->5->NULL
	// 输出: 1->3->5->2->4->NULL
	// 示例 2:

	// 输入: 2->1->3->5->6->4->7->NULL
	// 输出: 2->3->6->7->1->5->4->NULL
	// 说明:

	// 应当保持奇数节点和偶数节点的相对顺序。
	// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{
		Val:  -1,
		Next: nil,
	}

	var pc = dummyNode

	var pfirst = head

	var pSecond = head.Next

	for pfirst != nil && pSecond != nil {

		head = pSecond.Next

		pfirst.Next = nil
		pSecond.Next = nil

		pc.Next = pSecond
		pc.Next.Next = pfirst
		pc = pc.Next.Next

		fmt.Printf("%d,%d   ", pfirst.Val, pSecond.Val)
		pfirst = head
		pSecond = head.Next

	}

	if pfirst != nil {
		pc.Next = pfirst
		fmt.Printf("  %d   ", pfirst.Val)
	}

	return dummyNode.Next
}
func OddEvenList(head *ListNode) *ListNode {
	//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

	// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

	// 示例 1:

	// 输入: 1->2->3->4->5->NULL
	// 输出: 1->3->5->2->4->NULL
	// 示例 2:

	// 输入: 2->1->3->5->6->4->7->NULL
	// 输出: 2->3->6->7->1->5->4->NULL
	// 说明:

	// 应当保持奇数节点和偶数节点的相对顺序。
	// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

	if head == nil || head.Next == nil {
		return head
	}

	var dummyNode = &ListNode{
		Val:  -1,
		Next: nil,
	}

	var pc = dummyNode

	var p1 = head

	for p1 != nil {
		var p2 = p1.Next

		if p2 == nil {
			break
		}

		p1.Next = p2.Next

		p2.Next = nil

		pc.Next = p2

		pc = pc.Next

		if p1.Next != nil {
			p1 = p1.Next
		}

	}
	printListNode(dummyNode)
	p1.Next = dummyNode.Next

	return head
}
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	//     相交链表
	// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

	// 图示两个链表在节点 c1 开始相交：

	// ListA = [4,0,8,4,5], listB = [3,0,1,8,4,5],  焦急为  [8,4,5]

	//说明:
	/*
	   LA + LB  =  [4,0,8,4,5,3,0,1,8,4,5]
	   LB + LA  =  [3,0,1,8,4,5,4,1,8,4,5]
	*/

	var a = headA
	var b = headB
	for a != b {

		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}

	}
	return a

}

func InorderTraversal(root *TreeNode) []int {
	//给定一个二叉树的根节点 root ，返回它的 中序 遍历。

	/*
	                    1
	   		2 --------------3
	   	4                             5
	   	    6
	       7 8

	   	如图1所示，三种遍历方法(人工)得到的结果分别是：

	   先序：1 2 4 6 7 8 3 5
	   中序：4 7 6 8 2 1 3 5
	   后序：7 8 6 4 2 5 3 1

	   三种遍历方法的考查顺序一致，得到的结果却不一样，原因在于：

	   先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)

	   中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)

	   后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

	*/

	/*
			思路：
		1、中序遍历 ： 左、根、右


		    public List<Integer> inorderTraversal(TreeNode root) {
		        List<Integer> list = new ArrayList<>();
		        midFor(root, list);
		        return list;
		    }

		    private void midFor(TreeNode root, List<Integer> list) {
		        if (root==null){
		            return;
		        }
		        midFor(root.left, list);

		        list.add(root.val);

		        midFor(root.right, list);
		    }

	*/
	if root == nil {
		return nil
	}

	var res []int = nil

	var pr = root

	var queuq []*TreeNode = nil

	for pr != nil || len(queuq) != 0 {

		for pr != nil {

			queuq = append(queuq, pr)
			pr = pr.Left
		}

		if len(queuq) > 0 {
			var qlen = len(queuq)
			var lq = queuq[qlen-1]
			res = append(res, lq.Val)
			pr = lq.Right
			queuq = queuq[0 : qlen-1]
		}

	}

	return res

}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	// 写解决层次遍历
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {

		return [][]int{
			{root.Val},
		}
	}

	var result [][]int = nil

	var pr = root

	var stack []*TreeNode = nil

	stack = append(stack, pr)
	stack = append(stack, nil)
	result = append(result, []int{root.Val})
	var exchangeDirection = true

	for len(stack) > 0 {

		cNode := stack[0]
		stack = stack[1:]
		if cNode == nil {
			if len(stack) > 0 {

				var tr []int

				var ls = len(stack)
				if exchangeDirection {
					for i := ls - 1; i >= 0; i-- {
						tr = append(tr, stack[i].Val)
					}
				} else {
					for i := 0; i < ls; i++ {
						tr = append(tr, stack[i].Val)
					}
				}

				result = append(result, tr)
				stack = append(stack, nil)
				exchangeDirection = !exchangeDirection
			}

		} else {

			if cNode.Left != nil {
				stack = append(stack, cNode.Left)
			}

			if cNode.Right != nil {
				stack = append(stack, cNode.Right)
			}

		}

	}

	return result

}
