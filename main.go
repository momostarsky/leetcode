package main

import (
	"fmt"
	leetcode "leetcode/lcode"
)

func createListNode(nums []int) *leetcode.ListNode {

	var res = &leetcode.ListNode{
		Val:  nums[0],
		Next: nil,
	}
	var cp = res

	for i := 1; i < len(nums); i++ {
		cp.Next = &leetcode.ListNode{
			Val:  nums[i],
			Next: nil,
		}
		cp = cp.Next
	}

	return res
}

func createListNodeWith2(header *leetcode.ListNode, tail *leetcode.ListNode) *leetcode.ListNode {

	var ph = header
	for ; ph != nil && ph.Next != nil; ph = ph.Next {

	}
	ph.Next = tail

	return ph

}

func printListNode(node *leetcode.ListNode) {

	fmt.Print("\n[")
	for node != nil {
		fmt.Printf("%d,", node.Val)
		node = node.Next
	}
	fmt.Println("]")
}

func toInt(l1 *leetcode.ListNode) int {

	var aStep = 10
	var arr1 = l1.Val
	for l1.Next != nil {
		arr1 = arr1 + aStep*l1.Next.Val
		l1 = l1.Next
		aStep = aStep * 10

	}
	return arr1
}

func createTreeNode() *leetcode.TreeNode {

	var t7 = &leetcode.TreeNode{Val: 7, Left: nil, Right: nil}
	var t8 = &leetcode.TreeNode{Val: 8, Left: nil, Right: nil}
	var t6 = &leetcode.TreeNode{Val: 6, Left: t7, Right: t8}
	var t5 = &leetcode.TreeNode{Val: 5, Left: nil, Right: nil}
	var t4 = &leetcode.TreeNode{Val: 4, Left: nil, Right: t6}
	var t3 = &leetcode.TreeNode{Val: 3, Left: nil, Right: t5}
	var t2 = &leetcode.TreeNode{Val: 2, Left: t4, Right: nil}

	return &leetcode.TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	/*
	                    1
	   		2 --------------3
	   	4                             5
	   	    6
	   	7       8

	   	如图1所示，三种遍历方法(人工)得到的结果分别是：

	   先序：1 2 4 6 7 8 3 5
	   中序：4 7 6 8 2 1 3 5
	   后序：7 8 6 4 2 5 3 1

	   三种遍历方法的考查顺序一致，得到的结果却不一样，原因在于：

	   先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)

	   中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)

	   后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

	*/

}

func executeNode() {
	var cx = createListNode([]int{8, 4, 5})

	var c1 = createListNode([]int{4, 1})

	var c2 = createListNode([]int{5, 6, 1})
	var a = createListNodeWith2(c1, cx)
	var b = createListNodeWith2(c2, cx)

	var mg = leetcode.GetIntersectionNode(a, b)

	printListNode(mg)
	//leetcode.AddTwoNumbers()
}

func exchangeSlice(slice []int) {
	// for k, v := range slice {
	// 	slice[k] = v * 2
	// }

	slice = append(slice, 33)
}

func main() {

	leetcode.TestLevelOrder()

	fmt.Println("LLLK")

	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)
	// exchangeSlice(slice)
	// fmt.Println(slice)

	// var rest []int = nil

	// rest = append(rest, 1)

	// rest = append(rest, 2)

	// fmt.Printf("%v\n", rest)
	var tn = createTreeNode()

	var mg = leetcode.ZigzagLevelOrder(tn)

	fmt.Printf("%v\n", mg)

}
