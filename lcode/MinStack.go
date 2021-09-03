package leetcode

type vNode struct {
	Val  int
	Min  int
	Next *vNode
}

type MinStack struct {
	nodes *vNode
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {

	return MinStack{}

}

func (p *MinStack) Push(val int) {

	if p.nodes == nil {
		p.nodes = &vNode{
			Val: val, Min: val,
		}

	} else {

		var mv = p.nodes.Min
		if val < mv {
			mv = val
		}
		cp := &vNode{
			Val: val, Min: mv,
		}
		cp.Next = p.nodes
		p.nodes = cp
	}

}

func (p *MinStack) Pop() {

	if p.nodes != nil {
		p.nodes = p.nodes.Next
	}
}

func (p *MinStack) Top() int {
	return p.nodes.Val

}

func (p *MinStack) GetMin() int {
	return p.nodes.Min

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
