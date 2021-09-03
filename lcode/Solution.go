package leetcode

import "math/rand"

type Solution struct {
	Nums []int
}

func SolutionConstructor(nums []int) Solution {
	return Solution{
		Nums: nums,
	}

}

/** Resets the array to its original configuration and return it. */
func (p *Solution) Reset() []int {

	return p.Nums

}

func (p *Solution) Shuffle() []int {
	lx := len(p.Nums)
	var nArr = make([]int, lx)
	copy(nArr, p.Nums)

	for i := lx - 1; i > 0; i-- {
		//--随机生成一个下表 rd , 取值范围在[0,i]  ( rd>=0 && rd <= i )
		var rd = rand.Int() % (i + 1)
		//  nArr[rd] 最后一个 (nArr[i]) 交换
		//  由于i在递减,所有每次交换之后 nArr[rd]从下次交换中剔除
		nArr[rd], nArr[i] = nArr[i], nArr[rd]

	}

	return nArr

}

/** Returns a random shuffling of the array. */

///==================递归算法不对
// func (p *Solution) Shuffle() []int {
// 	lx := len(p.Nums)
// 	var nArr = make([]int, lx)
// 	copy(nArr, p.Nums)
// 	p.innerShuffle(nArr[0:])
// 	return nArr
// }

// func (p *Solution) innerShuffle(nums []int) {
// 	var lx = len(nums)
// 	if lx <= 1 {
// 		return
// 	} else if lx == 2 {
// 		nums[0], nums[1] = nums[1], nums[0]
// 	} else {
// 		var rd = rand.Int() % lx
// 		nums[rd], nums[0] = nums[0], nums[rd]
// 		p.innerShuffle(nums[1:])
// 	}
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
