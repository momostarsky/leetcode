package leetcode

//常数时间插入、删除和获取随机元素
//实现RandomizedSet 类：
//
//RandomizedSet() 初始化 RandomizedSet 对象
//bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
//bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
//int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
//你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

//
//作者：dfzhou6
//链接：https://leetcode-cn.com/problems/insert-delete-getrandom-o1/solution/gojian-dan-yi-dong-by-dfzhou6-f6du/

import "math/rand"

type RandomizedSet struct {
	sl []int
	m  map[int]int
}

func MakeRandomizedSet() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

func (p *RandomizedSet) Insert(val int) bool {
	if _, ok := p.m[val]; ok {
		return false
	}
	p.sl = append(p.sl, val)
	p.m[val] = len(p.sl) - 1
	return true
}

func (p *RandomizedSet) Remove(val int) bool {
	if _, ok := p.m[val]; !ok {
		return false
	}
	curIdx := p.m[val]
	lastIdx := len(p.sl) - 1
	lastVal := p.sl[lastIdx]
	p.sl[curIdx], p.sl[lastIdx] = lastVal, val
	p.sl = p.sl[:lastIdx]
	p.m[lastVal] = curIdx
	delete(p.m, val)
	return true
}

func (p *RandomizedSet) GetRandom() int {
	return p.sl[rand.Int()%len(p.sl)]
}
