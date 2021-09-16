package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serialize   a tree to a single string.
func (p *Codec) serialize(root *TreeNode) string {

	sb := &strings.Builder{}

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {

		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(fmt.Sprintf("%d,", node.Val))

		dfs(node.Left)

		dfs(node.Right)

	}

	dfs(root)

	return sb.String()

}

// Deserializes your encoded data to tree.
func (p *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()

}
