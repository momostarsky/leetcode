package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSorut(t *testing.T) {
	var ax = []int{5, 1, 3, 7, 0, 100, 11}

	QuickSort(ax)
	fmt.Printf("%v\n", ax)

	assert.EqualValues(t, ax, []int{
		0, 1, 3, 5, 7, 11, 100,
	})
}

func TestQuickSort2(t *testing.T) {
	var ax = []int{5, 1, 3, 7, 0, 100, 11}

	QuickSortIt(ax)
	fmt.Printf("%v\n", ax)

	assert.EqualValues(t, ax, []int{
		0, 1, 3, 5, 7, 11, 100,
	})
}

func TestTitleToNumber(t *testing.T) {
	var ax = titleToNumber("AZ")

	assert.Equal(t, 27, ax)
}
