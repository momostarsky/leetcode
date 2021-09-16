package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func treeNodeEqual(t1 *TreeNode, t2 *TreeNode) bool {

	if t1 == nil && t2 == nil {
		return true
	}
	if assert.ObjectsAreEqual(t1.Val, t2.Val) {

		return treeNodeEqual(t1.Left, t2.Left) && treeNodeEqual(t1.Right, t2.Right)

	} else {
		return false
	}

}

func TestCodec_Serialize(t *testing.T) {

	var tcode = CodecConstructor()

	var t5 = &TreeNode{Val: 5, Left: nil, Right: nil}
	var t4 = &TreeNode{Val: 4, Left: nil, Right: nil}
	var t3 = &TreeNode{Val: 3, Left: t4, Right: t5}
	var t2 = &TreeNode{Val: 2, Left: nil, Right: nil}
	var t1 = &TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	var res = tcode.serialize(t1)
	var dt = tcode.deserialize(res)

	assert.True(t, treeNodeEqual(t1, dt))

}
func TestFib(t *testing.T) {
	var Fib6a = Fib(6)
	var Fib6b = Fib2(6)

	assert.True(t, Fib6a == Fib6b)

	var f3 = Fib3(6)
	var f4 = Fib4(6)

	assert.True(t, f3 == f4)

}

func TestMergeSort(t *testing.T) {
	var data = []int{0, 1, 9, 3, 2, 100, 77, 13, 20}
	var RES = MergeSort(data)

	assert.EqualValues(t, []int{
		0, 1, 2, 3, 9, 13, 20, 77, 100,
	}, RES)
}
