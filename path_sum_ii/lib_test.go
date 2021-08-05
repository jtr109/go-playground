package pathsumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	input := TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{13, nil, nil},
			&TreeNode{
				4,
				&TreeNode{5, nil, nil},
				&TreeNode{1, nil, nil},
			},
		},
	}
	targetSum := 22
	expected := [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}
	assert.Equal(t, expected, pathSum(&input, targetSum))
}

func TestSubmission1(t *testing.T) {
	var input *TreeNode = nil
	targetSum := 1
	expected := [][]int{}
	assert.Equal(t, expected, pathSum(input, targetSum))
}

func TestSubmission2(t *testing.T) {
	input := TreeNode{
		1,
		nil,
		&TreeNode{2, nil, nil},
	}
	targetSum := 1
	expected := [][]int{}
	assert.Equal(t, expected, pathSum(&input, targetSum))

}
