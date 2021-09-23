package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	var nums =[]int {1,2,3,4}
	v:= Hello(nums)

	assert.Equal(t, v, 10 )

}