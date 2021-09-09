package leetcode

import "fmt"

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	fmt.Printf("%b , %d\n", 1<<n, 1<<n)
	maskLen := 1 << n
	for mask := 0; mask < maskLen; mask++ {
		set := []int{}
		fmt.Printf("mask:%b,%d\n", mask, mask)
		for i, v := range nums {
			var bw = mask >> i
			fmt.Printf("-----mask>>%d:%b,%d\n", i, bw, bw)
			if bw&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func Runsubset() {

	var nums = []int{1, 2, 3, 4}

	var rx = subsets(nums)

	fmt.Printf("\n%v\n", rx)

}

/*
deteck
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xvkwe2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func deteck(board [][]byte, i, j int, word []byte, rows int, cols int, wo string) []string {

	var dataGrid [][]byte = nil

	var res0 []byte = nil
	var res1 []byte = nil
	var res2 []byte = nil
	var res3 []byte = nil
	res0 = append(res0, word[0])
	res1 = append(res1, word[0])
	res2 = append(res2, word[0])
	res3 = append(res3, word[0])

	dataGrid = append(dataGrid, res0)
	dataGrid = append(dataGrid, res1)
	dataGrid = append(dataGrid, res2)
	dataGrid = append(dataGrid, res3)

	return []string{
		string(dataGrid[0]),
		string(dataGrid[1]),
		string(dataGrid[2]),
		string(dataGrid[3]),
	}

}

/*
思路与算法


如果当前已经访问到字符串的末尾，且对应字符依然匹配，此时直接返回 \texttt{true}true。
否则，遍历当前位置的所有相邻位置。如果从某个相邻位置出发，能够搜索到子串 \textit{word}[k+1..]word[k+1..]，则返回 \texttt{true}true，否则返回 \texttt{false}false。
这样，我们对每一个位置 (i,j)(i,j) 都调用函数 \text{check}(i, j, 0)check(i,j,0) 进行检查：只要有一处返回 \texttt{true}true，就说明网格中能够找到相应的单词，否则说明不能找到。

为了防止重复遍历相同的位置，需要额外维护一个与 \textit{board}board 等大的 \textit{visited}visited 数组，用于标识每个位置是否被访问过。每次遍历相邻位置时，需要跳过已经被访问的位置。

*/

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {

	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	/// 从 第 [i,j]个位置开始判断 [i,j]的字符是否和 word[k]相等
	var check func(i, j, k int) bool

	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { //最后一个字母了, 说明单词存在于网格中
			return true
		}
		visited[i][j] = true
		//---------------------回溯时还原已访问的单元格
		defer func() {
			visited[i][j] = false
		}()

		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < rows && 0 <= newJ && newJ < cols && !visited[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false

}

func Runexist() {

	// var grid = [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }

	// var exis = exist(grid, "ABCCED")

	// fmt.Printf("%v\n", exis)
	//"SEE"

	/*
					[["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]]
				"ABCESEEEFS"

				[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
		      "ABCB"
	*/
	var g2 = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'F'},
	}
	var exis2 = exist(g2, "ABCB")

	fmt.Printf("%v\n", exis2)

}
