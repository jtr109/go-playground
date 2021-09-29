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

func TestMove(t *testing.T) {
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
	assert.Equal(t, 5, move(listNode, 4).Val)
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
	assert.Equal(t, []*ListNode{head, head.Next, head.Next.Next}, splitListToParts(head, 3))
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
	assert.Equal(
		t, []*ListNode{head, head.Next.Next.Next.Next, head.Next.Next.Next.Next.Next.Next.Next},
		splitListToParts(head, 3),
	)
}
