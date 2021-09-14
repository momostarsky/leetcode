package leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serialize   a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	sb := &strings.Builder{}

	var queue []*TreeNode = nil

	queue = append(queue, root)

	//  queue=append(queue, nil)

	for len(queue) > 0 {
		var topNode = queue[0]
		queue = queue[1:]
		if topNode == nil {
			sb.WriteString("null,")
			//  queue is Empty And current Level is Over
			//if len(queue)  >0 {
			//   // has more than 1 node waiting for processing
			//    for len(queue)>0 &&  queue[0] == nil {
			//	    pstr= pstr+"null,"
			//		queue =queue[1:]
			//   }
			//   if len(queue) >0 {
			//	   queue = append(queue,nil )
			//   }
			//
			//}
		} else {

			sb.WriteString(strconv.Itoa(topNode.Val))
			sb.WriteString(",")
			queue = append(queue, topNode.Left)

			queue = append(queue, topNode.Right)

		}

		queue = queue
	}

	sb.WriteString("]")
	return sb.String()

}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	return nil
}
