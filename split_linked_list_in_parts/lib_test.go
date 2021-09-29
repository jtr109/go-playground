package splitlinkedlistinparts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	assert.Equal(t, 3, count(listNode))
}

func TestLengthList(t *testing.T) {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 10,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	assert.Equal(t, []int{4, 3, 3}, lengthList(listNode, 3))
}

func TestExample1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	result := splitListToParts(head, 3)
	assert.Equal(t, 1, result[0].Val)
	assert.Equal(t, (*ListNode)(nil), result[0].Next)
	assert.Equal(t, 2, result[1].Val)
	assert.Equal(t, (*ListNode)(nil), result[1].Next)
	assert.Equal(t, 3, result[2].Val)
	assert.Equal(t, (*ListNode)(nil), result[2].Next)
}

func TestExample2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 10,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	result := splitListToParts(head, 3)
	assert.Equal(t, 1, result[0].Val)
	assert.Equal(t, 2, result[0].Next.Val)
	assert.Equal(t, 3, result[0].Next.Next.Val)
	assert.Equal(t, 4, result[0].Next.Next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), result[0].Next.Next.Next.Next)
	assert.Equal(t, 5, result[1].Val)
	assert.Equal(t, 6, result[1].Next.Val)
	assert.Equal(t, 7, result[1].Next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), result[1].Next.Next.Next)
	assert.Equal(t, 8, result[2].Val)
	assert.Equal(t, 9, result[2].Next.Val)
	assert.Equal(t, 10, result[2].Next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), result[2].Next.Next.Next)
}

func TestSubmission1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	res := splitListToParts(head, 5)
	assert.Equal(t, head, res[0])
	assert.Equal(t, 2, res[1].Val)
	assert.Equal(t, (*ListNode)(nil), res[1].Next)
	assert.Equal(t, 3, res[2].Val)
	assert.Equal(t, (*ListNode)(nil), res[2].Next)
	assert.Equal(t, (*ListNode)(nil), res[3])
	assert.Equal(t, (*ListNode)(nil), res[4])
}

func TestCut(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	next := cut(head, 1)
	assert.Equal(t, 1, head.Val)
	assert.Equal(t, (*ListNode)(nil), head.Next)
	assert.Equal(t, 2, next.Val)
	assert.Equal(t, 3, next.Next.Val)
	assert.Equal(t, (*ListNode)(nil), next.Next.Next)
}

// func TestSubmission2(t *testing.T) {
// 	head := (*ListNode)(nil)
// 	assert.Equal(t, []*ListNode{nil, nil, nil}, splitListToParts(head, 3))
// }
