// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge-2021/638/week-3-september-15th-september-21st/3977/

package spiralmatrix

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

type Cursor struct {
	left      int
	right     int
	top       int
	bottom    int
	row       int
	col       int
	direction Direction
}

func newCursor(matrixWidth int, matrixHeight int) Cursor {
	return Cursor{
		right:  matrixWidth - 1,
		bottom: matrixHeight - 1,
	}
}

func (cursor *Cursor) exhausted() bool {
	// the cursor out of boundary when exhaust all cells
	return cursor.col < cursor.left ||
		cursor.col > cursor.right ||
		cursor.row < cursor.top ||
		cursor.row > cursor.bottom
}

func (cursor *Cursor) value(matrix *[][]int) int {
	return (*matrix)[cursor.row][cursor.col]
}

func (cursor *Cursor) reachBoundary() bool {
	switch cursor.direction {
	case Right:
		return cursor.col == cursor.right
	case Down:
		return cursor.row == cursor.bottom
	case Left:
		return cursor.col == cursor.left
	default: // up
		return cursor.row == cursor.top
	}
}

func (cursor *Cursor) narrow() {
	switch cursor.direction {
	case Right:
		cursor.top += 1
	case Down:
		cursor.right -= 1
	case Left:
		cursor.bottom -= 1
	default: // up
		cursor.left += 1
	}
}

func (cursor *Cursor) changeDirection() {
	switch cursor.direction {
	case Right:
		cursor.direction = Down
	case Down:
		cursor.direction = Left
	case Left:
		cursor.direction = Up
	default: // up
		cursor.direction = Right
	}
}

func (cursor *Cursor) move() {
	switch cursor.direction {
	case Right:
		cursor.col += 1
	case Down:
		cursor.row += 1
	case Left:
		cursor.col -= 1
	default: // up
		cursor.row -= 1
	}
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	cursor := newCursor(len(matrix[0]), len(matrix))
	for {
		if cursor.exhausted() {
			break
		}
		result = append(result, cursor.value(&matrix))
		if cursor.reachBoundary() {
			cursor.narrow()
			cursor.changeDirection()
		}
		cursor.move()
	}
	return result
}
