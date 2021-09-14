package leetcode

import (
	"fmt"
	"testing"
)

func TestCodec_Serialize(t *testing.T) {

	var tcode = Constructor()

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

	fmt.Printf("%s", res)

}
