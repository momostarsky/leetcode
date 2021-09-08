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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 给定一棵树的前序遍历 preorder 与中序遍历inorder。请构造二叉树并返回其根节点。
	/*
		                1
		   		2 --------------3
		   	4                        5
		   	    6
		       7  8
			如图1所示，三种遍历方法(人工)得到的结果分别是：
			先序：1 2 4 6 7 8 3 5
			中序：4 7 6 8 2 1 3 5

			后序：7 8 6 4 2 5 3 1
			三种遍历方法的考查顺序一致，得到的结果却不一样，原因在于：

			先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)

			中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)

			后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

			作者：阿菜的博客
			链接：https://www.jianshu.com/p/456af5480cee
			来源：简书
			著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	*/

	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	if len(preorder) == 2 {

		var rootNode = &TreeNode{
			Val: preorder[0],
		}
		if preorder[0] == inorder[0] {
			rootNode.Right = &TreeNode{
				Val: preorder[1],
			}
		} else {
			rootNode.Left = &TreeNode{
				Val: preorder[1],
			}
		}
		return rootNode

	}
	/*
		1) preorder 的第一个是跟节点
	*/
	var rootVal = preorder[0]

	var rootIndex = -1

	for i := 0; i < len(inorder); i++ {

		if rootVal == inorder[i] {
			rootIndex = i
			break
		}
	}

	//	var res []*TreeNode = nil

	/// 跟节点的值

	inSubLeft := inorder[:rootIndex]
	inSubRight := inorder[rootIndex+1:]

	preSubLeft := preorder[1 : len(inSubLeft)+1]
	preSubRight := preorder[len(inSubLeft)+1:]

	// fmt.Printf("leftNodes:%v\n", inSubLeft)
	// fmt.Printf("rightNodes:%v\n", inSubRight)

	// fmt.Printf("leftNodes:%v\n", preSubLeft)
	// fmt.Printf("rightNodes:%v\n", preSubRight)

	var leftTree = buildTree(preSubLeft, inSubLeft)

	var rightTree = buildTree(preSubRight, inSubRight)

	return &TreeNode{
		Val:   rootVal,
		Left:  leftTree,
		Right: rightTree,
	}
}

//buildTreeWithoutRecurive  非递归构建二插树
func buildTreeWithoutRecurive(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var pIn = 0

	var nlen = len(inorder)

	var cNode *TreeNode = nil
	//			     3
	//			   	/ \
	//			   9  20
	//			  /  /  \
	//			 8  15   7
	//			/ \
	//		   5  10
	//		  /
	//		 4

	//  preorder = [3, 9, 8, 5, 4, 10, 20, 15, 7]
	//  inorder = [4, 5, 8, 10, 9, 3, 15, 20, 7]

	for i := 1; i < nlen; i++ {
		// 当期节点是最后一个节点
		var preorderVal = preorder[i]
		cNode = stack[len(stack)-1]
		// 读取当前节点的全部左子树, 此时 inorder[pIn] 就是最左子树
		if cNode.Val != inorder[pIn] {
			cNode.Left = &TreeNode{
				Val: preorder[i],
			}
			stack = append(stack, cNode.Left)

		} else {
			//  当前堆栈顶部的左节点
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[pIn] {
				cNode = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				pIn++
			}

			cNode.Right = &TreeNode{
				Val: preorderVal,
			}
			stack = append(stack, cNode.Right)
		}

	}

	return root

}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	/*
			方法二：迭代
		思路

		迭代法是一种非常巧妙的实现方法。

		对于前序遍历中的任意两个连续节点 uu 和 vv，根据前序遍历的流程，我们可以知道 uu 和 vv 只有两种可能的关系：

		vv 是 uu 的左儿子。这是因为在遍历到 uu 之后，下一个遍历的节点就是 uu 的左儿子，即 vv；

		uu 没有左儿子，并且 vv 是 uu 的某个祖先节点（或者 uu 本身）的右儿子。
		如果 uu 没有左儿子，那么下一个遍历的节点就是 uu 的右儿子。
		如果 uu 没有右儿子，我们就会向上回溯，直到遇到第一个有右儿子（且 uu 不在它的右儿子的子树中）的节点 u_au
		a
		​
		 ，那么 vv 就是 u_au
		​
		  的右儿子。
	*/

	return buildTreeWithoutRecurive(preorder, inorder)
}

///
/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


示例：



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvijdh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectCore(l *Node, r *Node) {

	l.Next = r

	// 因为是满二叉树，所以进入递归前只需判断一个分支是否为空
	if l.Left == nil {
		return
	}

	// 递归，左右节点的所有子树从左往右依次连接
	// 这样到最后最右子树的next必定指向null，而其它子树也已经依次连接
	connectCore(l.Left, l.Right)
	connectCore(l.Right, r.Left)
	connectCore(r.Left, r.Right)

}

func connect(root *Node) *Node {
	///给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

	if root == nil || root.Left == nil {

		return root
	}

	connectCore(root.Left, root.Right)
	return root

}

func enumNode(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var left []int = nil
	if root.Left != nil {
		left = append(left, enumNode(root.Left)...)
	}
	left = append(left, root.Val)

	if root.Right != nil {
		left = append(left, enumNode(root.Right)...)
	}
	return left

}

func KthSmallest(root *TreeNode, k int) int {
	var res []int = nil
	var lq []*TreeNode = nil

	for root != nil || len(lq) > 0 {

		// 左树进入堆栈
		for root != nil {
			lq = append(lq, root)

			root = root.Left
		}
		//---
		cr := lq[len(lq)-1]

		res = append(res, cr.Val)
		root = cr.Right
		lq = lq[:len(lq)-1]

	}

	fmt.Printf("中序列遍历结果:%v\n", res)
	return res[k-1]

}

