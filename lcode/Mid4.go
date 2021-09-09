package leetcode

import "fmt"

type PNode struct {
	Val  int
	Pre  *PNode
	Subs []*PNode
}

func createSubNode(pre *PNode, pv int, subVal []int) *PNode {

	var p = &PNode{
		Val: pv,
		Pre: pre,
	}
	fmt.Printf("createNode :%d, %v\n", pv, subVal)
	for i := 0; i < len(subVal); i++ {
		var sub []int = nil
		sub = append(sub, subVal[:i]...)
		sub = append(sub, subVal[i+1:]...)
		var cNode = createSubNode(p, subVal[i], sub)
		p.Subs = append(p.Subs, cNode)
	}
	return p

}

func makeNodes(nums []int) []*PNode {

	var nodes []*PNode = nil
	for i := 0; i < len(nums); i++ {
		var sub []int = nil
		sub = append(sub, nums[:i]...)
		sub = append(sub, nums[i+1:]...)
		var cNode = createSubNode(nil, nums[i], sub)
		nodes = append(nodes, cNode)
	}
	return nodes

}

var res [][]int = nil

func enumNodes(p *PNode) {

	if len(p.Subs) == 0 {
		fmt.Printf("%d ", p.Val)

		var mg []int = nil
		var cp = p
		for cp != nil {
			mg = append(mg, cp.Val)
			cp = cp.Pre
		}
		res = append(res, mg)
		fmt.Printf("-->: %v\n ", mg)
		return
	}
	for i := 0; i < len(p.Subs); i++ {
		enumNodes(p.Subs[i])
	}

}

func permute(nums []int) [][]int {
	//-- 构建生成树
	var nl = len(nums)
	if nl == 1 {
		return [][]int{nums}
	}
	if nl == 2 {
		return [][]int{
			{nums[0], nums[1]},
			{nums[1], nums[0]},
		}
	}
	res = nil
	var p = makeNodes(nums)
	for _, v := range p {
		enumNodes(v)
	}

	return res

}
func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

//----优化算法
func permuteOption(nums []int) [][]int {
	N := len(nums)
	if N == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	var backtrack func(int)
	backtrack = func(index int) {
		fmt.Printf("index is :%d\n", index)
		// 结束条件
		if index == N {
			// 保存结果
			res = append(res, append([]int(nil), nums...))
			return
		}
		for i := index; i < N; i++ {
			// 做出选择
			swap(nums, i, index)
			// 递归
			backtrack(index + 1)
			// 撤销选择
			swap(nums, i, index)
		}
	}
	backtrack(0)
	return res

}

/*
回溯模版
void backtrack(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtrack(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

思路：回溯
1、路径：已经选择过的数字
2、选择列表： 可以选择的数字（顺序从左往右，不能回头）
3、结束条件： 无
ps: 拿 nums={1,2,3} 手动排一下, 就会有思路了



*/

func Runpermute() {

	var nums = []int{1, 2, 3}

	var rx = permuteOption(nums)

	fmt.Printf("\n%v\n", rx)

}
