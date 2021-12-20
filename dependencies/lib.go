package dependencies

type Operation struct {
	ID           int
	marked       bool
	dependencies []*Operation
}

var closeList []int

func main() {
	closeList = []int{}
	dependencies := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},
	}
	operations := []*Operation{}
	for i := 0; i < 4; i++ {
		op := &Operation{
			ID:           i,
			dependencies: []*Operation{},
		}
		operations = append(operations, op)
	}
	createGraph(operations, dependencies)
	for _, op := range operations {
		dfs(op)
	}
}

func createGraph(operations []*Operation, dependencies [][]int) {
	for _, d := range dependencies {
		operations[d[0]].dependencies = append(operations[d[0]].dependencies, operations[d[1]])
	}
}

func dfs(op *Operation) {
	if op.marked {
		return
	}
	for _, d := range op.dependencies {
		dfs(d)
	}
	op.marked = true
	closeList = append(closeList, op.ID)
}