//---深度优先便利会内存溢出
func dfs(grid [][]byte, x, y int, m, n int) {

	if x >= 0 && x < m && y >= 0 && y < n { //防止越界
		if grid[x][y] == '1' {
			grid[x][y] = '0'
			dfs(grid, x+1, y, m, n)
			dfs(grid, x-1, y, m, n)
			dfs(grid, x, y+1, m, n)
			dfs(grid, x, y-1, m, n)
		}
	}
}

// ----广度优先便利
func bfs(grid [][]byte, x, y int, m, n int) {

	var lq [][2]int = nil

	lq = append(lq, [2]int{x, y})
	for len(lq) > 0 {
		var arr = lq[0]
		var i = arr[0]
		var j = arr[1]
		lq = lq[1:]
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == '1' {

			grid[i][j] = '0'
			lq = append(lq, [2]int{i + 1, j})
			lq = append(lq, [2]int{i - 1, j})
			lq = append(lq, [2]int{i, j + 1})
			lq = append(lq, [2]int{i, j - 1})

		}
	}

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var m = len(grid)    // rows
	var n = len(grid[0]) // cols
	var res = 0

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == '1' {
				bfs(grid, x, y, m, n) //找到一个起点全都玷污一遍
				res = res + 1         //记录小岛个数
			}
		}
	}
	return res
}

func JpNumIslands(grid [][]byte) int {
	return numIslands(grid)
}

var dmap map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var res []string = nil

	if len(digits) == 1 {
		var ax, _ = dmap[digits[0]]
		for i := 0; i < len(ax); i++ {
			res = append(res, string(ax[i]))
		}
		return res
	}

	if len(digits) == 2 {
		var ax, _ = dmap[digits[0]]
		var arr []string = nil
		for i := 0; i < len(ax); i++ {
			arr = append(arr, string(ax[i]))
		}

		var bx, _ = dmap[digits[1]]

		var barr []string = nil
		for j := 0; j < len(bx); j++ {
			barr = append(barr, string(bx[j]))
		}

		for _, v := range arr {

			for _, b := range barr {

				var x = v + b

				res = append(res, x)

			}

		}
		return res
	}

	var ax, _ = dmap[digits[0]]
	var arr []string = nil
	for i := 0; i < len(ax); i++ {
		arr = append(arr, string(ax[i]))
	}
	var brr = LetterCombinations(digits[1:])
	for _, v := range arr {

		for _, b := range brr {

			var x = v + b

			res = append(res, x)

		}
	}
	return res

}

func GenerateParenthesis(n int) []string {
	//1 <= n <= 8

	var dp [][]string = nil
	dp = append(dp, []string{
		"",
	})

	for i := 1; i <= n; i++ {
		var cur []string = nil
		for j := 0; j < i; j++ {

			var str1 = dp[j]
			var str2 = dp[i-1-j]
			for _, s1 := range str1 {
				for _, s2 := range str2 {
					// 枚举右括号的位置

					cur = append(cur, "("+s1+")"+s2)

				}
			}
		}
		dp = append(dp, cur)

	}
	return dp[n]

}

func valid(current string) bool {
	var balance = 0
	for _, c := range current {
		if c == '(' {
			balance = balance + 1
		} else {
			balance = balance - 1
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func isValidate(pm string) bool {

	var lq []byte = nil
	var nl = len(pm)
	var pos = 0
	for pos < nl {

		if pm[pos] == '(' {
			lq = append(lq, '(')
		} else {
			if len(lq) == 0 {
				return false
			}
			lq = lq[:len(lq)-1]
		}
		pos = pos + 1
	}

	return len(lq) == 0

}

func GenerateParenthesisWithDP(n int) []string {
	//1 <= n <= 8

	if n == 1 {
		return []string{
			"()",
		}
	}
	if n == 2 {
		return []string{
			"(())",
			"()()",
		}
	}
	//  为了生成所有序列，我们可以使用递归。长度为 n 的序列就是在长度为 n-1 的序列前加一个 '(' 或 ')'。
	var res []string = nil
	//  把所有可能的（）对都生成出来

	var subx = GenerateParenthesisWithDP(n - 1)

	for _, v := range subx {
		for i := 0; i < len(v); i++ {

			v1 := string(v[:i]) + "()" + string(v[i:])
			if valid(v1) {
				res = append(res, v1)
			}
			v2 := string(v[:i]) + ")(" + string(v[i:])
			if valid(v2) {
				res = append(res, v2)
			}
		}

	}

	var ojk []string = nil
	var mp map[string]bool = make(map[string]bool)
	for _, v := range res {
		if _, mok := mp[v]; !mok {
			mp[v] = true
			ojk = append(ojk, v)
		}

	}

	return ojk

}

var dfsList []string = nil

func backtrack(S []byte, left, right, n int) {
	if len(S) == 2*n {
		dfsList = append(dfsList, string(S))
		return
	}
	if left < n {
		S = append(S, '(')
		backtrack(S, left+1, right, n)
		S = S[:len(S)-1]
	}
	if right < left {
		S = append(S, ')')
		backtrack(S, left, right+1, n)

	}
}

func GenerateParenthesisWithDFS(n int) []string {
	dfsList = nil
	var prx []byte
	backtrack(prx, 0, 0, n)
	return dfsList
}
